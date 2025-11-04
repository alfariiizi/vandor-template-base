# Test Directory

This directory contains project-wide tests organized by test type.

## Structure

```
test/
├── integration/    # Cross-domain integration tests
├── e2e/            # End-to-end API tests
└── fixture/        # Shared test fixtures and helpers
```

## Test Strategy

### 1. Unit Tests (Co-located)
Located next to the code they test:
```
domain/order/application/place_order.go
domain/order/application/place_order_test.go
```

**Focus**: Test individual components in isolation with mocks.

### 2. Integration Tests (`test/integration/`)
Test workflows across multiple components:
```
test/integration/
└── order_flow_test.go    # Tests order placement through payment
```

**Focus**: Test real dependencies (database, cache) using testcontainers.

### 3. E2E Tests (`test/e2e/`)
Test complete system through HTTP API:
```
test/e2e/
└── api_test.go    # Full API workflow tests
```

**Focus**: Test user-facing functionality end-to-end.

### 4. Fixtures (`test/fixture/`)
Shared test data and builders:
```
test/fixture/
├── user.go       # User test builders
├── order.go      # Order test builders
└── database.go   # Database test helpers
```

## Running Tests

```bash
# Run all tests
task test

# Run with coverage
task test:coverage

# Run only unit tests
go test ./internal/domain/... -v

# Run integration tests
go test ./test/integration/... -v

# Run e2e tests
go test ./test/e2e/... -v
```

## Test Conventions

### Naming
- Unit tests: `*_test.go`
- BDD specs (optional): `*_spec_test.go`
- Integration tests: `*_integration_test.go`

### Build Tags
Use build tags to separate test types:

```go
//go:build integration
package integration_test
```

Run with: `go test -tags=integration ./...`

## Testing with Testcontainers

Integration tests use testcontainers for real dependencies:

```go
import "github.com/testcontainers/testcontainers-go"

func TestOrderRepository(t *testing.T) {
    ctx := context.Background()

    // Start PostgreSQL container
    pgContainer, err := postgres.RunContainer(ctx)
    require.NoError(t, err)
    defer pgContainer.Terminate(ctx)

    // Run tests against real database
    // ...
}
```

## BDD Testing (Optional)

For behavior-driven development, use Ginkgo/Gomega or Godog:

```go
// domain/order/application/place_order_spec_test.go
var _ = Describe("PlaceOrder", func() {
    Context("when order is valid", func() {
        It("should create order successfully", func() {
            // ...
        })
    })
})
```

## Learn More

- [Testing Guide](https://docs.vandor.dev/testing)
- [BDD with Ginkgo](https://docs.vandor.dev/testing/bdd)
- [Integration Testing](https://docs.vandor.dev/testing/integration)
