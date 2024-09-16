package users

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
)

func SetupRoutes(router fiber.Router, appState *app.AppState) {
	profilesGroup := router.Group("/users")
	profilesGroup.Get("/me", GetUserMe(appState))
	profilesGroup.Put("/me", UpdateUserMe(appState))
	profilesGroup.Post("/me/pfp", UploadProfilePicture(appState))
	profilesGroup.Post("/me/pfbanner", UploadProfileBanner(appState))
	profilesGroup.Get("/me/feed", ListFeed(appState))
	profilesGroup.Get("/:username", GetUser(appState))
	profilesGroup.Get("/:username/likes", GetUserLikes(appState))
	profilesGroup.Get("/:username/bookmarks", GetUserBookmarks(appState))
	profilesGroup.Get("/:username/posts", ListPosts(appState))
}
