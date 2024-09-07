package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
)

type UpdatePostRequest struct {
	Content string `json:"content" validate:"required"`
}

func UpdatePost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return err
		}

		var req UpdatePostRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := appState.Validator.Struct(req); err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		post, err := appState.Services.PostService.UpdatePost(int64(id), userID, req.Content)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"data": post})
	}
}
