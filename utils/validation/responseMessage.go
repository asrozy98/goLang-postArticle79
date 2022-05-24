package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Message(err any) any {
	errorMessages := []any{}
	for _, e := range err.(validator.ValidationErrors) {
		switch e.ActualTag() {
		case "min":
			errorMessage := fmt.Sprintf("Error on %s, because: minimum value should be greater than %s character", e.Field(), e.Param())
			errorMessages = append(errorMessages, errorMessage)
		case "enum":
			errorMessage := fmt.Sprintf("Error on %s, because: value must be %s", e.Field(), e.Param())
			errorMessages = append(errorMessages, errorMessage)
		default:
			errorMessage := fmt.Sprintf("Error on %s, because: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	return errorMessages
}
