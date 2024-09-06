package posts

import (
	"yabro.io/social-api/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appState *app.AppState) {
	postsGroup := router.Group("/posts/v1")
	{
		postsGroup.POST("", CreatePost(appState))
		postsGroup.GET("", ListPosts(appState))
		postsGroup.GET("/:id", GetPost(appState))
		postsGroup.PUT("/:id", UpdatePost(appState))
		postsGroup.DELETE("/:id", DeletePost(appState))
	}
}
