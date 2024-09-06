package api

import (
	"yabro.io/social-api/app"
	"yabro.io/social-api/middleware"
	"yabro.io/social-api/routes/api/posts"
	"yabro.io/social-api/routes/api/users"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, appState *app.AppState) {
	apiGroup := router.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware(appState))
	{
		users.SetupRoutes(apiGroup, appState)
		posts.SetupRoutes(apiGroup, appState)
	}
}
