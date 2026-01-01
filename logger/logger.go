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

	switch env {
	case constants.Production:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	case constants.Staging:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	return slog.New(handler)
}
