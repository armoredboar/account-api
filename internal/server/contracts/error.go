package contracts

import (
	"fmt"

	"github.com/armoredboar/account-api/internal/validation"
)

// Error indicates the field and reason of an processing error.
type Error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func CreateError(code int, field string, parameters ...interface{}) Error {

	var message string

	var newError Error

	// Check if the specified code exists as an error resource.
	if resourceMessage, ok := validation.Errors[code]; ok {

		if len(parameters) > 0 {

			args := append([]interface{}{field}, parameters...)

			message = fmt.Sprintf(resourceMessage, args...)
		} else {
			message = fmt.Sprintf(resourceMessage, field)
		}

		newError = Error{Code: code, Field: field, Message: message}

	} else {
		if len(parameters) > 0 {
			message = fmt.Sprintf("%s", parameters[0])
		}

		newError = Error{Field: field, Message: message}
	}

	return newError
}
