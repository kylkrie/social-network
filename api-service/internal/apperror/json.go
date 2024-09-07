package apperror

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type JSONUnmarshalError struct {
	Errors []ValidationError `json:"errors"`
}

func (je JSONUnmarshalError) Error() string {
	var messages []string
	for _, err := range je.Errors {
		messages = append(messages, fmt.Sprintf("%s: %s", err.Field, err.Message))
	}
	return strings.Join(messages, "; ")
}

func NewJSONUnmarshalError(err interface{}) *AppError {
	var jsonError JSONUnmarshalError

	switch e := err.(type) {
	case *json.UnmarshalTypeError:
		jsonError.Errors = append(jsonError.Errors, ValidationError{
			Field: e.Field,

			Message: formatJSONUnmarshalTypeError(e),
		})
	case *json.SyntaxError:
		jsonError.Errors = append(jsonError.Errors, ValidationError{
			Field:   "JSON",
			Message: fmt.Sprintf("Invalid JSON syntax at position %d", e.Offset),
		})
	default:

		// If it's neither UnmarshalTypeError nor SyntaxError, return a generic error
		log.Error().Any("error", e).Msg("Unhandled JSON Error")
		return New(http.StatusBadRequest, "Invalid JSON input")
	}

	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "JSON parsing failed",

		Data: jsonError,
	}
}

func formatJSONUnmarshalTypeError(err *json.UnmarshalTypeError) string {
	return fmt.Sprintf("Invalid type. Expected %s, got %s", err.Type, err.Value)
}
