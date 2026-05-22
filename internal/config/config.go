package config

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	HTTPPort string
	LogLevel slog.Level
}

func Load() Config {
	return Config{
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		LogLevel: parseLogLevel(getEnv("LOG_LEVEL", "info")),
	}
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
