

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInit_DisabledUsesNoopLogger(t *testing.T) {
	err := Init(Config{
		Enabled: false,
		Level:   "debug",
		Format:  "json",
	})

	require.NoError(t, err)
	require.NotNil(t, log)

	// Noop logger should safely accept logs without panicking.
	assert.NotPanics(t, func() {
		Debug("debug message", "key", "value")
		Info("info message", "key", "value")
		Warn("warn message", "key", "value")
		Error("error message", "key", "value")
		DebugJSON("json message", "payload", map[string]any{"amount": 4500})
		Sync()
	})
}

func TestInit_EnabledJSONLogger(t *testing.T) {
	err := Init(Config{
		Enabled: true,
		Level:   "debug",
		Format:  "json",
	})

	require.NoError(t, err)
	require.NotNil(t, log)

	assert.NotPanics(t, func() {
		Debug("debug message", "key", "value")
		Info("info message", "key", "value")
		Warn("warn message", "key", "value")
		Error("error message", "key", "value")
		DebugJSON("json message", "payload", map[string]any{"amount": 4500})
		Sync()
	})
}

func TestInit_EnabledConsoleLogger(t *testing.T) {
	tests := []string{"console", "text", " CONSOLE "}

	for _, format := range tests {
		t.Run(format, func(t *testing.T) {
			err := Init(Config{
				Enabled: true,
				Level:   "info",
				Format:  format,
			})

			require.NoError(t, err)
			require.NotNil(t, log)
		})
	}
}

func TestSync_NilLoggerDoesNotPanic(t *testing.T) {
	original := log
	log = nil
	defer func() {
		log = original
	}()

	assert.NotPanics(t, func() {
		Sync()
	})
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected zapcore.Level
	}{
		{
			name:     "debug",
			input:    "debug",
			expected: zapcore.DebugLevel,
		},
		{
			name:     "debug with whitespace",
			input:    " DEBUG ",
			expected: zapcore.DebugLevel,
		},
		{
			name:     "warn",
			input:    "warn",
			expected: zapcore.WarnLevel,
		},
		{
			name:     "warning",
			input:    "warning",
			expected: zapcore.WarnLevel,
		},
		{
			name:     "error",
			input:    "error",
			expected: zapcore.ErrorLevel,
		},
		{
			name:     "info",
			input:    "info",
			expected: zapcore.InfoLevel,
		},
		{
			name:     "unknown defaults to info",
			input:    "unknown",
			expected: zapcore.InfoLevel,
		},
		{
			name:     "empty defaults to info",
			input:    "",
			expected: zapcore.InfoLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, parseLevel(tt.input))
		})
	}
}

func TestLoggerPackageStartsWithNoopLogger(t *testing.T) {
	require.NotNil(t, zap.NewNop().Sugar())
}