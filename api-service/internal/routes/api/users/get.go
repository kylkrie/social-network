package users

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/service"
)

type GetUserRequest struct {
	Username string `params:"username" validate:"required"`
	Profile  bool   `query:"profile"`
}

type GetUserMeRequest struct {
	Profile bool `query:"profile"`
}

type GetUserResponse struct {
	Data service.PublicUser `json:"data"`
}

var validate = validator.New()

func GetUser(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req GetUserRequest

		req.Username = c.Params("username")
		if err := c.QueryParser(&req); err != nil {
			return err
		}
		if err := validate.Struct(req); err != nil {
			return err
		}

		user, err := appState.Services.UserService.GetUserByUsername(req.Username, req.Profile)
		if err != nil {
			return err
		}

		return c.JSON(GetUserResponse{Data: *user})
	}
}

func GetUserMe(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req GetUserMeRequest

		if err := c.QueryParser(&req); err != nil {
			return err
		}
		if err := validate.Struct(req); err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		user, err := appState.Services.UserService.GetUserByID(userID, req.Profile)
		if err != nil {
			return err
		}

		return c.JSON(GetUserResponse{Data: *user})
	}
}
