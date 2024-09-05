package auth

import (
	"yabro.io/social-api/app"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, appState *app.AppState) {
	apiGroup := router.Group("/auth/v1")
	{
		apiGroup.POST("/token", GetAuthToken(appState))
	}
}
