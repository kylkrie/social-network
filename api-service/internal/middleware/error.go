package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"yabro.io/social-api/internal/apperror"
)

type ErrorResponse struct {
	Error apperror.AppError `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	appErr := apperror.ToAppError(err)
	if appErr.Code >= 500 && appErr.Code < 600 {
		log.Error().Int("code", appErr.Code).Str("error", appErr.Message).Msg("AppError")
	} else {
		log.Warn().Int("code", appErr.Code).Str("error", appErr.Message).Msg("AppError")
	}
	return c.Status(appErr.Code).JSON(ErrorResponse{Error: *appErr})
}
