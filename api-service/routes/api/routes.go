package api

import (
	"yabro.io/social-api/app"
	"yabro.io/social-api/auth"
	"yabro.io/social-api/routes/api/profiles"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, appState *app.AppState) {
	apiGroup := router.Group("/api")
	apiGroup.Use(auth.AuthMiddleware(appState.JWKS))
	{
		profiles.SetupRoutes(apiGroup, appState)
	}
}
