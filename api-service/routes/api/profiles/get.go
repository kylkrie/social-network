package profiles

import (
	"net/http"

	"yabro.io/social-api/app"

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
			c.Error(err)
			return

		}

		profile, err := appState.Stores.Profile.Get(id)
		if err != nil {
			c.Error(err)
			return
		}

		response := GetProfileResponse{
			Profile: FromEntity(profile),
		}

		c.JSON(http.StatusOK, response)
	}
}
