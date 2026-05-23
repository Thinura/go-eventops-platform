package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Thinura/go-eventops-platform/internal/config"
	"github.com/Thinura/go-eventops-platform/internal/infrastructure/logger"
	httptransport "github.com/Thinura/go-eventops-platform/internal/transport/http"
)

type serveFunc func(server *http.Server) error

func main() {
	os.Exit(run(context.Background(), config.Load(), defaultServe, systemSignalChannel()))
}

func run(ctx context.Context, cfg config.Config, serve serveFunc, shutdown <-chan os.Signal) int {
	if err := logger.Init(logger.Config{
		Enabled: cfg.AppLogging,
		Level:   cfg.LogLevel,
		Format:  cfg.LogFormat,
	}); err != nil {
		return 1
	}
	defer logger.Sync()

	router := httptransport.NewRouter(httptransport.RouterConfig{
		AppLogging: cfg.AppLogging,
	})

	server := newServer(cfg, router)
	serverErrors := make(chan error, 1)

	go func() {
		logger.Info("starting api server", "port", cfg.HTTPPort)
		if err := serve(server); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case err := <-serverErrors:
		logger.Error("server error", "error", err)
		return 1
	case sig := <-shutdown:
		logger.Info("shutdown signal received", "signal", sig)
		return shutdownServer(server)
	case <-ctx.Done():
		logger.Info("context cancelled, shutting down api server")
		return shutdownServer(server)
	}
}

func newServer(cfg config.Config, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func shutdownServer(server *http.Server) int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("server shutdown error", "error", err)
		return 1
	}

	logger.Info("server gracefully stopped")
	return 0
}

func defaultServe(server *http.Server) error {
	return server.ListenAndServe()
}

func systemSignalChannel() <-chan os.Signal {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	return shutdown
}
