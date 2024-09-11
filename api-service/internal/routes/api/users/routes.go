package users

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
)

func SetupRoutes(router fiber.Router, appState *app.AppState) {
	profilesGroup := router.Group("/users")
	profilesGroup.Get("", GetUserMe(appState))
	profilesGroup.Put("", UpdateUserMe(appState))
	profilesGroup.Get("/:username", GetUser(appState))
	profilesGroup.Get("/:username/likes", GetUserLikes(appState))
	profilesGroup.Get("/:username/bookmarks", GetUserBookmarks(appState))
	profilesGroup.Post("/pfp", UploadProfilePicture(appState))
	profilesGroup.Post("/pfbanner", UploadProfileBanner(appState))
}
