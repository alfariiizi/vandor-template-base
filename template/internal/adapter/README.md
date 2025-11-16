# Adapter Layer

This directory contains infrastructure adapters that connect your domain to
external systems.

## Structure

```
adapter/
├── vandor/http/           # HTTP handlers, middleware (installed via VPKG)
├── vandor/storage/        # Repository implementations (installed via VPKG)
├── vandor/messaging/      # Message queue adapters (installed via VPKG)
└── vandor/cache/          # Cache implementations (installed via VPKG)
```

## Adding Adapters

Adapters are added via VPKG packages:

```bash
# Add HTTP server
vandor vpkg add vandor/http

# Add MinIO object storage
vandor vpkg add vandor/minio

# See all available packages
vandor vpkg list
```

## Adapter Responsibilities

### HTTP Adapter (`adapter/http/`)

- HTTP handlers
- Request/response mapping
- Middleware
- Route registration

### Storage Adapter (`adapter/storage/`)

- Repository implementations (ports from domain layer)
- Database queries
- Transaction management
- Data mapping

### Messaging Adapter (`adapter/messaging/`)

- Event publishers
- Event subscribers
- Message transformation

### Cache Adapter (`adapter/cache/`)

- Cache implementations
- Cache key management
- TTL handling

## Dependency Direction

```
Domain Layer (ports) ← Adapter Layer (implementations)
```

Adapters depend on domain, never the reverse.

## Learn More

- [Adapter Pattern Guide](https://docs.vandor.dev/architecture/adapters)
- [VPKG Packages](https://docs.vandor.dev/vpkg)
