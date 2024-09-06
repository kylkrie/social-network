package users

import (
	"yabro.io/social-api/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appState *app.AppState) {
	profilesGroup := router.Group("/users/v1")
	{
		profilesGroup.GET("", GetUserMe(appState))
		profilesGroup.PUT("", UpdateUserMe(appState))
		profilesGroup.GET("/:username", GetUser(appState))
	}
}
