package logger

import (
	"log/slog"
	"os"
	"sync"

	"github.com/teacinema-go/core/constants"
)

var (
	log  *slog.Logger
	once sync.Once
)

func Init(env constants.Env) {
	once.Do(func() {
		if env == "" {
			env = constants.Local
		}

		opts := &slog.HandlerOptions{
			Level:     levelByEnv(env),
			AddSource: env != constants.Production,
		}

		var handler slog.Handler
		switch env {
		case constants.Local:
			handler = slog.NewTextHandler(os.Stdout, opts)
		default:
			handler = slog.NewJSONHandler(os.Stdout, opts)
		}

		log = slog.New(handler).With(
			slog.String("env", string(env)),
		)

		slog.SetDefault(log)
	})
}

func levelByEnv(env constants.Env) slog.Level {
	switch env {
	case constants.Production:
		return slog.LevelInfo
	default:
		return slog.LevelDebug
	}
}

func get() *slog.Logger {
	if log == nil {
		return slog.Default()
	}
	return log
}

func Debug(msg string, args ...any) {
	get().Debug(msg, args...)
}

func Info(msg string, args ...any) {
	get().Info(msg, args...)
}

func Warn(msg string, args ...any) {
	get().Warn(msg, args...)
}

func Error(msg string, args ...any) {
	get().Error(msg, args...)
}

func With(args ...any) *slog.Logger {
	return get().With(args...)
}
