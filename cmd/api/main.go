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
	httptransport "github.com/Thinura/go-eventops-platform/internal/transport/http"
	"github.com/Thinura/go-eventops-platform/internal/infrastructure/logger"
)

func main() {
	cfg := config.Load()

	if err := logger.Init(logger.Config{
		Enabled: cfg.AppLogging,
		Level:   cfg.LogLevel,
		Format:  cfg.LogFormat,
	}); err != nil {
		panic(err)
	}
	defer logger.Sync()

	router := httptransport.NewRouter(
		httptransport.RouterConfig{
		AppLogging: cfg.AppLogging,
		},
	)

	server := &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		logger.Info("starting api server", "port", cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		logger.Error("server error", "error", err)
		os.Exit(1)
	case sig := <-shutdown:
		logger.Info("shutdown signal received", "signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("server shutdown error", "error", err)
			os.Exit(1)
		}
		logger.Info("server gracefully stopped")
	}
}
