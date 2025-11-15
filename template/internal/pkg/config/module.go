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

// NewConfig creates a new config from file
func NewConfig() (*Config, error) {
	// Parse config path from command line
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	return Load(*configPath)
}
