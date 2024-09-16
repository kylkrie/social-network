package logger

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func addContext(ctx context.Context, e *zerolog.Event) *zerolog.Event {
	fiberCtx, ok := ctx.Value("fiber_req").(*fiber.Ctx)
	if ok {
		e = e.Str("request_id", fiberCtx.Locals("requestid").(string))
	}

	return e
}

func Info(ctx context.Context) *zerolog.Event {
	return addContext(ctx, log.Info())
}

func Warn(ctx context.Context) *zerolog.Event {
	return addContext(ctx, log.Warn())
}

func Debug(ctx context.Context) *zerolog.Event {
	return addContext(ctx, log.Debug())
}

func Error(ctx context.Context) *zerolog.Event {
	return addContext(ctx, log.Error())
}

func Err(ctx context.Context, err error) *zerolog.Event {
	return addContext(ctx, log.Error()).Err(err)
}
