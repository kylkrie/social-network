package profiles

import (
	"yabro.io/social-api/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appState *app.AppState) {
	profilesGroup := router.Group("/profiles/v1")
	{
		profilesGroup.POST("", CreateProfile(appState))
		profilesGroup.GET("/:id", GetProfile(appState))
		profilesGroup.PUT("/:id", UpdateProfile(appState))
		profilesGroup.DELETE("/:id", DeleteProfile(appState))
		profilesGroup.GET("", ListProfiles(appState))
	}
}
