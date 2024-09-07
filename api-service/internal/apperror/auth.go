package apperror

import (
	"net/http"
)

// AuthErrorCode is an enumeration of authentication error codes
type AuthErrorCode int

const (
	ErrUnauthorized AuthErrorCode = iota
	ErrInvalidToken
	ErrTokenExpired
	ErrInvalidCredentials
	ErrTokenExchangeFailed
	ErrTokenRefreshFailed
)

// AuthError represents an authentication error
type AuthError struct {
	Code    AuthErrorCode
	Message string
	Status  int
}

func (e AuthError) Error() string {
	return e.Message
}

var authErrors = map[AuthErrorCode]AuthError{
	ErrUnauthorized:        {Code: ErrUnauthorized, Message: "Unauthorized", Status: http.StatusUnauthorized},
	ErrInvalidToken:        {Code: ErrInvalidToken, Message: "Invalid token", Status: http.StatusUnauthorized},
	ErrTokenExpired:        {Code: ErrTokenExpired, Message: "Token expired", Status: http.StatusUnauthorized},
	ErrInvalidCredentials:  {Code: ErrInvalidCredentials, Message: "Invalid credentials", Status: http.StatusUnauthorized},
	ErrTokenExchangeFailed: {Code: ErrTokenExchangeFailed, Message: "Failed to exchange code for token", Status: http.StatusInternalServerError},
	ErrTokenRefreshFailed:  {Code: ErrTokenRefreshFailed, Message: "Failed to refresh token", Status: http.StatusInternalServerError},
}

func newAuthError(code AuthErrorCode) AuthError {
	if err, ok := authErrors[code]; ok {
		return err
	}
	return AuthError{Code: code, Message: "Unknown auth error", Status: http.StatusInternalServerError}
}

func HandleAuthError(code AuthErrorCode) *AppError {
	authErr := newAuthError(code)
	return New(authErr.Status, authErr.Message)
}
