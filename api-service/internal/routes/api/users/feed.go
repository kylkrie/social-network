package users

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/db/postdb"
	"yabro.io/social-api/internal/routes/api/posts"
)

type ListFeedQuery struct {
	Limit  *int   `query:"limit" validate:"omitempty,min=1,max=100"`
	Cursor *int64 `query:"cursor" validate:"omitempty"`
}

func ListFeed(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query ListFeedQuery
		if err := c.QueryParser(&query); err != nil {
			return err
		}

		if err := appState.Validator.Struct(query); err != nil {
			return err
		}

		limit := 10
		if query.Limit != nil {
			limit = *query.Limit
		}

		ctx := app.CreateContext(c)
		postDatas, nextCursor, err := appState.Services.PostService.ListPosts(ctx, postdb.ListPostParams{
			Limit:  limit,
			Cursor: query.Cursor,
		})
		if err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		includes, err := appState.Services.IncludeService.GetIncludesForPosts(ctx, postDatas, userID)
		if err != nil {
			return err
		}

		response := posts.ToPostListResponse(postDatas, *includes, nextCursor)
		return c.JSON(response)
	}
}
