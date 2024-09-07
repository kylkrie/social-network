package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yabro.io/social-api/app"
	"yabro.io/social-api/auth"
)

type UpdatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

func UpdatePost(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Error(err)
			return
		}

		var req UpdatePostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		userID := auth.GetUserID(c)
		post, err := appState.Services.PostService.UpdatePost(id, userID, req.Content)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": post})
	}
}
