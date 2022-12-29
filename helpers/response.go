package helper

import "github.com/labstack/echo"

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Success struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Validation struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return Response(c, statusCode, Success{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}

func MessageResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}

func ValidationResponse(c echo.Context, statusCode int, message string, errors interface{}) error {
	return Response(c, statusCode, Validation{
		Code:    statusCode,
		Message: message,
		Errors:  errors,
	})
}
