package apperror

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

// New creates a new AppError
func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func ToAppError(err interface{}) *AppError {
	switch e := err.(type) {
	case *AppError:
		return e
	case AuthErrorCode:
		return HandleAuthError(e)
	case error:
		return handleStandardError(e)
	default:
		log.Error().Stack().Msg("Unknown AppError")
		return New(http.StatusInternalServerError, "An unknown error occurred")
	}
}

// handleStandardError deals with standard error types
func handleStandardError(err error) *AppError {
	// Delegate to specific error handlers
	if dbErr := HandleDBError(err); dbErr != nil {
		return dbErr
	}

	// If it's not a specific type we recognize, return a generic error
	return New(http.StatusInternalServerError, err.Error())
}
