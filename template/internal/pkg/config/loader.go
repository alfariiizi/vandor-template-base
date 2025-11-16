package config

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Load loads configuration from file and environment variables
func Load(configPath string) (*Config, error) {
	v := viper.New()

	// Set defaults from struct tags
	if err := setDefaultsFromTags(v, &Config{}); err != nil {
		return nil, fmt.Errorf("failed to set defaults: %w", err)
	}

	// Set config file
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath("./config")
		v.AddConfigPath(".")
	}

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		// Config file not found is okay if we have environment variables
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Environment variables
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Unmarshal config
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate config
	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}

// setDefaultsFromTags sets default values from struct tags
func setDefaultsFromTags(v *viper.Viper, cfg interface{}) error {
	return setDefaultsRecursive(v, "", reflect.ValueOf(cfg).Elem(), reflect.TypeOf(cfg).Elem())
}

// setDefaultsRecursive recursively sets defaults from struct tags
func setDefaultsRecursive(v *viper.Viper, prefix string, val reflect.Value, typ reflect.Type) error {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		// Get mapstructure tag for the key name
		mapstructureTag := field.Tag.Get("mapstructure")
		if mapstructureTag == "" || mapstructureTag == "-" {
			continue
		}

		// Build the full key path
		var key string
		if prefix != "" {
			key = prefix + "." + mapstructureTag
		} else {
			key = mapstructureTag
		}

		// Handle nested structs
		if field.Type.Kind() == reflect.Struct {
			if err := setDefaultsRecursive(v, key, fieldValue, field.Type); err != nil {
				return err
			}
			continue
		}

		// Get default tag
		defaultTag := field.Tag.Get("default")
		if defaultTag == "" {
			continue
		}

		// Set default value based on field type
		switch field.Type.Kind() {
		case reflect.String:
			v.SetDefault(key, defaultTag)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(defaultTag, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int default for %s: %w", key, err)
			}
			v.SetDefault(key, intVal)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintVal, err := strconv.ParseUint(defaultTag, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid uint default for %s: %w", key, err)
			}
			v.SetDefault(key, uintVal)
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(defaultTag)
			if err != nil {
				return fmt.Errorf("invalid bool default for %s: %w", key, err)
			}
			v.SetDefault(key, boolVal)
		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(defaultTag, 64)
			if err != nil {
				return fmt.Errorf("invalid float default for %s: %w", key, err)
			}
			v.SetDefault(key, floatVal)
		default:
			return fmt.Errorf("unsupported type for default tag: %s (%s)", key, field.Type.Kind())
		}
	}

	return nil
}
