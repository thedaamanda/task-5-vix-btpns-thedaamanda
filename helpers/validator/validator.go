package validator

import (
	"github.com/go-playground/validator/v10"
)

type GoPlaygroundValidator struct {
	Validator *validator.Validate
}

func (v *GoPlaygroundValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
