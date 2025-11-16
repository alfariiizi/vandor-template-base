package logger

import (
	"go.uber.org/fx"
)

// Module provides logger for dependency injection
var Module = fx.Module("logger",
	fx.Provide(NewLogger),
)

// NewLogger creates a new logger from config
func NewLogger(cfg Config) Logger {
	loggerCfg := &Config{
		Level:  cfg.Level,
		Format: cfg.Format,
	}
	return New(loggerCfg)
}
