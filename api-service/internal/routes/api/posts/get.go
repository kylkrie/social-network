package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/db/postdb"
)

func GetPost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return err
		}

		post, err := appState.Services.PostService.GetPostByID(int64(id), true)
		if err != nil {
			return err
		}

		includes, err := appState.Services.IncludeService.GetIncludesForPost(post)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"data": post, "includes": includes})
	}
}

type ListPostsQuery struct {
	Limit          *int   `query:"limit" validate:"omitempty,min=1,max=100"`
	Cursor         *int64 `query:"cursor" validate:"omitempty"`
	Replies        *bool  `query:"replies" validate:"omitempty"`
	ConversationID *int64 `query:"conversation_id" validate:"omitempty"`
}

func ListPosts(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query ListPostsQuery
		if err := c.QueryParser(&query); err != nil {
			return err
		}

		if err := appState.Validator.Struct(query); err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		limit := 20
		if query.Limit != nil {
			limit = *query.Limit
		}
		isReply := false
		if query.Replies != nil {
			isReply = *query.Replies
		}

		posts, nextCursor, err := appState.Services.PostService.ListPosts(postdb.ListPostParams{
			UserID:         userID,
			Limit:          limit,
			Cursor:         query.Cursor,
			IsReply:        isReply,
			ConversationID: query.ConversationID,
		})
		if err != nil {
			return err
		}

		includes, err := appState.Services.IncludeService.GetIncludesForPosts(posts)
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
