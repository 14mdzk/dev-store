package validator

import "github.com/go-playground/validator"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Check(value interface{}) bool {
	err := validate.Struct(value)

	if err != nil {
		return false
	}

	return true
}
