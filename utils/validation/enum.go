package validation

import "github.com/go-playground/validator/v10"

type Enum interface {
	IsValid() bool
}

type Status string

const (
	Draft   Status = "Draft"
	Publish Status = "Publish"
	Trash   Status = "Trash"
)

func EnumValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	switch value {
	case string(Draft), string(Publish), string(Trash):
		return true
	}

	return false
}
