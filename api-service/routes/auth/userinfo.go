package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserInfo struct {
	Sub               string `json:"sub"`
	PreferredUsername string `json:"preferred_username"`
	Name              string `json:"name"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	Email             string `json:"email"`
}

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		token, exists := c.Get("token")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not found in context"})
			return
		}

		claims, ok := token.(*jwt.Token).Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token claims"})
			return
		}

		userInfo := UserInfo{
			Sub:               userID.(string),
			PreferredUsername: getStringClaim(claims, "preferred_username"),
			Name:              getStringClaim(claims, "name"),
			GivenName:         getStringClaim(claims, "given_name"),
			FamilyName:        getStringClaim(claims, "family_name"),
			Email:             getStringClaim(claims, "email"),
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
