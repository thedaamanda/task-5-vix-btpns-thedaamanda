package delivery

import (
	"net/http"
	helper "project/helpers"
	middleware "project/middlewares"
	model "project/models"
	"project/request"

	"github.com/labstack/echo"
)

type photoDelivery struct {
	photoUsecase model.PhotoUsecase
}

type PhotoDelivery interface {
	Mount(group *echo.Group)
}

func NewPhotoDelivery(PhotoUsecase model.PhotoUsecase) PhotoDelivery {
	return &photoDelivery{
		photoUsecase: PhotoUsecase,
	}
}

func (p *photoDelivery) Mount(group *echo.Group) {
	customMiddleware := middleware.Init()
	group.Use(customMiddleware.Authentication)
	group.Use(customMiddleware.Authorization)
	group.POST("", p.StorePhotoHandler)
	group.GET("", p.FetchPhotoHandler)
	group.PUT("/:id", p.UpdatePhotoHandler, customMiddleware.PhotoAuthorization)
	group.DELETE("/:id", p.DeletePhotoHandler, customMiddleware.PhotoAuthorization)
}

func (p *photoDelivery) StorePhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var request request.PhotoRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ValidationResponse(c, http.StatusBadRequest, "Invalid Request Body", helper.GetErrorMessages(err))
	}

	id := helper.GetUserID(c)

	photo, err := p.photoUsecase.CreatePhoto(ctx, id, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusCreated, "Your photo has been successfully created", photo)
}

func (p *photoDelivery) FetchPhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()

	photos, err := p.photoUsecase.GetPhotoList(ctx)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusOK, "List of photos", photos)
}

func (p *photoDelivery) UpdatePhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var request request.PhotoRequest

	paramID := c.Param("id")

	id, err := helper.StringToInt(paramID)
	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ValidationResponse(c, http.StatusBadRequest, "Invalid Request Body", helper.GetErrorMessages(err))
	}

	photo, err := p.photoUsecase.UpdatePhoto(ctx, id, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusOK, "Your photo has been successfully updated", photo)
}

func (p *photoDelivery) DeletePhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")

	id, err := helper.StringToInt(paramID)
	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	if err := p.photoUsecase.DeletePhoto(ctx, id); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.MessageResponse(c, http.StatusOK, "Your photo has been successfully deleted")
}
