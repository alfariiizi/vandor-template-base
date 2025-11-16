package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"

	"{{.ModuleName}}/cmd/app/config"
	"{{.ModuleName}}/internal/pkg/logger"

	// Import VPKG modules here:
	// "{{.ModuleName}}/internal/vpkg/transport/http"
	// "{{.ModuleName}}/internal/vpkg/storage/postgres"
)

func main() {
	app := fx.New(
		// Core infrastructure (always present)
		config.Module,
		logger.Module,

		// VPKG modules (user adds these):
		// http.Module,
		// postgres.Module,

		// Lifecycle hooks
		fx.Invoke(registerHooks),
	)

	// Start the application
	app.Run()
}

func registerHooks(lc fx.Lifecycle, log logger.Logger, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting application...",
				"env", cfg.App.Env,
				"version", cfg.App.Version,
			)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shutting down application...")
			return nil
		},
	})

	// Graceful shutdown
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				sigChan := make(chan os.Signal, 1)
				signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
				<-sigChan
				log.Info("Received shutdown signal")
			}()
			return nil
		},
	})
}
