
package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad_DefaultValuesForLocalEnv(t *testing.T) {
	t.Setenv("APP_ENV", "")
	t.Setenv("HTTP_PORT", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("LOG_FORMAT", "")
	t.Setenv("APP_LOGGING", "")

	cfg := Load()

	assert.Equal(t, "local", cfg.AppEnv)
	assert.Equal(t, "8080", cfg.HTTPPort)
	assert.Equal(t, "info", cfg.LogLevel)
	assert.Equal(t, "text", cfg.LogFormat)
	assert.True(t, cfg.AppLogging)
}

func TestLoad_UsesEnvironmentValues(t *testing.T) {
	t.Setenv("APP_ENV", "staging")
	t.Setenv("HTTP_PORT", "9090")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("LOG_FORMAT", "json")
	t.Setenv("APP_LOGGING", "false")

	cfg := Load()

	assert.Equal(t, "staging", cfg.AppEnv)
	assert.Equal(t, "9090", cfg.HTTPPort)
	assert.Equal(t, "debug", cfg.LogLevel)
	assert.Equal(t, "json", cfg.LogFormat)
	assert.False(t, cfg.AppLogging)
}

func TestLoad_TrimsEmptyEnvironmentValues(t *testing.T) {
	t.Setenv("APP_ENV", "   ")
	t.Setenv("HTTP_PORT", "   ")
	t.Setenv("LOG_LEVEL", "   ")
	t.Setenv("LOG_FORMAT", "   ")
	t.Setenv("APP_LOGGING", "   ")

	cfg := Load()

	assert.Equal(t, "local", cfg.AppEnv)
	assert.Equal(t, "8080", cfg.HTTPPort)
	assert.Equal(t, "info", cfg.LogLevel)
	assert.Equal(t, "text", cfg.LogFormat)
	assert.True(t, cfg.AppLogging)
}

func TestLoad_DefaultAppLoggingForProduction(t *testing.T) {
	t.Setenv("APP_ENV", "production")
	t.Setenv("APP_LOGGING", "")

	cfg := Load()

	assert.Equal(t, "production", cfg.AppEnv)
	assert.False(t, cfg.AppLogging)
}

func TestLoad_DefaultAppLoggingForProdAlias(t *testing.T) {
	t.Setenv("APP_ENV", "prod")
	t.Setenv("APP_LOGGING", "")

	cfg := Load()

	assert.Equal(t, "prod", cfg.AppEnv)
	assert.False(t, cfg.AppLogging)
}

func TestLoad_ExplicitAppLoggingOverridesEnvironmentDefault(t *testing.T) {
	t.Setenv("APP_ENV", "production")
	t.Setenv("APP_LOGGING", "true")

	cfg := Load()

	assert.True(t, cfg.AppLogging)
}

func TestGetEnv(t *testing.T) {
	t.Setenv("TEST_CONFIG_VALUE", " actual-value ")
	assert.Equal(t, "actual-value", getEnv("TEST_CONFIG_VALUE", "fallback"))

	t.Setenv("TEST_CONFIG_EMPTY", "   ")
	assert.Equal(t, "fallback", getEnv("TEST_CONFIG_EMPTY", "fallback"))
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "true", input: "true", expected: true},
		{name: "one", input: "1", expected: true},
		{name: "yes", input: "yes", expected: true},
		{name: "y", input: "y", expected: true},
		{name: "false", input: "false", expected: false},
		{name: "zero", input: "0", expected: false},
		{name: "no", input: "no", expected: false},
		{name: "empty", input: "", expected: false},
		{name: "unknown", input: "random", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, parseBool(tt.input))
		})
	}
}

func TestDefaultRequestLogging(t *testing.T) {
	tests := []struct {
		name     string
		appEnv   string
		expected string
	}{
		{name: "local", appEnv: "local", expected: "true"},
		{name: "development", appEnv: "development", expected: "true"},
		{name: "production", appEnv: "production", expected: "false"},
		{name: "prod", appEnv: "prod", expected: "false"},
		{name: "production with whitespace and uppercase", appEnv: " PRODUCTION ", expected: "false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, defaultRequestLogging(tt.appEnv))
		})
	}
}
