package apperror

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ErrResourceNotFound = "Resource not found"
	ErrResourceExists   = "Resource already exists"
	ErrDatabaseTimeout  = "Database operation timed out"
)

func HandleDBError(err error) *AppError {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return New(http.StatusNotFound, ErrResourceNotFound)
	case errors.Is(err, context.DeadlineExceeded):
		return New(http.StatusGatewayTimeout, ErrDatabaseTimeout)
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation
			return New(http.StatusConflict, ErrResourceExists)
		}
	}

	return nil
}
