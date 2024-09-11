package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/auth"
)

func UploadProfilePicture(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := auth.GetUserID(c)

		// Get the file from the request
		file, err := c.FormFile("pfp")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Failed to get file from request")
		}

		// Check file size (e.g., max 5MB)
		if file.Size > 5*1024*1024 {
			return fiber.NewError(fiber.StatusBadRequest, "File size exceeds maximum limit of 5MB")
		}

		// Check file type
		if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
			return fiber.NewError(fiber.StatusBadRequest, "Only JPEG and PNG files are allowed")
		}

		// Open the file
		src, err := file.Open()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to open uploaded file")
		}
		defer src.Close()

		// Upload the file
		err = appState.Services.UserService.UploadProfilePicture(userID, src, file.Size, file.Filename)
		if err != nil {
			log.Error().Err(err).Msg("pfp")
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload profile picture")
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func UploadProfileBanner(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := auth.GetUserID(c)

		// Get the file from the request
		file, err := c.FormFile("banner")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Failed to get file from request")
		}

		// Check file size (e.g., max 5MB)
		if file.Size > 5*1024*1024 {
			return fiber.NewError(fiber.StatusBadRequest, "File size exceeds maximum limit of 5MB")
		}

		// Check file type
		if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
			return fiber.NewError(fiber.StatusBadRequest, "Only JPEG and PNG files are allowed")
		}

		// Open the file
		src, err := file.Open()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to open uploaded file")
		}
		defer src.Close()

		// Upload the file
		err = appState.Services.UserService.UploadProfileBanner(userID, src, file.Size, file.Filename)
		if err != nil {
			log.Error().Err(err).Msg("banner")
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload profile banner")
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}
