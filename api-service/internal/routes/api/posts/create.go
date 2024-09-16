package posts

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/service"
	"yabro.io/social-api/internal/util"
)

func CreatePost(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := auth.GetUserID(c)

		content := c.FormValue("content")
		var replyToPostID *int64
		replyToPostIDStr := c.FormValue("reply_to_post_id")
		if replyToPostIDStr != "" {
			var err error
			replyToPostID, err = util.NullableStringToInt64(&replyToPostIDStr)
			if err != nil {
				return err
			}
		}

		var quotePostID *int64
		quotePostIDStr := c.FormValue("quote_post_id")
		if quotePostIDStr != "" {
			var err error
			quotePostID, err = util.NullableStringToInt64(&quotePostIDStr)
			if err != nil {
				return err
			}
		}

		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		var media []*multipart.FileHeader
		if mediaFiles := form.File["media"]; len(mediaFiles) > 0 {
			media = mediaFiles
		}

		ctx := app.CreateContext(c)
		err = appState.Services.PostService.CreatePost(ctx, service.CreatePostParams{
			UserID:        userID,
			Content:       content,
			ReplyToPostID: replyToPostID,
			QuotePostID:   quotePostID,
			Media:         media,
		})
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}
