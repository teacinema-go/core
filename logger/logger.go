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
			Level:     slog.LevelInfo,
			AddSource: true,
		})
	case constants.Staging:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	}

	return slog.New(handler)
}

func WithMethod(logger *slog.Logger, method string) *slog.Logger {
	return logger.With(slog.String("method", method))
}

func WithService(logger *slog.Logger, service string) *slog.Logger {
	return logger.With(slog.String("service", service))
}
