package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServer_UsesConfigAndRouter(t *testing.T) {
	router := http.NewServeMux()
	cfg := testConfig()
	cfg.HTTPPort = "9090"

	server := newServer(cfg, router)

	require.NotNil(t, server)
	assert.Equal(t, ":9090", server.Addr)
	assert.Same(t, router, server.Handler)
	assert.Equal(t, 5*time.Second, server.ReadTimeout)
	assert.Equal(t, 10*time.Second, server.WriteTimeout)
	assert.Equal(t, 60*time.Second, server.IdleTimeout)
}

func TestRun_ReturnsOneWhenServerFails(t *testing.T) {
	shutdown := make(chan os.Signal)
	serveErr := errors.New("listen failed")

	code := run(context.Background(), testConfig(), func(server *http.Server) error {
		return serveErr
	}, shutdown)

	assert.Equal(t, 1, code)
}

func TestRun_ReturnsZeroOnShutdownSignal(t *testing.T) {
	shutdown := make(chan os.Signal, 1)
	shutdown <- os.Interrupt

	code := run(context.Background(), testConfig(), func(server *http.Server) error {
		return http.ErrServerClosed
	}, shutdown)

	assert.Equal(t, 0, code)
}

func TestRun_ReturnsZeroWhenContextIsCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	shutdown := make(chan os.Signal)

	code := run(ctx, testConfig(), func(server *http.Server) error {
		return http.ErrServerClosed
	}, shutdown)

	assert.Equal(t, 0, code)
}

func TestShutdownServer_ReturnsZeroForNewServer(t *testing.T) {
	server := &http.Server{}

	code := shutdownServer(server)

	assert.Equal(t, 0, code)
}

func TestSystemSignalChannel(t *testing.T) {
	shutdown := systemSignalChannel()

	require.NotNil(t, shutdown)
}

func testConfig() config.Config {
	return config.Config{
		AppEnv:     "test",
		HTTPPort:  "0",
		LogLevel:  "error",
		LogFormat: "json",
		AppLogging: false,
	}
}
