package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) (map[string]string, error) {
	err := validate.Struct(s)
	if err == nil {
		return nil, nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, err
	}

	errorsMap := make(map[string]string)
	for _, e := range validationErrors {
		errorsMap[e.Field()] = getErrorMsg(e)
	}

	return errorsMap, nil
}

func getErrorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "email":
		return "Invalid email address"
	default:
		return e.Field() + " is invalid"
	}
}
