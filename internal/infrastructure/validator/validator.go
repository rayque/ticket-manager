package validator

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Error struct {
	Message string `json:"message"`
}
type Errors []Error

func ValidateRequest(input interface{}) Errors {
	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		var validationErrors Errors

		if err := validate.Struct(input); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors, Error{
					Message: strings.Split(e.Error(), "Error:")[1],
				})
			}

			return validationErrors
		}

	}

	return nil
}
