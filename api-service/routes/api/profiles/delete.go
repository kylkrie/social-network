package profiles

import (
	"net/http"

	"yabro.io/social-api/app"
	"yabro.io/social-api/apperror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteProfileResponse struct {
	Message string `json:"message"`
}

func DeleteProfile(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.Error(apperror.ToAppError(err))
			return
		}

		err = appState.Stores.Profile.Delete(id)
		if err != nil {
			c.Error(apperror.ToAppError(err))
			return
		}

		response := DeleteProfileResponse{
			Message: "Profile successfully deleted",
		}

		c.JSON(http.StatusOK, response)
	}
}
