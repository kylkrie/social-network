package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/service"
	"yabro.io/social-api/internal/util"
)

type CreatePostRequest struct {
	Content       string  `json:"content" validate:"required"`
	ReplyToPostID *string `json:"reply_to_post_id"`
	QuotePostID   *string `json:"quote_post_id"`
}

func CreatePost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreatePostRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := appState.Validator.Struct(req); err != nil {
			return err
		}

		replyToPostID, err := util.NullableStringToInt64(req.ReplyToPostID)
		if err != nil {
			return err
		}
		quotePostID, err := util.NullableStringToInt64(req.QuotePostID)
		if err != nil {
			return err
		}

		userID := auth.GetUserID(c)
		post, err := appState.Services.PostService.CreatePost(service.CreatePostParams{
			UserID:        userID,
			Content:       req.Content,
			ReplyToPostID: replyToPostID,
			QuotePostID:   quotePostID,
		})
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": post})
	}
}
