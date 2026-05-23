
package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogHelpers_WriteExpectedLevelsAndFields(t *testing.T) {
	observedLogs := useObservedLogger(t, zapcore.DebugLevel)

	Debug("debug message", "key", "debug-value")
	Info("info message", "key", "info-value")
	Warn("warn message", "key", "warn-value")
	Error("error message", "key", "error-value")

	entries := observedLogs.All()
	require.Len(t, entries, 4)

	assert.Equal(t, zapcore.DebugLevel, entries[0].Level)
	assert.Equal(t, "debug message", entries[0].Message)
	assert.Equal(t, "debug-value", entries[0].ContextMap()["key"])

	assert.Equal(t, zapcore.InfoLevel, entries[1].Level)
	assert.Equal(t, "info message", entries[1].Message)
	assert.Equal(t, "info-value", entries[1].ContextMap()["key"])

	assert.Equal(t, zapcore.WarnLevel, entries[2].Level)
	assert.Equal(t, "warn message", entries[2].Message)
	assert.Equal(t, "warn-value", entries[2].ContextMap()["key"])

	assert.Equal(t, zapcore.ErrorLevel, entries[3].Level)
	assert.Equal(t, "error message", entries[3].Message)
	assert.Equal(t, "error-value", entries[3].ContextMap()["key"])
}

func TestDebugJSON_LogsPrettyJSONString(t *testing.T) {
	observedLogs := useObservedLogger(t, zapcore.DebugLevel)

	DebugJSON("payload received", "payload", map[string]any{
		"amount": float64(4500),
		"reason": "card_declined",
	})

	entries := observedLogs.All()
	require.Len(t, entries, 1)

	payload, ok := entries[0].ContextMap()["payload"].(string)
	require.True(t, ok)

	assert.Contains(t, payload, "\n")
	assert.Contains(t, payload, `"amount": 4500`)
	assert.Contains(t, payload, `"reason": "card_declined"`)
}

func TestDebugJSON_WhenMarshalFails_LogsOriginalValue(t *testing.T) {
	observedLogs := useObservedLogger(t, zapcore.DebugLevel)

	value := map[string]any{
		"invalid": func() {},
	}

	assert.NotPanics(t, func() {
		DebugJSON("payload received", "payload", value)
	})

	entries := observedLogs.All()
	require.Len(t, entries, 1)

	assert.Equal(t, "payload received", entries[0].Message)
	assert.Equal(t, value, entries[0].ContextMap()["payload"])
}

func TestDebugLogsAreFilteredByLevel(t *testing.T) {
	observedLogs := useObservedLogger(t, zapcore.InfoLevel)

	Debug("debug message", "key", "value")
	Info("info message", "key", "value")

	entries := observedLogs.All()
	require.Len(t, entries, 1)
	assert.Equal(t, zapcore.InfoLevel, entries[0].Level)
	assert.Equal(t, "info message", entries[0].Message)
}

func useObservedLogger(t *testing.T, level zapcore.Level) *observer.ObservedLogs {
	t.Helper()

	original := log
	core, observedLogs := observer.New(level)
	log = zap.New(core).Sugar()

	t.Cleanup(func() {
		log = original
	})

	return observedLogs
}
