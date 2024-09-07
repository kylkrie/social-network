package middleware

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/apperror"
)

func AuthMiddleware(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Auth header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return apperror.ToAppError(apperror.ErrUnauthorized)
		}

		// parse out bearer token
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			return apperror.ToAppError(apperror.ErrInvalidToken)
		}

		// validate token with JWKS
		token, err := jwt.Parse(bearerToken[1], appState.JWKS.GetKey)
		if err != nil {
			return apperror.ToAppError(apperror.ErrInvalidToken)
		}

		// parse out claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			sub := claims["sub"].(string)
			authUUID, err := uuid.Parse(sub)
			if err != nil {
				return apperror.ToAppError(apperror.ErrInvalidToken)
			}

			// lookup userID
			userID, err := appState.Services.UserService.GetUserID(authUUID)
			if err != nil {
				// user not found, new user
				if errors.Is(err, sql.ErrNoRows) {
					log.Info().Msg("User not found, creating")
					user, err := appState.Services.UserService.CreateUser(
						authUUID,
						claims["name"].(string),
						claims["preferred_username"].(string),
					)
					if err != nil {
						log.Error().Err(err).Msg("Error creating User")
						return err
					}
					userID = user.ID
					log.Info().Int64("userID", userID).Msg("User created")
				} else {
					return err
				}
			}

			// set ids in context
			c.Locals("authUUID", authUUID)
			c.Locals("userID", userID)
			return c.Next()
		} else {
			log.Info().Msg("Invalid token claims")
			return apperror.ToAppError(apperror.ErrInvalidToken)
		}
	}
}
