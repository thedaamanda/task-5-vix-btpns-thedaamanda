package delivery

import (
	"net/http"
	helper "project/helpers"
	middleware "project/middlewares"
	model "project/models"
	"project/request"

	"github.com/labstack/echo"
)

type userDelivery struct {
	userUsecase model.UserUsecase
}

type UserDelivery interface {
	Mount(group *echo.Group)
}

func NewUserDelivery(UserUsecase model.UserUsecase) UserDelivery {
	return &userDelivery{
		userUsecase: UserUsecase,
	}
}

func (u *userDelivery) Mount(group *echo.Group) {
	customMiddleware := middleware.Init()
	group.POST("/register", u.RegisterHandler)
	group.POST("/login", u.LoginHandler)
	group.PUT("/:id", u.EditUserHandler, customMiddleware.Authentication)
	group.DELETE("/:id", u.DeleteUserHandler, customMiddleware.Authentication)
}

func (u *userDelivery) RegisterHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var request request.CreateUserRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ValidationResponse(c, http.StatusBadRequest, "Invalid Request Body", helper.GetErrorMessages(err))
	}

	user, err := u.userUsecase.Register(ctx, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusCreated, "Your account has been successfully created", user)
}

func (u *userDelivery) LoginHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var request request.LoginRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ValidationResponse(c, http.StatusBadRequest, "Invalid Request Body", helper.GetErrorMessages(err))
	}

	user, err := u.userUsecase.Login(ctx, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusOK, "Login Success", user)
}

func (u *userDelivery) EditUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")

	id, err := helper.StringToInt(paramID)
	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	var request request.UpdateUserRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return helper.ValidationResponse(c, http.StatusBadRequest, "Invalid Request Body", helper.GetErrorMessages(err))
	}

	user, err := u.userUsecase.UpdateUser(ctx, id, request)

	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.SuccessResponse(c, http.StatusOK, "Your account has been successfully updated", user)
}

func (u *userDelivery) DeleteUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	paramID := c.Param("id")

	id, err := helper.StringToInt(paramID)
	if err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	if err := u.userUsecase.DeleteUser(ctx, id); err != nil {
		return helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return helper.MessageResponse(c, http.StatusOK, "Your account has been successfully deleted")
}
