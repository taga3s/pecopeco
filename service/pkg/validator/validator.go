package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func Get() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}
