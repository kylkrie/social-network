package posts

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
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

func ListPosts(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit := c.QueryInt("limit", 20)
		if limit > 100 {
			limit = 100
		}

		var cursor *int64
		if cursorStr := c.Query("cursor"); cursorStr != "" {
			cursorVal, err := strconv.ParseInt(cursorStr, 10, 64)
			if err != nil {
				return err
			}
			cursor = &cursorVal
		}

		userID := auth.GetUserID(c)
		posts, nextCursor, err := appState.Services.PostService.ListPosts(userID, limit, cursor)
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
