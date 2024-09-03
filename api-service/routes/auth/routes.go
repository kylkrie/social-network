package auth

import (
	"yabro.io/social-api/app"
	"yabro.io/social-api/auth"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, appState *app.AppState) {
	apiGroup := router.Group("/auth/v1")
	{
		apiGroup.POST("/token", GetAuthToken(appState))
		apiGroup.GET("/userinfo", auth.AuthMiddleware(appState.JWKS), GetUserInfo())
	}
}
