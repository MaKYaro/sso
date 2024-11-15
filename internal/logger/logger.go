package logger

import (
	"log/slog"
	"os"
)

const (
	local = "local"
	prod  = "prod"
)

func New(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case local:
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdin,
				&slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case prod:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdin,
				&slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}

	return logger
}
