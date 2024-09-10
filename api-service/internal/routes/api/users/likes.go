package users

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/util"
)

type GetUserLikesQuery struct {
	Limit  *int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Cursor *string `query:"cursor" validate:"omitempty"`
}

func GetUserLikes(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query GetUserLikesQuery
		if err := c.QueryParser(&query); err != nil {
			return err
		}

		if err := appState.Validator.Struct(query); err != nil {
			return err
		}

		username := c.Params("username")
		user, err := appState.Services.UserService.GetUserByUsername(username, false)
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

		posts, nextCursor, err := appState.Services.PostService.ListUserLikes(util.StringToInt64MustParse(user.ID), limit, cursor)
		if err != nil {
			return err
		}

		myUserID := auth.GetUserID(c)
		includes, err := appState.Services.IncludeService.GetIncludesForPosts(posts, myUserID)
		if err != nil {
			return err
		}

		response := fiber.Map{
			"data": posts,

			"includes": includes,
		}
		if nextCursor != nil {
			response["next_cursor"] = *nextCursor
		}

		return c.JSON(response)
	}
}
