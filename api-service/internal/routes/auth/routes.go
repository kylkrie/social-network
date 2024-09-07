package auth

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
)

// SetupRoutes configures the API routes
func SetupRoutes(app *fiber.App, appState *app.AppState) {
	apiGroup := app.Group("/auth/v1")

	apiGroup.Post("/token", GetAuthToken(appState))
}
