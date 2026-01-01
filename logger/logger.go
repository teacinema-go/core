package logger

import (
	"log/slog"
	"os"

	"github.com/teacinema-go/core/constants"
)

func New(env constants.Env) *slog.Logger {
	if env == "" {
		env = constants.Development
	}
	var handler slog.Handler

	if env.IsProduction() {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else if env.IsStaging() {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	return slog.New(handler)
}
