package profiles

import (
	"net/http"

	"yabro.io/social-api/app"
	"yabro.io/social-api/apperror"
	"yabro.io/social-api/auth"
	"yabro.io/social-api/stores/profilestore"

	"github.com/gin-gonic/gin"
)

type CreateProfileRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type CreateProfileResponse struct {
	Profile PublicProfile `json:"profile"`
}

func CreateProfile(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(apperror.ToAppError(err))
			return
		}

		userID, err := auth.GetUserId(c)
		if err != nil {
			c.Error(apperror.ToAppError(err))
			return
		}

		profile, err := appState.Stores.Profile.Create(profilestore.CreateProfileParams{
			UserID:   userID.String(),
			Username: req.Username,
			Email:    req.Email,
		})
		if err != nil {
			c.Error(apperror.ToAppError(err))
			return
		}

		response := CreateProfileResponse{
			Profile: FromEntity(profile),
		}

		c.JSON(http.StatusCreated, response)
	}
}
