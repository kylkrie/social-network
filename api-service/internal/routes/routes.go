package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/middleware"
	"yabro.io/social-api/internal/routes/api"
	"yabro.io/social-api/internal/routes/auth"
	"yabro.io/social-api/internal/routes/health"
	"yabro.io/social-api/internal/routes/public"
)

func SetupRoutes(app *fiber.App, appState *app.AppState) {
	app.Use(middleware.CORS(appState.AuthConfig.AllowedOrigin))
	app.Use(requestid.New())
	app.Use(middleware.RequestLogger())

	health.SetupRoutes(app, appState)
	api.SetupRoutes(app, appState)
	auth.SetupRoutes(app, appState)
	public.SetupRoutes(app, appState)
}
