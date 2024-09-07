package main

import (
	"context"
	"embed"
	"os"
	"os/signal"
	"syscall"
	"time"

	"yabro.io/social-api/internal/app"
	"yabro.io/social-api/internal/middleware"
	"yabro.io/social-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	appState, err := app.CreateAppState()
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating app state")
	}

	// migrations
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Err(err).Msg("Error setting goose dialect")
	}

	if err := goose.Up(appState.DB.DB, "migrations"); err != nil {
		log.Fatal().Err(err).Msg("Error running migrations")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Setup routes
	routes.SetupRoutes(app, appState)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Start the server in a goroutine
	go func() {
		log.Info().Msgf("Starting server on port %s", port)
		if err := app.Listen(":" + port); err != nil {
			log.Fatal().Err(err).Msg("Error starting server")
		}
	}()

	// Setup channel to listen for termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-quit
	log.Info().Msg("Shutting down gracefully...")

	// Create a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	// Now that the server has shut down, we can close other resources
	appState.Close()

	log.Info().Msg("Server exiting")
}
