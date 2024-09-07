package api

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/middleware"
	"yabro.io/social-api/internal/routes/api/posts"
	"yabro.io/social-api/internal/routes/api/users"
)

// SetupRoutes configures the API routes
func SetupRoutes(app *fiber.App, appState *app.AppState) {
	apiGroup := app.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware(appState))

	users.SetupRoutes(apiGroup, appState)
	posts.SetupRoutes(apiGroup, appState)
}
