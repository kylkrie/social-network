package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/util"
)

func BookmarkPost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID, err := util.StringToInt64(c.Params("id"))
		if err != nil {
			return err
		}

		userID := auth.GetUserID(c)

		ctx := app.CreateContext(c)
		err = appState.Services.PostService.BookmarkPost(ctx, postID, userID)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func UnbookmarkPost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID, err := util.StringToInt64(c.Params("id"))
		if err != nil {
			return err
		}

		userID := auth.GetUserID(c)

		ctx := app.CreateContext(c)
		err = appState.Services.PostService.UnbookmarkPost(ctx, postID, userID)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
