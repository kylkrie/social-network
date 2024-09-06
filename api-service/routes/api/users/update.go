package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"yabro.io/social-api/app"
	"yabro.io/social-api/auth"
	"yabro.io/social-api/service"
)

type UpdateUserMeRequest struct {
	Name         *string    `json:"name"`
	Protected    *bool      `json:"protected"`
	Bio          *string    `json:"bio"`
	Website      *string    `json:"website"`
	Location     *string    `json:"location"`
	Birthday     *time.Time `json:"birthday"`
	PinnedPostID *int64     `json:"pinned_post_id"`
}

type UpdateUserMeResponse struct {
	Data service.PublicUser `json:"data"`
}

func UpdateUserMe(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UpdateUserMeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		userID := auth.GetUserID(c)

		updateUser := &service.UpdateUserParams{
			Name:         req.Name,
			Protected:    req.Protected,
			Bio:          req.Bio,
			Website:      req.Website,
			Location:     req.Location,
			Birthday:     req.Birthday,
			PinnedPostID: req.PinnedPostID,
		}

		err := appState.Services.UserService.UpdateUser(userID, updateUser)
		if err != nil {
			c.Error(err)
			return
		}

		c.Status(http.StatusOK)
	}
}
