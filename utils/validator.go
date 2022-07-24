package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/maelsilva96/base-api/entity/errors"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidEntity(entity interface{}, mapFields map[string]string) error {
	err := validate.Struct(entity)
	if err != nil {
		errorMessages := make(map[string][]string)
		for _, m := range err.(validator.ValidationErrors) {
			field := m.StructField()
			if val, ok := mapFields[field]; ok {
				field = val
			}
			if _, ok := errorMessages[field]; !ok {
				errorMessages[field] = []string{}
			}
			message := getMessageToError(m.Tag(), m.Param())
			errorMessages[field] = append(errorMessages[field], message)
		}
		return errors.NewFieldError(errorMessages)
	}
	return nil
}

func getMessageToError(errorType string, param string) string {
	switch errorType {
	case "required":
		return "Obrigatório o preenchimento do campo!"
	case "min":
		return fmt.Sprintf("Mínimo de caracteres desse campo é %s!", param)
	case "max":
		return fmt.Sprintf("Máximo de caracteres desse campo é %s!", param)
	default:
		return "O valor enviado não é valido!"
	}
}
