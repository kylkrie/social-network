package posts

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/service"
	"yabro.io/social-api/internal/util"
)

func ListRepliesForPost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID, err := util.StringToInt64(c.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid post ID")
		}

		limit := 10 // Default limit
		if limitStr := c.Query("limit"); limitStr != "" {
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid limit parameter")
			}
			if limit < 1 || limit > 100 {
				return fiber.NewError(fiber.StatusBadRequest, "Limit must be between 1 and 100")
			}
		}

		var cursor *int64
		if cursorStr := c.Query("cursor"); cursorStr != "" {
			cursorVal, err := util.StringToInt64(cursorStr)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid cursor parameter")
			}
			cursor = &cursorVal
		}

		ctx := app.CreateContext(c)
		replies, nextCursor, err := appState.Services.PostService.ListRepliesForPost(ctx, service.ListRepliesForPostParams{
			PostID: postID,
			Limit:  limit,
			Cursor: cursor,
		})
		if err != nil {
			return err
		}

		// Fetch the authenticated user's ID for including user interactions
		userID := auth.GetUserID(c)

		// Fetch additional data for includes
		includes, err := appState.Services.IncludeService.GetIncludesForPosts(ctx, replies, &userID)
		if err != nil {
			return err
		}

		response := ToPostListResponse(replies, *includes, nextCursor)
		return c.JSON(response)
	}
}
