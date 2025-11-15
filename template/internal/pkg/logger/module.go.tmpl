package logger

import (
	"go.uber.org/fx"

	"{{.Module}}/internal/pkg/config"
)

// Module provides logger for dependency injection
var Module = fx.Module("logger",
	fx.Provide(NewLogger),
)

// NewLogger creates a new logger from config
// Uses config.LoggerConfig instead of full *config.Config to avoid circular dependencies
func NewLogger(cfg config.LoggerConfig) Logger {
	loggerCfg := &Config{
		Level:  cfg.Level,
		Format: cfg.Format,
	}
	return New(loggerCfg)
}
