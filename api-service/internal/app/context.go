package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type contextKey string

const fiberContextKey contextKey = "fiber_ctx"

func CreateContext(c *fiber.Ctx) context.Context {
	return context.WithValue(c.Context(), fiberContextKey, c)
}
