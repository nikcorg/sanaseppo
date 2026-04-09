package logx

import (
	"log/slog"
	"os"
)

var logLevel = slog.LevelError

func init() {
	if lev, ok := os.LookupEnv("LOG_LEVEL"); ok {
		switch lev {
		case "debug":
			logLevel = slog.LevelDebug
		case "info":
			logLevel = slog.LevelInfo
		}
	}
}

func New() *slog.Logger {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	})

	return slog.New(h)
}
