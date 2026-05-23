package logger

import (
	"encoding/json"
)

func Debug(message string, args ...any) {
	log.Debugw(message, args...)
}

func Info(message string, args ...any) {
	log.Infow(message, args...)
}

func Warn(message string, args ...any) {
	log.Warnw(message, args...)
}

func Error(message string, args ...any) {
	log.Errorw(message, args...)
}

func DebugJSON(message string, key string, value any) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		log.Debugw(message, key, value)
		return
	}

	log.Debugw(message, key, string(data))
}
