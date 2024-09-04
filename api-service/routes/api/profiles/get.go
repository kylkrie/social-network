package profiles

import (
	"net/http"

	"yabro.io/social-api/app"
	"yabro.io/social-api/apperror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetProfileResponse struct {
	Profile PublicProfile `json:"profile"`
}

func GetProfile(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Error(apperror.ToAppError(err))
			return

		}

		profile, err := appState.Stores.Profile.Get(id)
		if err != nil {
			c.Error(apperror.ToAppError(err))
			return
		}

		response := GetProfileResponse{
			Profile: FromEntity(profile),
		}

		c.JSON(http.StatusOK, response)
	}
}
