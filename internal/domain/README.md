# Domain Layer

This directory contains your business logic organized by **business capabilities**, not database tables.

## Creating Domains

Domains represent business capabilities:

```bash
# Create a domain (business capability)
vandor add domain order
```

For hexagonal architecture, this creates:

```
domain/order/
├── core/
│   ├── entity/          # Domain entities
│   ├── value/           # Value objects
│   └── aggregate/       # Aggregate roots
├── port/                # Interfaces (ports)
│   ├── repository.go    # Repository port
│   └── service.go       # Service port
├── application/         # Use cases
│   ├── place_order.go
│   └── cancel_order.go
└── test/
    ├── integration/     # Integration tests
    ├── fixture/         # Test builders
    └── mock/            # Mocked interfaces
```

## Domain Examples

**E-Commerce:**
- `identity/` - Users, authentication, profiles
- `catalog/` - Products, categories, inventory
- `order/` - Orders, payments, shipments
- `review/` - Customer reviews, ratings

**SaaS Platform:**
- `identity/` - Users, organizations
- `workspace/` - Projects, tasks
- `billing/` - Subscriptions, invoices
- `integration/` - API keys, webhooks

## Anti-Patterns to Avoid

❌ One table = one domain (leads to anemic model)
❌ Too many tiny domains (over-fragmentation)
❌ Circular dependencies between domains

## Learn More

- [Domain-Driven Design Guide](https://docs.vandor.dev/ddd)
- [Hexagonal Architecture](https://docs.vandor.dev/architecture/hexagonal)
