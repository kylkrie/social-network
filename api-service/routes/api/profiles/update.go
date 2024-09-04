package profiles

import (
	"net/http"

	"yabro.io/social-api/app"
	"yabro.io/social-api/stores/profilestore"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateProfileRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"omitempty,email"`
}

type UpdateProfileResponse struct {
	Profile PublicProfile `json:"profile"`
}

func UpdateProfile(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Error(err)
			return
		}

		var req UpdateProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		params := profilestore.UpdateProfileParams{
			Username: req.Username,
			Email:    req.Email,
		}

		profile, err := appState.Stores.Profile.Update(id, params)
		if err != nil {
			c.Error(err)
			return
		}

		response := UpdateProfileResponse{
			Profile: FromEntity(profile),
		}

		c.JSON(http.StatusOK, response)
	}
}
