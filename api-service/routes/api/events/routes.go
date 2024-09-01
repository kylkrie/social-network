package events

import (
	"yabro.io/social-api/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, appState *app.AppState) {
	apiGroup := router.Group("/events/v1")
	{
		apiGroup.GET("", ListEvents(appState))
	}
}
