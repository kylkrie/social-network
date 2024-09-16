package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
)

func DeletePost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		ctx := app.CreateContext(c)
		err = appState.Services.PostService.DeletePost(ctx, int64(id), userID)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
