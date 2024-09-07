package users

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
	"yabro.io/social-api/internal/service"
)

type UpdateUserMeRequest struct {
	Name         *string    `json:"name"`
	Protected    *bool      `json:"protected"`
	Bio          *string    `json:"bio"`
	Website      *string    `json:"website"`
	Location     *string    `json:"location"`
	Birthday     *time.Time `json:"birthday"`
	PinnedPostID *int64     `json:"pinned_post_id"`
}

type UpdateUserMeResponse struct {
	Data service.PublicUser `json:"data"`
}

func UpdateUserMe(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateUserMeRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		userID := auth.GetUserID(c)

		updateUser := &service.UpdateUserParams{
			Name:         req.Name,
			Protected:    req.Protected,
			Bio:          req.Bio,
			Website:      req.Website,
			Location:     req.Location,
			Birthday:     req.Birthday,
			PinnedPostID: req.PinnedPostID,
		}

		err := appState.Services.UserService.UpdateUser(userID, updateUser)
		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
