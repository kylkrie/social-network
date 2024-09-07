package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// RequestLogger middleware to log each request
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		requestID := c.Locals("requestid").(string)

		// Log request received
		log.Info().
			Str("request_id", requestID).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("ip", c.IP()).
			Msg("in")

		// Process the request
		err := c.Next()

		// Log request completed
		log.Info().
			Str("request_id", requestID).
			Dur("latency", time.Since(start)).
			Int("status", c.Response().StatusCode()).
			Msg("out")

		return err
	}
}
