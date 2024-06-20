package context

import (
	"log/slog"

	tele "github.com/EgorMizerov/telebot"
)

var ctxKeyLogger = "ctxKeyLogger"

func WithLogger(ctx tele.Context, logger *slog.Logger) {
	ctx.Set(ctxKeyLogger, logger)
}

func Logger(ctx tele.Context) *slog.Logger {
	logger, ok := ctx.Get(ctxKeyLogger).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}
