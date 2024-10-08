package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
)

func SetupRoutes(router fiber.Router, appState *app.AppState) {
	postsGroup := router.Group("/posts")

	postsGroup.Post("", CreatePost(appState))
	postsGroup.Get("/:id", GetPost(appState))
	postsGroup.Delete("/:id", DeletePost(appState))
	postsGroup.Get("/:id/replies", ListRepliesForPost(appState))
	postsGroup.Post("/:id/likes", LikePost(appState))
	postsGroup.Delete("/:id/likes", UnlikePost(appState))
	postsGroup.Post("/:id/bookmarks", BookmarkPost(appState))
	postsGroup.Delete("/:id/bookmarks", UnbookmarkPost(appState))
}
