package middleware

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"yabro.io/social-api/app"
	"yabro.io/social-api/apperror"
	"yabro.io/social-api/service"
)

func AuthMiddleware(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Auth header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(apperror.ToAppError(apperror.ErrUnauthorized))
			c.Abort()
			return
		}

		// parse out bearer token
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.Error(apperror.ToAppError(apperror.ErrInvalidToken))
			c.Abort()
			return
		}

		// validate token with JWKS
		token, err := jwt.Parse(bearerToken[1], appState.JWKS.GetKey)
		if err != nil {
			c.Error(apperror.ToAppError(apperror.ErrInvalidToken))
			c.Abort()
			return
		}

		// parse out claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			sub := claims["sub"].(string)
			authUUID, err := uuid.Parse(sub)
			if err != nil {
				c.Error(apperror.ToAppError(apperror.ErrInvalidToken))
				c.Abort()
				return
			}

			// lookup userID
			userID, err := appState.Services.UserService.GetUserID(authUUID)
			if err != nil {
				// user not found, new user
				if errors.Is(err, sql.ErrNoRows) {
					log.Info().Msg("User not found, creating")
					user, err := appState.Services.UserService.CreateUser(service.CreateUserInput{
						AuthUUID: authUUID,
						Username: claims["preferred_username"].(string),
						Name:     claims["name"].(string),
					})
					if err != nil {
						log.Error().Err(err).Msg("Error creating User")
						c.Error(err)
						c.Abort()
						return
					}

					userID = user.ID
					log.Info().Int64("userID", userID).Msg("User created")
				} else {
					c.Error(err)
					c.Abort()
					return
				}
			}

			// set ids in context
			c.Set("authUUID", authUUID)
			c.Set("userID", userID)
			c.Next()
		} else {
			log.Info().Msg("Invalid token claims")
			c.Error(apperror.ToAppError(apperror.ErrInvalidToken))
			c.Abort()
			return
		}
	}
}
