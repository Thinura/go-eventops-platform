package logger

import (
	"encoding/json"
	"log/slog"
)

func Debug(message string, args ...any) {
	slog.Debug(message, args...)
}

func Info(message string, args ...any) {
	slog.Info(message, args...)
}

func Warn(message string, args ...any) {
	slog.Warn(message, args...)
}

func Error(message string, args ...any) {
	slog.Error(message, args...)
}

func DebugJSON(message string, key string, value any) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		slog.Debug(message, slog.Any(key, value))
		return
	}

	slog.Debug(message, slog.String(key, string(data)))
}
