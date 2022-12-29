package helper

import (
	"project/helpers/token"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetUserID(c echo.Context) int {
	userData := c.Get("userData").(*token.NewTokenData)
	return userData.UserID
}

func errorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This field must be a valid email"
	case "min":
		if fe.Type().String() == "string" {
			return "This field must be at least " + fe.Param() + " characters"
		}
		return "This field must be at least " + fe.Param()
	default:
		return "Unknown error"
	}
}

func GetErrorMessages(err error) []ErrorMessage {
	var errors []ErrorMessage

	for _, fe := range err.(validator.ValidationErrors) {
		errors = append(errors, ErrorMessage{
			Field:   fe.Field(),
			Message: errorMessage(fe),
		})
	}

	return errors
}

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func CheckAuthorization(requester, target int) bool {
	return requester == target
}

func GetID(c echo.Context) int {
	id, _ := StringToInt(c.Param("id"))
	return id
}
