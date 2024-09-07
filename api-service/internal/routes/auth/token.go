package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"yabro.io/social-api/internal/app"
)

func GetAuthToken(state *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request struct {
			Code         string `json:"code"`
			CodeVerifier string `json:"code_verifier"`
			RefreshToken string `json:"refresh_token"`
		}
		if err := c.BodyParser(&request); err != nil {
			return err
		}

		var token *oauth2.Token
		var err error

		if request.Code != "" {
			// Exchange code for token using PKCE
			token, err = state.AuthConfig.OAuthConfig.Exchange(
				context.Background(),
				request.Code,
				oauth2.SetAuthURLParam("code_verifier", request.CodeVerifier),
			)
			if err != nil {
				log.Err(err).Msg("Failed to exchange code for token")
				return err
			}
		} else if request.RefreshToken != "" {
			// Refresh the token
			tokenSource := state.AuthConfig.OAuthConfig.TokenSource(context.Background(), &oauth2.Token{
				RefreshToken: request.RefreshToken,
			})
			token, err = tokenSource.Token()
			if err != nil {
				log.Err(err).Msg("Failed to refresh token")
				return err
			}
		} else {
			return err
		}

		return c.JSON(fiber.Map{
			"access_token":  token.AccessToken,
			"token_type":    token.TokenType,
			"refresh_token": token.RefreshToken,
			"expiry":        token.Expiry,
		})
	}
}
