package logger

import (
	"go.uber.org/fx"

	"{{.ModuleName}}/internal/pkg/config"
)

// Module provides logger for dependency injection
var Module = fx.Module("logger",
	fx.Provide(NewLogger),
)

// NewLogger creates a new logger from config
func NewLogger(cfg *config.Config) Logger {
	loggerCfg := &Config{
		Level:  cfg.Logger.Level,
		Format: cfg.Logger.Format,
	}
	return New(loggerCfg)
}
