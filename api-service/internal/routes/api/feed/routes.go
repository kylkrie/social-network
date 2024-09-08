package feed

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
)

func SetupRoutes(router fiber.Router, appState *app.AppState) {
	postsGroup := router.Group("/feed")

	postsGroup.Get("", ListFeed(appState))
}
