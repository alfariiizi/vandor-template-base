# VPKG - Vandor Package System

This directory contains installed VPKG packages - reusable infrastructure components.

## What is VPKG?

VPKG is Vandor's extension mechanism for adding infrastructure components to your project.

## Package Types

### 1. FX Module Packages (`fx-module`)
Library components integrated into your application:
- `transport/http` - HTTP server
- `storage/postgres` - PostgreSQL database
- `cache/redis` - Redis cache
- `messaging/kafka` - Kafka messaging

### 2. CLI Command Packages (`cli-command`)
Executable tools:
- Code generators
- Migration tools
- Development utilities

## Installing Packages

```bash
# Install a package
vandor vpkg add transport/http

# List available packages
vandor vpkg list

# Search packages
vandor vpkg search cache

# Remove a package
vandor vpkg remove transport/http

# Run CLI package
vandor vpkg exec migration-tool
```

## Package Structure

When you install a package like `transport/http`:

```
vpkg/transport/http/
├── module.go         # FX module
├── server.go         # Implementation
├── middleware.go     # Middleware
├── vpkg.yaml         # Package metadata
└── README.md         # Usage documentation
```

## Using Installed Packages

After installing, add the module to `cmd/app/main.go`:

```go
import "{{.ModuleName}}/internal/vpkg/transport/http"

fx.New(
    config.Module,
    logger.Module,
    http.Module,  // Add installed VPKG module
)
```

## Available Package Categories

- `transport/*` - HTTP, gRPC, GraphQL, WebSocket
- `storage/*` - PostgreSQL, MongoDB, Redis
- `cache/*` - Redis, Memcached, in-memory
- `messaging/*` - Kafka, RabbitMQ, NATS, Asynq
- `observability/*` - LGMT, Prometheus, traces
- `auth/*` - JWT, OAuth, sessions
- `external/*` - Third-party API clients

## Learn More

- [VPKG Documentation](https://docs.vandor.dev/vpkg)
- [Creating Custom Packages](https://docs.vandor.dev/vpkg/custom)
- [Package Registry](https://registry.vandor.dev)
