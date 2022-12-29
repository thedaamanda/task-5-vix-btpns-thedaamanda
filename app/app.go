package app

import (
	"fmt"
	"log"
	"net/http"
	"project/config"
	"project/delivery"
	"project/helpers/validator"
	"project/repository"
	"project/usecase"

	validatorEngine "github.com/go-playground/validator/v10"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	server struct {
		httpServer *echo.Echo
		cfg        config.Config
	}

	Server interface {
		Run()
	}
)

func InitServer(cfg config.Config) Server {
	e := echo.New()
	e.HideBanner = true
	e.Validator = &validator.GoPlaygroundValidator{
		Validator: validatorEngine.New(),
	}

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	return &server{
		httpServer: e,
		cfg:        cfg,
	}
}

func (s *server) Run() {
	userRepo := repository.NewUserRepository(s.cfg)
	userUsercase := usecase.NewUserUsecase(userRepo)
	userDelivery := delivery.NewUserDelivery(userUsercase)
	userGroup := s.httpServer.Group("/users")
	userDelivery.Mount(userGroup)

	if err := s.httpServer.Start(fmt.Sprintf("%s:%d", s.cfg.ServiceHost(), s.cfg.ServicePort())); err != nil {
		log.Panic(err)
	}
}
