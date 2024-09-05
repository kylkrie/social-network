package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"yabro.io/social-api/app"
)

func GetAuthToken(state *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Code         string `json:"code"`
			CodeVerifier string `json:"code_verifier"`
			RefreshToken string `json:"refresh_token"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
				return
			}
		} else if request.RefreshToken != "" {
			// Refresh the token
			tokenSource := state.AuthConfig.OAuthConfig.TokenSource(context.Background(), &oauth2.Token{
				RefreshToken: request.RefreshToken,
			})
			token, err = tokenSource.Token()
			if err != nil {
				log.Err(err).Msg("Failed to refresh token")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Either code or refresh_token must be provided"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token":  token.AccessToken,
			"token_type":    token.TokenType,
			"refresh_token": token.RefreshToken,
			"expiry":        token.Expiry,
		})
	}
}
