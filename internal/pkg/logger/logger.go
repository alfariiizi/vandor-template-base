package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the interface for structured logging
type Logger interface {
	// Debug logs a debug message
	Debug(msg string, keysAndValues ...interface{})

	// Info logs an info message
	Info(msg string, keysAndValues ...interface{})

	// Warn logs a warning message
	Warn(msg string, keysAndValues ...interface{})

	// Error logs an error message
	Error(msg string, keysAndValues ...interface{})

	// Fatal logs a fatal message and exits
	Fatal(msg string, keysAndValues ...interface{})

	// With returns a logger with additional context
	With(keysAndValues ...interface{}) Logger

	// WithContext returns a logger with context
	WithContext(ctx context.Context) Logger
}

// zapLogger is the zap implementation of Logger
type zapLogger struct {
	logger *zap.SugaredLogger
}

// Config holds logger configuration
type Config struct {
	Level  string // debug, info, warn, error
	Format string // json, console
}

// DefaultConfig returns default logger configuration
func DefaultConfig() *Config {
	return &Config{
		Level:  "info",
		Format: "console",
	}
}

// New creates a new logger instance
func New(cfg *Config) Logger {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	// Parse log level
	level := parseLevel(cfg.Level)

	// Create zap config
	var zapConfig zap.Config
	if cfg.Format == "json" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig.Level = zap.NewAtomicLevelAt(level)

	// Build logger
	logger, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &zapLogger{logger: logger.Sugar()}
}

func parseLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func (l *zapLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *zapLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) With(keysAndValues ...interface{}) Logger {
	return &zapLogger{logger: l.logger.With(keysAndValues...)}
}

func (l *zapLogger) WithContext(ctx context.Context) Logger {
	// For now, just return the logger
	// In the future, we can extract values from context if needed
	return l
}
