package public

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
)

// SetupRoutes configures the API routes
func SetupRoutes(app *fiber.App, appState *app.AppState) {
	publicGroup := app.Group("/public/v1")

	publicGroup.Get("/feed", ListFeed(appState))
	publicGroup.Get("/posts/:id", GetPost(appState))
}
