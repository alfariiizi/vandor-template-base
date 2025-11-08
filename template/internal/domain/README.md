# Domain Layer

This directory contains domain logic organized by business capabilities (NOT database tables).

## Structure

Each domain follows DDD aggregate-centric structure:

```
domain/{name}/
├── core/
│   ├── {aggregate}.go      # Aggregate root + nested entities
│   ├── value_objects.go    # Value objects and domain errors
│   └── service/            # Domain services (optional)
├── port/
│   └── repository.go       # Repository interface (one per aggregate)
├── application/            # Use cases
└── module.go               # FX dependency injection module
```

## Aggregates vs Entities

**Aggregate**: Cluster of domain objects treated as a single unit
- Has ONE aggregate root (entry point)
- Contains nested entities
- Enforces business invariants
- Defines transaction boundary

**Example**: Order aggregate
```
Order (aggregate root)
├── OrderItem (nested entity)
├── Payment (nested entity)
├── ShippingAddress (value object)
└── OrderStatus (value object)
```

## Creating Domains

```bash
# Create domain
vandor add domain order

# Add aggregate to domain
vandor add aggregate order Order --entities=OrderItem,Payment
```

## Domain Examples

**E-Commerce:**
- `identity/` - User aggregate (users, authentication, profiles)
- `catalog/` - Product aggregate (products, categories, inventory)
- `order/` - Order aggregate (orders, order items, payments, shipments)
- `review/` - Review aggregate (customer reviews, ratings)

**SaaS Platform:**
- `identity/` - User aggregate (users, organizations, memberships)
- `workspace/` - Project aggregate (projects, tasks, attachments)
- `billing/` - Subscription aggregate (subscriptions, invoices, payments)
- `integration/` - Integration aggregate (API keys, webhooks, audit logs)

## Best Practices

1. ✅ Think in business capabilities (identity, catalog, order)
2. ✅ One repository per aggregate root
3. ✅ Keep aggregates small and focused
4. ✅ Use `vandor add aggregate` to create new aggregates
5. ❌ Don't create one domain per database table
6. ❌ Don't create repositories for nested entities
7. ❌ Don't have circular dependencies between domains

## Learn More

- [Domain-Driven Design Guide](https://docs.vandor.dev/ddd)
- [Hexagonal Architecture](https://docs.vandor.dev/architecture/hexagonal)
- [Aggregate Pattern](https://docs.vandor.dev/ddd/aggregates)
