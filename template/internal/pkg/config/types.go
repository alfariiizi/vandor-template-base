package config

// AppConfig holds application-level configuration
// This is a common type that can be shared across all commands
type AppConfig struct {
	Name    string `mapstructure:"name" validate:"required,min=1" default:"vandor-app" example:"my-awesome-app" doc:"Application name used for logging and metrics"`
	Version string `mapstructure:"version" validate:"required,semver" default:"0.1.0" example:"1.2.3" doc:"Application version in semantic versioning format"`
	Env     string `mapstructure:"env" validate:"required,oneof=development staging production" default:"development" example:"production" doc:"Environment mode: development, staging, or production"`
}
