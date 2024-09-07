package users

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
)

func SetupRoutes(router fiber.Router, appState *app.AppState) {
	profilesGroup := router.Group("/users/v1")
	profilesGroup.Get("", GetUserMe(appState))
	profilesGroup.Put("", UpdateUserMe(appState))
	profilesGroup.Get("/:username", GetUser(appState))
}
