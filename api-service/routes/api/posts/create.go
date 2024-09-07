package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yabro.io/social-api/app"
	"yabro.io/social-api/auth"
)

type CreatePostRequest struct {
	Content        string `json:"content" binding:"required"`
	ConversationID *int64 `json:"conversation_id"`
}

func CreatePost(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreatePostRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		userID := auth.GetUserID(c)
		post, err := appState.Services.PostService.CreatePost(userID, req.Content, req.ConversationID)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": post})
	}
}
