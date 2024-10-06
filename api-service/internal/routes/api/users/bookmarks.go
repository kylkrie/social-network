package users

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/routes/api/posts"
	"yabro.io/social-api/internal/util"
)

type GetUserBookmarksQuery struct {
	Limit  *int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Cursor *string `query:"cursor" validate:"omitempty"`
}

func GetUserBookmarks(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query GetUserBookmarksQuery
		if err := c.QueryParser(&query); err != nil {
			return err
		}

		if err := appState.Validator.Struct(query); err != nil {
			return err
		}

		username := c.Params("username")
		userData, err := appState.Services.UserService.GetUserByUsername(username, false)
		if err != nil {
			return err
		}

		limit := 10
		if query.Limit != nil {
			limit = *query.Limit
		}

		cursor, err := util.NullableStringToInt64(query.Cursor)
		if err != nil {
			return err
		}

		ctx := app.CreateContext(c)
		postDatas, nextCursor, err := appState.Services.PostService.ListUserBookmarks(
			ctx,
			userData.User.ID,
			limit,
			cursor,
		)
		if err != nil {
			return err
		}

		myUserID := auth.GetUserID(c)
		includes, err := appState.Services.IncludeService.GetIncludesForPosts(ctx, postDatas, &myUserID)
		if err != nil {
			return err
		}

		response := posts.ToPostListResponse(postDatas, *includes, nextCursor)
		return c.JSON(response)
	}
}
