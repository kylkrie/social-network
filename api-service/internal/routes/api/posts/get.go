package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/db/postdb"
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

		post, err := appState.Services.PostService.GetPostByID(postID, true)
		if err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		includes, err := appState.Services.IncludeService.GetIncludesForPost(post, userID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"data": post, "includes": includes})
	}
}

type ListPostsQuery struct {
	Limit          *int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Cursor         *string `query:"cursor" validate:"omitempty"`
	Replies        *bool   `query:"replies" validate:"omitempty"`
	ConversationID *string `query:"conversation_id" validate:"omitempty"`
	UserID         *string `query:"user_id" validate:"omitempty"`
	Username       *string `query:"username" validate:"omitempty"`
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

		var userID int64
		if query.UserID != nil {
			uID, err := util.StringToInt64(*query.UserID)
			if err != nil {
				return err
			}
			userID = uID
		} else if query.Username != nil {
			user, err := appState.Services.UserService.GetUserByUsername(*query.Username, false)
			if err != nil {
				return err
			}

			userID = util.StringToInt64MustParse(user.ID)
		} else {
			userID = auth.GetUserID(c)
		}
		limit := 10
		if query.Limit != nil {
			limit = *query.Limit
		}
		isReply := false
		if query.Replies != nil {
			isReply = *query.Replies
		}
		conversationID, err := util.NullableStringToInt64(query.ConversationID)
		if err != nil {
			return nil
		}
		cursor, err := util.NullableStringToInt64(query.Cursor)
		if err != nil {
			return err
		}

		posts, nextCursor, err := appState.Services.PostService.ListPosts(postdb.ListPostParams{
			UserID:         &userID,
			Limit:          limit,
			Cursor:         cursor,
			IsReply:        isReply,
			ConversationID: conversationID,
		})
		if err != nil {
			return err
		}

		myUserID := auth.GetUserID(c)
		includes, err := appState.Services.IncludeService.GetIncludesForPosts(posts, myUserID)
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
