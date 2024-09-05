package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"yabro.io/social-api/app"
)

type UserInfo struct {
	Name       string  `json:"name"`
	Username   string  `json:"username"`
	ProfileUrl *string `json:"profile_url"`
}

func GetUserInfo(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		user, err := appState.Services.UserService.GetUser(userID)
		if err != nil {
			c.Error(err)
			return
		}

		userInfo := UserInfo{
			Name:       user.Name,
			Username:   user.Username,
			ProfileUrl: nil,
		}

		c.JSON(http.StatusOK, userInfo)
	}
}

func getStringClaim(claims jwt.MapClaims, key string) string {
	if value, exists := claims[key]; exists {
		if strValue, ok := value.(string); ok {
			return strValue
		}
	}
	return ""
}
