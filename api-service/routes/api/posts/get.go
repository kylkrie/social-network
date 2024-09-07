package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yabro.io/social-api/app"
	"yabro.io/social-api/auth"
)

func GetPost(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Error(err)
			return
		}

		post, err := appState.Services.PostService.GetPostByID(id, true)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": post})
	}
}

func ListPosts(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		if limit > 100 {
			limit = 100
		}

		var cursor *int64
		if cursorStr := c.Query("cursor"); cursorStr != "" {
			cursorVal, err := strconv.ParseInt(cursorStr, 10, 64)
			if err != nil {
				c.Error(err)
				return
			}
			cursor = &cursorVal
		}

		userID := auth.GetUserID(c)
		posts, nextCursor, err := appState.Services.PostService.ListPosts(userID, limit, cursor)
		if err != nil {
			c.Error(err)
			return
		}

		response := gin.H{
			"data": posts,
		}
		if nextCursor != nil {
			response["next_cursor"] = *nextCursor
		}

		c.JSON(http.StatusOK, response)
	}
}
