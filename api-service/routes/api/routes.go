package api

import (
	"yabro.io/social-api/app"
	"yabro.io/social-api/middleware"
	"yabro.io/social-api/routes/api/user"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, appState *app.AppState) {
	apiGroup := router.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware(appState))
	{
		user.SetupRoutes(apiGroup, appState)
	}
}
