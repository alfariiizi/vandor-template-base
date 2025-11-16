# Vandor Base Template

Base template for creating new Vandor projects with hexagonal architecture.

## Overview

This is the official base template for [Vandor](https://github.com/alfariiizi/vandor) - a Go scaffolding tool for building backend applications with clean architecture.

**Version:** 1.0.0  
**Compatible with:** Vandor CLI v0.1.0+

## What's Included

### Architecture Structure

**Hexagonal Architecture (Ports & Adapters)**

```
template/
├── cmd/
│   ├── app/              # Main application entry point
│   ├── enum/             # Code generation for enums
│   ├── scheduler/        # Job scheduler commands
│   └── seed/             # Database seeding commands
│
├── internal/
│   ├── domain/           # Business logic (domain-driven)
│   │   └── domain-template/  # Template for new domains
│   ├── adapter/          # Infrastructure & I/O
│   ├── pkg/              # Shared utilities
│   │   ├── config/       # Configuration management
│   │   ├── logger/       # Logging utilities
│   │   └── testutil/     # Testing helpers
│   └── vpkg/             # VPKG packages directory
│
├── config/
│   └── config.yaml.example  # Configuration template
│
├── test/
│   ├── e2e/              # End-to-end tests
│   ├── integration/      # Integration tests
│   └── fixture/          # Test fixtures
│
├── scripts/
│   └── ent-tools.sh      # Ent code generation helpers
│
├── taskfile.yaml         # Task runner configuration
├── go.mod.tmpl           # Go module template
└── README.md.tmpl        # Project README template
```

### Features

- **Clean Architecture**: Hexagonal architecture with clear domain boundaries
- **Code Generation**: Built-in generators for domains, use cases, services
- **Testing Support**: Structure for unit, integration, and e2e tests
- **Task Runner**: Pre-configured Taskfile for common operations
- **Configuration**: Per-command config with smart merging and secret management
- **Logging**: Structured logging setup
- **VPKG System**: Package management for reusable components
- **Multi-Command Support**: Create multiple entry points (app, worker, migrate) with separate configs

## Usage

### With Vandor CLI (Recommended)

```bash
# Install Vandor CLI
curl -fsSL https://raw.githubusercontent.com/alfariiizi/vandor/main/install.sh | bash

# Create new project
vandor init myproject
```

The CLI will automatically:
- Fetch this template from GitHub
- Replace template variables
- Initialize Go module
- Set up git repository

### Manual Usage

```bash
# Clone this repository
git clone https://github.com/alfariiizi/vandor-template-base.git myproject
cd myproject

# Move template files to root
mv template/* .
rmdir template

# Process template variables manually
# Replace {{.ModulePath}} with your module path (e.g., github.com/username/myproject)
# Replace {{.ProjectName}} with your project name
# Replace {{.Version}} with your version (e.g., 0.1.0)

# Rename template files
mv go.mod.tmpl go.mod
mv README.md.tmpl README.md

# Initialize
go mod tidy
```

## Template Variables

These variables are automatically replaced during project generation:

| Variable | Description | Example |
|----------|-------------|---------|
| `{{.ModulePath}}` | Go module path | `github.com/username/myproject` |
| `{{.ProjectName}}` | Project name | `myproject` |
| `{{.Version}}` | Initial version | `0.1.0` |

## Configuration System

### Per-Command Configuration

Each command has its own configuration package at `cmd/{name}/config/`:

```
cmd/app/config/
├── config.go      # Config struct (defines what this command needs)
├── module.go      # FX module with smart loader
└── providers.go   # Provider functions
```

### Smart Config Merging

Config files in `config/` directory are automatically discovered and merged:

```
config/
├── config.yaml            # Base (shared by all commands)
├── config.app.yaml        # App-specific overrides
├── config.worker.yaml     # Worker-specific overrides
└── config.production.yaml # Environment-specific
```

**Merge order** (later overrides earlier):
1. `config.yaml` (base)
2. `config.*.yaml` (partials)
3. `config.{command}.yaml` (command-specific)
4. `config.{env}.yaml` (environment: development/production)
5. `.env` file (development only)
6. Environment variables (highest priority)

### Secret Management

**Development:**
```bash
# Copy template
cp .env.example .env

# Edit with your secrets
vim .env
```

**Production:**
```bash
# Use environment variables (e.g., from K8s Secrets)
export POSTGRES_PASSWORD=supersecret
export OPENAI_API_KEY=sk-...
```

**Security:**
- ✅ `.env.example` is committed (template)
- ❌ `.env` is gitignored (real secrets)
- ❌ Never commit secrets to `config/*.yaml`

## After Generation

Once your project is created, you can:

### 1. Build and Run

```bash
# Build the project
task build

# Run in development mode
task run:dev

# Run tests
task test
```

### 2. Add Domains

```bash
# Create a new domain
vandor add domain order

# Add use case
vandor add usecase order/PlaceOrder

# Add handler
vandor add handler order/Create
```

### 3. Install VPKG Packages

```bash
# Add HTTP server
vandor vpkg add transport/http

# Add PostgreSQL
vandor vpkg add storage/postgres

# Add Redis cache
vandor vpkg add cache/redis
```

## Structure Philosophy

### Domain-Driven Design

- **Domain**: Business capabilities (not database tables!)
  - ❌ Wrong: `user`, `product`, `order_item`
  - ✅ Right: `identity`, `catalog`, `order`
- Each domain is self-contained with its own:
  - Entities and value objects (`core/entity`, `core/value`)
  - Use cases (`application/`)
  - Ports/interfaces (`port/`)

### Hexagonal Architecture

- **Core** (`internal/domain/`): Pure business logic, no external dependencies
- **Ports** (`port/`): Interfaces defining how core interacts with outside
- **Adapters** (`internal/adapter/`): Implementations of ports (HTTP, database, etc.)
- **Infrastructure** (`internal/pkg/`): Shared utilities (logger, config)

## Development Workflow

```bash
# 1. Create domain
vandor add domain identity

# 2. Implement business logic
# Edit: internal/domain/identity/core/entity/user.go
#       internal/domain/identity/application/register_user.go

# 3. Add HTTP handler
vandor add handler identity/Register --method=POST

# 4. Generate code
vandor sync all

# 5. Test
task test

# 6. Run
task run:dev
```

## Customization

This template is designed to be modified for your needs:

- Add/remove directories in `internal/`
- Customize `Taskfile.yaml` commands
- Modify `config/config.yaml.example`
- Extend `internal/pkg/` utilities

## Requirements

- **Go**: 1.23 or higher
- **Git**: For version control
- **Task**: (Optional) Task runner, or use `go run`

## Template Versions

| Version | Vandor CLI | Changes |
|---------|------------|---------|
| 1.0.0 | >=0.1.0 | Initial release with hexagonal architecture |

## Contributing

This template is part of the Vandor project. To contribute:

1. Fork this repository
2. Create feature branch
3. Make changes
4. Submit pull request

See [Vandor contributing guide](https://github.com/alfariiizi/vandor/blob/main/CONTRIBUTING.md)

## Related Links

- **Vandor CLI**: https://github.com/alfariiizi/vandor
- **Documentation**: https://vandor.dev/docs
- **Examples**: https://vandor.dev/docs/examples
- **Discord**: https://discord.gg/vandor

## License

MIT License - See [LICENSE](LICENSE) file

---

**Created with ❤️ by the Vandor Team**
