package logger

// LoggerConfig holds logger configuration
type Config struct {
	Level  string `mapstructure:"level" validate:"required,oneof=debug info warn error" default:"info" example:"debug" doc:"Log level: debug, info, warn, or error"`
	Format string `mapstructure:"format" validate:"required,oneof=json console" default:"console" example:"json" doc:"Log format: json for production, console for development"`
}
