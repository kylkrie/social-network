package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yabro.io/social-api/app"
	"yabro.io/social-api/auth"
	"yabro.io/social-api/service"
)

type GetUserRequest struct {
	Username string `uri:"username" binding:"required"`
	Profile  bool   `form:"profile"`
}
type GetUserMeRequest struct {
	Profile bool `form:"profile"`
}

type GetUserResponse struct {
	Data service.PublicUser `json:"data"`
}

func GetUser(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetUserRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.Error(err)
			return
		}
		if err := c.ShouldBindQuery(&req); err != nil {
			c.Error(err)
			return
		}

		user, err := appState.Services.UserService.GetUserByUsername(req.Username, req.Profile)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, GetUserResponse{Data: *user})
	}
}

func GetUserMe(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetUserMeRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.Error(err)
			return
		}

		userID := auth.GetUserID(c)
		user, err := appState.Services.UserService.GetUserByID(userID, req.Profile)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, GetUserResponse{Data: *user})
	}
}
