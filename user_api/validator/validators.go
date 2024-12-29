package validator

import (
	lancet "github.com/duke-git/lancet/v2/validator"
	"github.com/go-playground/validator/v10"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	return lancet.IsChineseMobile(mobile)
}
