package config

import (
	"log/slog"
	"os"
	"strings"
)

// LoadLogger creates a custom app-wide logger
func LoadLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     LoadLogLevel(),
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	return logger
}

func LoadLogLevel() slog.Level {
	logLevel := strings.TrimSpace(strings.ToLower(os.Getenv("LOG_LEVEL")))

	switch logLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
