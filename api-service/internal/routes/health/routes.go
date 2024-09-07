// api/health/routes.go
package health

import (
	"context"

	"yabro.io/social-api/internal/app"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// SetupHealthRoutes registers health-related routes
func SetupRoutes(app *fiber.App, appState *app.AppState) {
	app.Get("/health", func(c *fiber.Ctx) error {
		if err := healthCheck(appState); err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"status": "error", "message": "Service Unhealthy", "error": err.Error()})
		}
		return c.Status(fiber.StatusOK).SendString("OK")
	})
}

// healthCheck performs the actual health checks for your services
func healthCheck(appState *app.AppState) error {
	// Database check
	if err := dbHealthCheck(appState.DB); err != nil {
		return err
	}

	return nil
}

// dbHealthCheck checks the health of the database connection
func dbHealthCheck(db *sqlx.DB) error {
	if err := db.PingContext(context.Background()); err != nil {
		return err
	}
	return nil
}
