package middleware

import (
	"net/http"

	"yabro.io/social-api/apperror"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				switch e := err.Err.(type) {
				case *apperror.AppError:
					log.Warn().Int("code", e.Code).Str("error", e.Message).Msg("AppError")
					c.JSON(e.Code, gin.H{"error": e.Message})
				default:
					log.Error().Err(e).Msg("Unexpected error")
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				}
			}
			c.Abort()
		}
	}
}
