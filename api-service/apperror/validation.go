// apperror/validation.go

package apperror

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func (ve ValidationErrors) Error() string {
	var messages []string
	for _, err := range ve.Errors {
		messages = append(messages, fmt.Sprintf("%s: %s", err.Field, err.Message))
	}
	return strings.Join(messages, "; ")
}

func NewValidationError(err error) *AppError {
	var validationErrors ValidationErrors

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			validationErrors.Errors = append(validationErrors.Errors, ValidationError{
				Field:   e.Field(),
				Message: formatValidationError(e),
			})
		}
	} else {
		// If it's not a validator.ValidationErrors, just return a generic error
		return New(http.StatusBadRequest, "Invalid input")
	}

	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		Data:    validationErrors,
	}
}

func formatValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return fmt.Sprintf("Maximum length is %s", err.Param())
	case "min":
		return fmt.Sprintf("Minimum length is %s", err.Param())
	case "email":
		return "Invalid email format"
	default:
		return "Invalid value"
	}
}

func RegisterCustomValidator() {
	// fix field names to show json name
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
