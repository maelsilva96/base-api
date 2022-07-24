package errors

import (
	"fmt"
	"strings"
)

type FieldError struct {
	errorMessages map[string][]string
}

func NewFieldError(errors map[string][]string) *FieldError {
	return &FieldError{
		errorMessages: errors,
	}
}

func (m *FieldError) Error() string {
	message := ""
	for n, messages := range m.errorMessages {
		message += fmt.Sprintf("O campo (%s) tem os seguintes erros: %s\n", n, strings.Join(messages, ","))
	}
	return message
}

func (m *FieldError) GetErrors() map[string][]string {
	return m.errorMessages
}
