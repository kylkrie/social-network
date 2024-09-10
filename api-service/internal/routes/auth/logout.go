package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/apperror"
)

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func Logout(appState *app.AppState) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LogoutRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := appState.Validator.Struct(req); err != nil {
			return err
		}

		// Construct the logout URL
		logoutURL := appState.AuthConfig.LogoutURL

		// Prepare the form data
		formData := url.Values{
			"client_id":     {appState.AuthConfig.OAuthConfig.ClientID},
			"client_secret": {appState.AuthConfig.OAuthConfig.ClientSecret},
			"refresh_token": {req.RefreshToken},
		}

		// Send POST request to Keycloak to invalidate the session
		resp, err := http.PostForm(logoutURL, formData)
		if err != nil {
			return apperror.New(fiber.StatusInternalServerError, "Failed to send logout request to authentication server")
		}
		defer resp.Body.Close()

		// Check the response status
		if resp.StatusCode != http.StatusNoContent {
			var errorResponse map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
				return apperror.New(fiber.StatusInternalServerError, "Failed to parse authentication server response")
			}
			return apperror.New(resp.StatusCode, fmt.Sprintf("Authentication server error: %v", errorResponse["error_description"]))
		}

		// If we've reached here, the logout was successful
		return c.SendStatus(fiber.StatusNoContent)
	}
}
