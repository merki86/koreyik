package logger

import (
	"log/slog"
	"os"

	"gitlab.com/greyxor/slogor"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = setupSlogor()
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // Default to production
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupSlogor() *slog.Logger {
	return slog.New(
		slogor.NewHandler(os.Stdout, slogor.Options{
			TimeFormat: "2006-01-02 15:04:05",
			Level:      slog.LevelDebug,
			ShowSource: false,
		}),
	)
}
