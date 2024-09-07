package apperror

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/logger"
)

type ErrorResponse struct {
	Error AppError `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	appErr := ToAppError(err)
	if appErr.Code >= 500 && appErr.Code < 600 {
		logger.Error(c).Int("code", appErr.Code).Str("error", appErr.Message).Msg("AppError")
	} else {
		logger.Warn(c).Int("code", appErr.Code).Str("error", appErr.Message).Msg("AppError")
	}
	return c.Status(appErr.Code).JSON(ErrorResponse{Error: *appErr})
}
