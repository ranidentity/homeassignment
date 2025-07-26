package common

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()

		switch tag {
		case "required":
			errors[field] = fmt.Sprintf("%s is required", field)
		case "gte":
			errors[field] = fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
		case "decimal2":
			errors[field] = fmt.Sprintf("%s doesn't follow allowed format 0.00", field)
		default:
			errors[field] = fmt.Sprintf("%s failed on %s validation", field, tag)
		}
	}
	return errors
}
