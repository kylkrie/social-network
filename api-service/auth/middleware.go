// ./auth/middleware.go
package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"yabro.io/social-api/apperror"
)

func AuthMiddleware(jwks *JWKS) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(apperror.ToAppError(apperror.ErrUnauthorized))
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.Error(apperror.ToAppError(apperror.ErrInvalidToken))
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearerToken[1], jwks.getKey)
		if err != nil {
			c.Error(apperror.ToAppError(apperror.ErrInvalidToken))
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["sub"].(string)
			c.Set("userID", userID)
			c.Set("token", token)
			c.Next()
		} else {
			log.Info().Msg("Invalid token claims")
			c.Error(apperror.ToAppError(apperror.ErrInvalidToken))
			c.Abort()
			return
		}
	}
}
