package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yabro.io/social-api/app"
	"yabro.io/social-api/auth"
)

func DeletePost(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Error(err)
			return
		}

		userID := auth.GetUserID(c)
		err = appState.Services.PostService.DeletePost(id, userID)
		if err != nil {
			c.Error(err)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
