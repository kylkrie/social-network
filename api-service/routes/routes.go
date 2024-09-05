package routes

import (
	"yabro.io/social-api/app"
	"yabro.io/social-api/middleware"
	"yabro.io/social-api/routes/auth"
	"yabro.io/social-api/routes/health"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, appState *app.AppState) {
	router.Use(middleware.CORS(appState.AuthConfig.AllowedOrigin))
	router.Use(middleware.Logger())
	router.Use(middleware.RequestID())
	router.Use(middleware.ErrorHandler())

	health.SetupRoutes(router, appState)
	// api.SetupRoutes(router, appState)
	auth.SetupRoutes(router, appState)
}
