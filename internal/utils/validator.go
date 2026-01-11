package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// FormatValidationError formats validation errors into a map
func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors[e.Field()] = getErrorMsg(e)
		}
	}

	return errors
}

func getErrorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	default:
		return "Invalid value"
	}
}
