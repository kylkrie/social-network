package posts

import (
	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
)

type CreatePostRequest struct {
	Content        string `json:"content" validate:"required"`
	ConversationID *int64 `json:"conversation_id"`
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

		userID := auth.GetUserID(c)
		post, err := appState.Services.PostService.CreatePost(userID, req.Content, req.ConversationID)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": post})
	}
}
