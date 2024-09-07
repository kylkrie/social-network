package apperror

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

const (
	ErrResourceNotFound     = "Resource not found"
	ErrResourceExists       = "Resource already exists"
	ErrDatabaseTimeout      = "Database operation timed out"
	ErrForeignKeyViolation  = "Operation violates foreign key constraint"
	ErrCheckConstraint      = "Operation violates check constraint"
	ErrSerializationFailure = "Concurrent transaction conflict, please retry"
	ErrUnknownDatabase      = "Unknown database error occurred"
)

func HandleDBError(err error) *AppError {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return New(http.StatusNotFound, ErrResourceNotFound)
	case errors.Is(err, context.DeadlineExceeded):
		return New(http.StatusGatewayTimeout, ErrDatabaseTimeout)
	case errors.Is(err, context.Canceled):
		return New(http.StatusInternalServerError, "Database operation was canceled")
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation
			return New(http.StatusConflict, ErrResourceExists)
		case "23503": // foreign_key_violation
			return New(http.StatusBadRequest, ErrForeignKeyViolation)
		case "23514": // check_violation
			return New(http.StatusBadRequest, ErrCheckConstraint)
		case "40001": // serialization_failure
			return New(http.StatusConflict, ErrSerializationFailure)
		case "3D000": // invalid_catalog_name
			return New(http.StatusInternalServerError, ErrUnknownDatabase)
		default:
			// Log the unknown postgres error for debugging
			log.Error().Str("pgcode", pgErr.Code).Str("message", pgErr.Message).Msg("Unhandled Postgres error")
			return New(http.StatusInternalServerError, fmt.Sprintf("Database error: %s", pgErr.Message))
		}
	}

	// If it's not a postgres error, log it and return a generic error
	log.Error().Err(err).Msg("Unhandled database error")
	return New(http.StatusInternalServerError, "An unexpected database error occurred")
}
