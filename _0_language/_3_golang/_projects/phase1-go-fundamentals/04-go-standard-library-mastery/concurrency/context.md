# context Package - Context Management 🔄

The `context` package provides a way to carry request-scoped values, cancellation signals, and deadlines across API boundaries and between processes. It's essential for managing goroutine lifecycles and request cancellation.

## 🎯 Key Concepts

### 1. **Context Types**
- `Context` - Interface for context
- `Background()` - Empty context
- `TODO()` - Placeholder context
- `WithCancel()` - Context with cancellation
- `WithTimeout()` - Context with timeout
- `WithDeadline()` - Context with deadline
- `WithValue()` - Context with values

### 2. **Context Operations**
- `Done()` - Channel for cancellation
- `Err()` - Error if context is cancelled
- `Value()` - Get value from context
- `Deadline()` - Get deadline if set
- `Cancel()` - Cancel context
- `CancelFunc` - Function to cancel context

### 3. **Context Propagation**
- Pass context as first parameter
- Use context in function signatures
- Propagate context through call chains
- Don't store context in structs
- Use context for cancellation

### 4. **Context Values**
- `WithValue()` - Add values to context
- `Value()` - Retrieve values from context
- Use custom types for keys
- Avoid using built-in types as keys
- Keep values immutable

### 5. **Context Cancellation**
- `WithCancel()` - Manual cancellation
- `WithTimeout()` - Time-based cancellation
- `WithDeadline()` - Deadline-based cancellation
- `Done()` - Check for cancellation
- `Err()` - Get cancellation reason

### 6. **Context Best Practices**
- Always pass context as first parameter
- Use context for cancellation
- Don't store context in structs
- Use context values sparingly
- Handle context cancellation

## 🚀 Common Patterns

### Basic Context Usage
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        return
    case <-time.After(5 * time.Second):
        // Do work
    }
}()
```

### Context with Timeout
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Use ctx in operations
result, err := doSomething(ctx)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        // Handle timeout
    }
}
```

### Context with Values
```go
type key string
const userKey key = "user"

ctx := context.WithValue(context.Background(), userKey, "john")
user := ctx.Value(userKey).(string)
```

### Context Propagation
```go
func processRequest(ctx context.Context, data string) error {
    // Pass context to sub-functions
    return processData(ctx, data)
}

func processData(ctx context.Context, data string) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
        // Process data
        return nil
    }
}
```

## ⚠️ Common Pitfalls

1. **Not passing context** - Always pass context as first parameter
2. **Storing context in structs** - Don't store context in structs
3. **Ignoring context cancellation** - Always check ctx.Done()
4. **Using wrong key types** - Use custom types for context keys
5. **Not handling context errors** - Always handle context errors

## 🎯 Best Practices

1. **Pass context first** - Always pass context as first parameter
2. **Use context for cancellation** - Use context for goroutine cancellation
3. **Don't store context** - Don't store context in structs
4. **Handle context errors** - Always handle context errors
5. **Use custom key types** - Use custom types for context keys

## 🔍 Advanced Features

### Custom Context
```go
type customContext struct {
    context.Context
    customValue string
}

func (c *customContext) Value(key interface{}) interface{} {
    if key == "custom" {
        return c.customValue
    }
    return c.Context.Value(key)
}
```

### Context Middleware
```go
func withTimeout(ctx context.Context, timeout time.Duration) context.Context {
    ctx, _ = context.WithTimeout(ctx, timeout)
    return ctx
}
```

### Context Chain
```go
func processWithChain(ctx context.Context) error {
    // Create context chain
    ctx = context.WithValue(ctx, "step", "1")
    ctx = context.WithTimeout(ctx, 5*time.Second)
    ctx = context.WithValue(ctx, "step", "2")
    
    return processStep(ctx)
}
```

## 📚 Real-world Applications

1. **HTTP Servers** - Request context propagation
2. **Database Operations** - Query cancellation
3. **API Clients** - Request timeout
4. **Background Jobs** - Job cancellation
5. **Microservices** - Request tracing

## 🧠 Memory Tips

- **context** = **C**ontext **O**perations **N**etwork **T**oolkit **E**ngine **X**ecution **T**ool
- **Background** = **B**ackground context
- **WithCancel** = **W**ith **C**ancel
- **WithTimeout** = **W**ith **T**imeout
- **WithValue** = **W**ith **V**alue
- **Done** = **D**one channel
- **Err** = **E**rror
- **Value** = **V**alue

Remember: The context package is your gateway to request lifecycle management in Go! 🎯
