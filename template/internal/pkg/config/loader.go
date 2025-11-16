package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// LoaderOptions configures the config loader
type LoaderOptions struct {
	ConfigDir   string   // Directory containing config files (default: "./config")
	CommandName string   // Command name for command-specific configs
	EnvFile     string   // Path to .env file (default: ".env")
	LoadEnvFile bool     // Whether to load .env file (default: true in development)
	EnvPrefix   string   // Prefix for environment variables (default: "APP")
	Patterns    []string // Additional file patterns to load
}

// DefaultLoaderOptions returns default loader options
func DefaultLoaderOptions() *LoaderOptions {
	return &LoaderOptions{
		ConfigDir:   "./config",
		EnvFile:     ".env",
		LoadEnvFile: os.Getenv("APP_ENV") != "production",
		EnvPrefix:   "APP",
		Patterns:    []string{},
	}
}

// LoadForCommand loads and merges all relevant config files for a command
// Merge order (later overrides earlier):
//  1. config.yaml (base)
//  2. config.*.yaml (all other partials)
//  3. config.{command}.yaml (command-specific)
//  4. config.{env}.yaml (environment-specific)
//  5. .env file (if enabled)
//  6. Environment variables
func LoadForCommand(opts *LoaderOptions) (map[string]interface{}, error) {
	if opts == nil {
		opts = DefaultLoaderOptions()
	}

	// 1. Load .env file first (if enabled)
	if opts.LoadEnvFile {
		if err := loadEnvFile(opts.EnvFile); err != nil {
			// Non-fatal: .env file is optional
			// fmt.Printf("Note: .env file not loaded: %v\n", err)
		}
	}

	// 2. Discover all config files
	files, err := discoverConfigFiles(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to discover config files: %w", err)
	}

	// 3. Merge all YAML files
	merged, err := mergeConfigFiles(files, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to merge config files: %w", err)
	}

	return merged, nil
}

// loadEnvFile loads environment variables from .env file
func loadEnvFile(envFile string) error {
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		return fmt.Errorf(".env file not found: %s", envFile)
	}

	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	return nil
}

// discoverConfigFiles finds all config files to load based on patterns
func discoverConfigFiles(opts *LoaderOptions) ([]string, error) {
	var files []string
	env := getEnvironment()

	// Priority order for file discovery
	patterns := []string{
		"config.yaml",              // Base config (highest priority base)
		"config.yml",               // Alternative extension
	}

	// Add all partial configs (config.*.yaml) except command and env specific
	partials, err := filepath.Glob(filepath.Join(opts.ConfigDir, "config.*.yaml"))
	if err == nil {
		for _, partial := range partials {
			base := filepath.Base(partial)
			// Skip command-specific and env-specific files
			if !strings.Contains(base, opts.CommandName) && !strings.Contains(base, env) {
				patterns = append(patterns, base)
			}
		}
	}

	// Add command-specific config
	if opts.CommandName != "" {
		patterns = append(patterns,
			fmt.Sprintf("config.%s.yaml", opts.CommandName),
			fmt.Sprintf("config.%s.yml", opts.CommandName),
		)
	}

	// Add environment-specific config (highest priority)
	patterns = append(patterns,
		fmt.Sprintf("config.%s.yaml", env),
		fmt.Sprintf("config.%s.yml", env),
	)

	// Add custom patterns
	patterns = append(patterns, opts.Patterns...)

	// Convert patterns to full paths and check existence
	for _, pattern := range patterns {
		path := filepath.Join(opts.ConfigDir, pattern)
		if _, err := os.Stat(path); err == nil {
			files = append(files, path)
		}
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no config files found in %s", opts.ConfigDir)
	}

	return files, nil
}

// mergeConfigFiles merges multiple config files using Viper
func mergeConfigFiles(files []string, opts *LoaderOptions) (map[string]interface{}, error) {
	v := viper.New()

	// Configure Viper
	v.SetEnvPrefix(opts.EnvPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Merge all config files in order
	for _, file := range files {
		v.SetConfigFile(file)
		if err := v.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("failed to merge config file %s: %w", file, err)
		}
	}

	// Convert to map
	return v.AllSettings(), nil
}

// getEnvironment returns current environment
func getEnvironment() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	return env
}

// DetectCommandName attempts to detect the command name from binary or path
func DetectCommandName() string {
	// Try from binary name first
	binary := filepath.Base(os.Args[0])
	if binary != "" && binary != "main" && binary != "go" {
		return binary
	}

	// Try from working directory (if running from cmd/{name}/)
	wd, err := os.Getwd()
	if err == nil && strings.Contains(wd, "/cmd/") {
		parts := strings.Split(wd, "/cmd/")
		if len(parts) > 1 {
			cmdPart := strings.Split(parts[1], "/")[0]
			if cmdPart != "" {
				return cmdPart
			}
		}
	}

	// Fallback: try from environment variable
	if cmdName := os.Getenv("APP_COMMAND"); cmdName != "" {
		return cmdName
	}

	// Default fallback
	return "app"
}
