package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yabro.io/social-api/app"
	"yabro.io/social-api/services"
)

func ExchangeCode(state *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Code string `json:"code" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		token, err := services.ExchangeCodeForToken(c, state.Config.OAuthConfig, request.Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
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
