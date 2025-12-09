---
id: 2025-12-09-map-string-any-support
title: Change Error Data from map[string]string to map[string]any
status: Implemented
created: 2025-12-09
authors: ["@bborbe"]
tags: [prd, error-handling, breaking-change]
related_issues: []
implemented_pr: ""
superseded_by: ""
---

# PRD: Change Error Data from map[string]string to map[string]any

**Status**: Implemented
**Created**: 2025-12-09
**Obsidian Task**: [[Improve Error Details with map string any Support]]
**Target**: github.com/bborbe/errors library

## Summary

Change the `HasData` interface and related functions from `map[string]string` to `map[string]any` to support richer error details including arrays and nested objects. This enables JSON error handlers to return structured data like `{"strategies": ["id1", "id2"]}` instead of comma-separated strings.

## Problem Statement

### Current Behavior

The `HasData` interface and `AddDataToError` function only support string values:

```go
type HasData interface {
    Data() map[string]string
}

func AddDataToError(err error, data map[string]string) DataError
```

### Limitation

When HTTP handlers need to include arrays in error details, they must join them into comma-separated strings:

```go
// Current workaround
errors.AddDataToError(err, map[string]string{
    "strategies": strings.Join(strategyIdentifiers.Strings(), ","),
})
```

**JSON output (current):**
```json
{
  "error": {
    "details": {
      "strategies": "ftmo-live-3-ORB_GER40.cash_15MIN_V26,another-strategy"
    }
  }
}
```

**JSON output (desired):**
```json
{
  "error": {
    "details": {
      "strategies": ["ftmo-live-3-ORB_GER40.cash_15MIN_V26", "another-strategy"]
    }
  }
}
```

### Real-World Impact (2025-12-09)

During `/start-day` command, the trading daily summary failed because `frontend-report` couldn't find strategy `ftmo-live-3-ORB_GER40.cash_15MIN_V26`. The error details should include the list of strategies that were requested, but the current implementation forces joining them into a string.

## Goals

### Primary Goals

1. **Richer Error Data** - Support arrays, numbers, booleans, and nested objects in error details
2. **Backward Compatible Usage** - Existing code passing `map[string]string` should still work
3. **JSON Serializable** - All values must serialize cleanly to JSON
4. **Simple Migration** - Minimal changes required for consumers

### Non-Goals

- ❌ Support for non-JSON-serializable types (channels, functions)
- ❌ Deep validation of value types
- ❌ Automatic type conversion

## Technical Specification

### Interface Changes

**Before:**
```go
type HasData interface {
    Data() map[string]string
}
```

**After:**
```go
type HasData interface {
    Data() map[string]any
}
```

### Function Changes

**errors_data-error.go:**

```go
// Before
func AddDataToError(err error, data map[string]string) DataError

// After
func AddDataToError(err error, data map[string]any) DataError

// Before
func DataFromError(err error) map[string]string

// After
func DataFromError(err error) map[string]any
```

**errors_data-context.go:**

```go
// Before
func AddToContext(ctx context.Context, key, value string) context.Context

// After - support any value type
func AddToContext(ctx context.Context, key string, value any) context.Context

// Before
func DataFromContext(ctx context.Context) map[string]string

// After
func DataFromContext(ctx context.Context) map[string]any
```

### Struct Changes

```go
// Before
type dataError struct {
    err  error
    data map[string]string
}

// After
type dataError struct {
    err  error
    data map[string]any
}
```

### Usage Examples

**Arrays:**
```go
errors.AddDataToError(err, map[string]any{
    "strategies": []string{"strategy-1", "strategy-2"},
    "invalidFields": []string{"email", "phone"},
})
```

**Mixed types:**
```go
errors.AddDataToError(err, map[string]any{
    "field": "columnGroup",
    "received": "",
    "expected": []string{"day", "week", "month", "year"},
    "count": 4,
})
```

**Nested objects:**
```go
errors.AddDataToError(err, map[string]any{
    "validation": map[string]any{
        "field": "email",
        "reason": "invalid_format",
    },
})
```

## Implementation Plan

### Phase 1: Update Core Types (30 minutes)

**File: errors_data-error.go**

1. Change `HasData` interface return type
2. Change `AddDataToError` parameter type
3. Change `dataError.data` field type
4. Change `Data()` method return type
5. Change `DataFromError` return type and internal map creation

### Phase 2: Update Context Functions (15 minutes)

**File: errors_data-context.go**

1. Change `AddToContext` value parameter to `any`
2. Change context value type from `map[string]string` to `map[string]any`
3. Change `DataFromContext` return type

### Phase 3: Update Tests (30 minutes)

**File: errors_data-error_test.go**

1. Update existing tests to use `map[string]any`
2. Add tests for array values
3. Add tests for mixed type values

**File: errors_data-context_test.go**

1. Update existing tests
2. Add tests for non-string context values

### Phase 4: Documentation (15 minutes)

1. Update README if needed
2. Add CHANGELOG entry

**Total Estimate**: ~1.5 hours

## Breaking Changes

### Interface Change

Any code implementing `HasData` interface must update return type:

```go
// Before
func (e *myError) Data() map[string]string

// After
func (e *myError) Data() map[string]any
```

### Type Assertions

Code reading from `Data()` expecting strings needs type assertions:

```go
// Before
value := data["key"]

// After - if you need string
value, ok := data["key"].(string)
```

### Mitigation

- Most consumers use `AddDataToError` and `DataFromError` functions
- JSON serialization handles `any` types automatically
- Passing `map[string]string` to functions expecting `map[string]any` requires explicit conversion

## Downstream Impact

### github.com/bborbe/http

Must update after this change:

1. `ErrorDetails.Details` field: `map[string]string` → `map[string]any`
2. `WrapWithDetails` function signature
3. Tests

### Other Consumers

Any library or service using `HasData` interface directly needs updates.

## Testing Strategy

### Unit Tests

```go
func TestAddDataToError_WithArray(t *testing.T) {
    err := errors.AddDataToError(
        errors.New(context.Background(), "test"),
        map[string]any{
            "ids": []string{"a", "b", "c"},
        },
    )
    data := errors.DataFromError(err)
    ids, ok := data["ids"].([]string)
    Expect(ok).To(BeTrue())
    Expect(ids).To(Equal([]string{"a", "b", "c"}))
}

func TestAddDataToError_WithMixedTypes(t *testing.T) {
    err := errors.AddDataToError(
        errors.New(context.Background(), "test"),
        map[string]any{
            "field": "email",
            "count": 5,
            "valid": false,
        },
    )
    data := errors.DataFromError(err)
    Expect(data["field"]).To(Equal("email"))
    Expect(data["count"]).To(Equal(5))
    Expect(data["valid"]).To(Equal(false))
}
```

### Integration Test (with http library)

Verify JSON serialization produces expected output with arrays.

## Rollout Plan

1. Implement changes in github.com/bborbe/errors
2. Run all tests
3. Tag new version (v1.x.0 - minor version bump for breaking interface change)
4. Update github.com/bborbe/http to use new version
5. Update consuming services (frontend-report, etc.)

## Open Questions

- [x] Should we keep backward compatibility helper for `map[string]string`? → No, explicit conversion is clear enough
- [ ] Should we validate that values are JSON-serializable? → Probably not, keep it simple

## Updates Log

**2025-12-09**: Initial PRD created based on frontend-report error handling needs
