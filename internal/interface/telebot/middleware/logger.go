package middleware

import (
	"log/slog"

	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/context"
	"github.com/EgorMizerov/telebot"
)

func LoggerMiddleware(logger *slog.Logger) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(ctx telebot.Context) error {
			context.WithLogger(ctx, logger)
			return next(ctx)
		}
	}
}
