# Configuration Directory

This directory contains the base configuration files for your application. These YAML files are loaded and merged by all commands created with `vandor add cmd`.

## How It Works

When you run any command (e.g., `cmd/app`, `cmd/worker`, `cmd/cli-tool`), the configuration system automatically:

1. Loads `config/config.yaml` as the base configuration
2. Merges any partial configs like `config/config.database.yaml`
3. Merges command-specific config like `config/config.app.yaml`
4. Merges environment-specific config like `config/config.production.yaml`
5. Loads secrets from `.env` file (development) or environment variables (production)

## File Structure

```
config/
├── config.yaml              # Base configuration (required)
├── config.app.yaml          # App-specific overrides (optional)
├── config.worker.yaml       # Worker-specific overrides (optional)
├── config.development.yaml  # Development environment (optional)
├── config.production.yaml   # Production environment (optional)
└── README.md                # This file
```

## Loading Priority

Configuration values are merged in this order (later values override earlier ones):

1. `config/config.yaml` - Base configuration (lowest priority)
2. `config/config.*.yaml` - All partial configs (alphabetically)
3. `config/config.{command}.yaml` - Command-specific (e.g., `config.app.yaml`)
4. `config/config.{env}.yaml` - Environment-specific (e.g., `config.production.yaml`)
5. `.env` file - Secrets for development
6. Environment variables - Highest priority (production secrets)

## Base Configuration

The `config.yaml` file contains all available configuration options with:
- Default values for development
- Validation rules in comments
- Examples for each field
- VPKG package configuration templates (commented out)

## Adding VPKG Package Configs

When you install a VPKG package, uncomment its configuration section in `config.yaml`:

```bash
# Install HTTP server package
vandor vpkg add vandor/http

# Then uncomment the http section in config/config.yaml
```

## Environment-Specific Configs

Create environment-specific files to override values:

**config/config.production.yaml:**
```yaml
app:
  env: production

logger:
  level: warn
  format: json

# http:
#   host: 0.0.0.0
#   port: 8080
```

## Command-Specific Configs

Create command-specific files for different entry points:

**config/config.worker.yaml:**
```yaml
logger:
  level: debug  # More verbose for background jobs
```

## Secrets Management

**Development:** Use `.env` file in project root
```bash
# .env
DATABASE_URL=postgres://user:pass@localhost:5432/mydb
REDIS_URL=redis://localhost:6379
```

**Production:** Use environment variables (Kubernetes Secrets, Docker secrets, etc.)
```bash
export DATABASE_URL=postgres://user:pass@db.example.com:5432/mydb
```

## Relationship with cmd/*/config

Each command created with `vandor add cmd` has its own `config/` package that:
- Defines the `Config` struct with all fields
- Loads and merges files from this `config/` directory
- Provides the configuration to the application

**For FX-module commands:**
- Uses `fx.Module` for dependency injection
- Has `providers.go` for extracting sub-configs

**For Cobra CLI commands:**
- Uses simple `LoadConfig()` function
- No FX dependency injection overhead

## Validation

All configuration values are validated using [go-playground/validator](https://github.com/go-playground/validator). Invalid configuration will cause the application to fail at startup with clear error messages.

Common validation tags:
- `required` - Field must be present
- `min=1`, `max=65535` - Numeric range
- `oneof=debug info warn error` - Must be one of these values
- `hostname_port` - Must be valid host:port format
- `url` - Must be valid URL
- `email` - Must be valid email
- `ip` - Must be valid IP address

For the full list of available validators, see the [validator documentation](https://pkg.go.dev/github.com/go-playground/validator/v10#readme-baked-in-validations).

## Next Steps

1. Review and customize `config.yaml` for your needs
2. Create environment-specific configs as needed
3. Set up `.env` for development secrets
4. Configure production secrets via environment variables
