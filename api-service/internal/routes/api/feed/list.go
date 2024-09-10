package feed

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/db/postdb"
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

		posts, nextCursor, err := appState.Services.PostService.ListPosts(postdb.ListPostParams{
			Limit:  limit,
			Cursor: query.Cursor,
		})
		if err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		includes, err := appState.Services.IncludeService.GetIncludesForPosts(posts, userID)
		if err != nil {
			return err
		}

		response := fiber.Map{
			"data":     posts,
			"includes": includes,
		}
		if nextCursor != nil {
			response["next_cursor"] = *nextCursor
		}

		return c.JSON(response)
	}
}
