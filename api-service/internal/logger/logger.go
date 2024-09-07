package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func addRequestID(c *fiber.Ctx, e *zerolog.Event) *zerolog.Event {
	return e.Str("request_id", c.Locals("requestid").(string))
}

func Info(c *fiber.Ctx) *zerolog.Event {
	return addRequestID(c, log.Info())
}

func Warn(c *fiber.Ctx) *zerolog.Event {
	return addRequestID(c, log.Warn())
}

func Debug(c *fiber.Ctx) *zerolog.Event {
	return addRequestID(c, log.Debug())
}

func Error(c *fiber.Ctx) *zerolog.Event {
	return addRequestID(c, log.Error())
}
