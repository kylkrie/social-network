package middleware

import (
	"yabro.io/social-api/apperror"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				appErr := apperror.ToAppError(err.Err)
				if appErr.Code >= 500 && appErr.Code < 600 {
					log.Error().Int("code", appErr.Code).Str("error", appErr.Message).Msg("AppError")
				} else {
					log.Warn().Int("code", appErr.Code).Str("error", appErr.Message).Msg("AppError")
				}
				c.JSON(appErr.Code, appErr)
			}
			c.Abort()
		}
	}
}
