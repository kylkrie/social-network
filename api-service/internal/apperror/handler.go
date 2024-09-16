package apperror

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/logger"
)

type ErrorResponse struct {
	Error AppError `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	ctx := context.WithValue(c.Context(), "fiber_req", c)
	appErr := ToAppError(err)
	if appErr.Code >= 500 && appErr.Code < 600 {
		logger.Error(ctx).Int("code", appErr.Code).Str("error", appErr.Message).Msg("AppError")
	} else {
		logger.Warn(ctx).Int("code", appErr.Code).Str("error", appErr.Message).Msg("AppError")
	}
	return c.Status(appErr.Code).JSON(ErrorResponse{Error: *appErr})
}
