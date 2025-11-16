package config

import (
	"flag"

	"go.uber.org/fx"
)

// Module provides config for dependency injection
var Module = fx.Module("config",
	fx.Provide(
		NewConfig,
		ProvideAppConfig,
		ProvideLoggerConfig,
	),
)

var configPath string

func init() {
	// Register flag only once
	flag.StringVar(&configPath, "config", "", "Path to config file")
}

// NewConfig creates a new config from file
func NewConfig() (*Config, error) {
	// Parse flags only if not already parsed
	if !flag.Parsed() {
		flag.Parse()
	}

	return Load(configPath)
}
