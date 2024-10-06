package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/service"
	"yabro.io/social-api/internal/util"
)

type GetPostRequest struct {
	ID string `params:"id" validate:"required"`
}

func GetPost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req GetPostRequest
		req.ID = c.Params("id")
		if err := c.QueryParser(&req); err != nil {
			return err
		}
		if err := appState.Validator.Struct(req); err != nil {
			return err
		}
		postID, err := util.StringToInt64(req.ID)
		if err != nil {
			return err
		}

		ctx := app.CreateContext(c)
		post, err := appState.Services.PostService.GetPostByID(ctx, postID, true, true)
		if err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		posts := []service.PostData{*post}
		includes, err := appState.Services.IncludeService.GetIncludesForPosts(ctx, posts, &userID)
		if err != nil {
			return err
		}

		return c.JSON(ToPostResponse(*post, *includes))
	}
}
