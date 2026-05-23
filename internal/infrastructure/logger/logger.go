package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Enabled bool
	Level   string
	Format  string
}

func New(cfg Config) *slog.Logger {
	var output io.Writer = os.Stdout

	if !cfg.Enabled {
		output = io.Discard
	}

	level := parseLevel(cfg.Level)
	var handler slog.Handler

	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(output, &slog.HandlerOptions{
			Level: level,
		})
	default:
		handler = slog.NewTextHandler(output, &slog.HandlerOptions{
			Level: level,
		})
	}

	return slog.New(handler)
}

func SetDefault(l *slog.Logger) {
	slog.SetDefault(l)
}



func parseLevel(level string) slog.Level {

	switch strings.ToLower(strings.TrimSpace(level)) {
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
