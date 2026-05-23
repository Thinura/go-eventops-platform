package config

import (
	"os"
	"strings"
)

type Config struct {
	AppEnv   string
	HTTPPort string
	LogLevel string
	LogFormat string
	AppLogging bool
}

func Load() Config {
	appEnv := getEnv("APP_ENV", "local")
	return Config{
		AppEnv: appEnv,
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
		LogFormat: getEnv("LOG_FORMAT", "text"),
		AppLogging: parseBool(getEnv("APP_LOGGING", defaultRequestLogging(appEnv))),
	}
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func parseBool(value string) bool {
	switch strings.ToLower(value) {
	case "true", "1", "yes", "y":
		return true
	default:
		return false
	}
}

func defaultRequestLogging(appEnv string) string {
	switch strings.ToLower(strings.TrimSpace(appEnv)) {
	case "production", "prod":
		return "false"
	default:
		return "true"
	}
}
