# 🎯 Level 3, Topic 1: Context Package Mastery

## 🚀 Overview
The `context` package is one of the most powerful and essential packages in Go for managing request-scoped data, cancellation signals, and timeouts across API boundaries and between processes. This topic will take you from basic context usage to advanced patterns that will make you a context master.

---

## 📚 Table of Contents

1. [Context Fundamentals](#context-fundamentals)
2. [Context Types and Creation](#context-types-and-creation)
3. [Context Propagation](#context-propagation)
4. [Cancellation Patterns](#cancellation-patterns)
5. [Timeout Patterns](#timeout-patterns)
6. [Value Context](#value-context)
7. [Context Hierarchies](#context-hierarchies)
8. [Advanced Patterns](#advanced-patterns)
9. [Performance Considerations](#performance-considerations)
10. [Best Practices](#best-practices)
11. [Common Pitfalls](#common-pitfalls)
12. [Real-World Applications](#real-world-applications)

---

## 🎯 Context Fundamentals

### What is Context?

Context is a Go package that provides:
- **Request-scoped values**: Data that can be passed through the call chain
- **Cancellation signals**: A way to signal cancellation across goroutines
- **Deadlines**: A way to set timeouts for operations
- **Request tracing**: A way to trace requests through the system

### Why Context Matters

```go
// Without context - no cancellation, no timeouts, no request tracing
func processRequest(data []byte) error {
    // What if this takes too long?
    // What if the client cancels?
    // How do we trace this request?
    return doWork(data)
}

// With context - full control over lifecycle
func processRequest(ctx context.Context, data []byte) error {
    // Can be cancelled
    // Can have timeouts
    // Can carry request-scoped data
    return doWork(ctx, data)
}
```

### Key Principles

1. **Context should be the first parameter** of functions
2. **Never store context in structs** - pass it explicitly
3. **Context is immutable** - create new contexts, don't modify existing ones
4. **Always check for cancellation** in long-running operations
5. **Use context for cancellation, not for passing optional parameters**

---

## 🏗️ Context Types and Creation

### 1. Background Context

```go
// The root context - never cancelled, no values, no deadline
ctx := context.Background()

// Use for:
// - Main function
// - Incoming requests to servers
// - Outgoing calls to servers
// - Top-level testing
```

### 2. TODO Context

```go
// A placeholder context - should be replaced
ctx := context.TODO()

// Use when:
// - You're not sure which context to use
// - You're refactoring code
// - The context is not yet available
```

### 3. WithCancel Context

```go
// Creates a context that can be cancelled
ctx, cancel := context.WithCancel(context.Background())

// Cancel the context
cancel()

// Use for:
// - Manual cancellation
// - Graceful shutdown
// - User-initiated cancellation
```

### 4. WithTimeout Context

```go
// Creates a context with a timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // Important: always call cancel

// Use for:
// - HTTP requests
// - Database operations
// - Any operation with a time limit
```

### 5. WithDeadline Context

```go
// Creates a context with a specific deadline
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Use for:
// - Operations that must complete by a specific time
// - Batch processing with deadlines
```

### 6. WithValue Context

```go
// Creates a context with key-value pairs
type key string
const userKey key = "user"

ctx := context.WithValue(context.Background(), userKey, "john@example.com")

// Use for:
// - Request-scoped data
// - User information
// - Request IDs
// - Authentication tokens
```

---

## 🔄 Context Propagation

### Basic Propagation

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Create context from request
    ctx := r.Context()
    
    // Pass context through call chain
    result, err := processRequest(ctx, r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Write(result)
}

func processRequest(ctx context.Context, body io.Reader) ([]byte, error) {
    // Check if context is cancelled
    if ctx.Err() != nil {
        return nil, ctx.Err()
    }
    
    // Pass context to next function
    return doWork(ctx, body)
}

func doWork(ctx context.Context, body io.Reader) ([]byte, error) {
    // Long-running operation with cancellation check
    for i := 0; i < 1000; i++ {
        select {
        case <-ctx.Done():
            return nil, ctx.Err()
        default:
            // Do work
            time.Sleep(1 * time.Millisecond)
        }
    }
    return []byte("done"), nil
}
```

### Context Inheritance

```go
func main() {
    // Root context
    rootCtx := context.Background()
    
    // Create child context with timeout
    childCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
    defer cancel()
    
    // Create grandchild context with value
    grandchildCtx := context.WithValue(childCtx, "requestID", "12345")
    
    // All contexts are linked
    // Cancelling childCtx will cancel grandchildCtx
    // But not rootCtx
}
```

---

## ⏹️ Cancellation Patterns

### 1. Manual Cancellation

```go
func manualCancellation() {
    ctx, cancel := context.WithCancel(context.Background())
    
    go func() {
        // Simulate work
        for i := 0; i < 100; i++ {
            select {
            case <-ctx.Done():
                fmt.Println("Work cancelled")
                return
            default:
                fmt.Printf("Working... %d\n", i)
                time.Sleep(100 * time.Millisecond)
            }
        }
    }()
    
    // Cancel after 2 seconds
    time.Sleep(2 * time.Second)
    cancel()
    
    // Wait for cancellation to propagate
    time.Sleep(100 * time.Millisecond)
}
```

### 2. Timeout Cancellation

```go
func timeoutCancellation() {
    // Create context with 3-second timeout
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    go func() {
        for i := 0; i < 100; i++ {
            select {
            case <-ctx.Done():
                fmt.Printf("Timeout reached: %v\n", ctx.Err())
                return
            default:
                fmt.Printf("Working... %d\n", i)
                time.Sleep(500 * time.Millisecond)
            }
        }
    }()
    
    // Wait for timeout
    time.Sleep(5 * time.Second)
}
```

### 3. Deadline Cancellation

```go
func deadlineCancellation() {
    // Set deadline to 2 seconds from now
    deadline := time.Now().Add(2 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()
    
    go func() {
        for i := 0; i < 100; i++ {
            select {
            case <-ctx.Done():
                fmt.Printf("Deadline reached: %v\n", ctx.Err())
                return
            default:
                fmt.Printf("Working... %d\n", i)
                time.Sleep(300 * time.Millisecond)
            }
        }
    }()
    
    // Wait for deadline
    time.Sleep(3 * time.Second)
}
```

---

## ⏱️ Timeout Patterns

### 1. HTTP Request Timeout

```go
func httpRequestWithTimeout() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Create HTTP request with context
    req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/data", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // Make request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            log.Println("Request timed out")
        } else {
            log.Printf("Request failed: %v", err)
        }
        return
    }
    defer resp.Body.Close()
    
    // Process response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Response: %s\n", body)
}
```

### 2. Database Operation Timeout

```go
func databaseOperationWithTimeout() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Simulate database operation
    result := make(chan string, 1)
    go func() {
        // Simulate slow database query
        time.Sleep(8 * time.Second)
        result <- "query result"
    }()
    
    select {
    case res := <-result:
        fmt.Printf("Database result: %s\n", res)
    case <-ctx.Done():
        fmt.Printf("Database operation timed out: %v\n", ctx.Err())
    }
}
```

### 3. Cascading Timeouts

```go
func cascadingTimeouts() {
    // Parent context with 10-second timeout
    parentCtx, parentCancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer parentCancel()
    
    // Child context with 5-second timeout
    childCtx, childCancel := context.WithTimeout(parentCtx, 5*time.Second)
    defer childCancel()
    
    // Grandchild context with 3-second timeout
    grandchildCtx, grandchildCancel := context.WithTimeout(childCtx, 3*time.Second)
    defer grandchildCancel()
    
    // The grandchild will timeout first (3 seconds)
    // Then child (5 seconds)
    // Then parent (10 seconds)
    
    go func() {
        for i := 0; i < 100; i++ {
            select {
            case <-grandchildCtx.Done():
                fmt.Printf("Grandchild timeout: %v\n", grandchildCtx.Err())
                return
            default:
                fmt.Printf("Working... %d\n", i)
                time.Sleep(500 * time.Millisecond)
            }
        }
    }()
    
    time.Sleep(6 * time.Second)
}
```

---

## 🔑 Value Context

### 1. Basic Value Context

```go
type key string

const (
    userKey    key = "user"
    requestKey key = "requestID"
    traceKey   key = "traceID"
)

func valueContextExample() {
    // Create context with values
    ctx := context.WithValue(context.Background(), userKey, "john@example.com")
    ctx = context.WithValue(ctx, requestKey, "req-12345")
    ctx = context.WithValue(ctx, traceKey, "trace-67890")
    
    // Use values in functions
    processWithValues(ctx)
}

func processWithValues(ctx context.Context) {
    // Extract values
    user := ctx.Value(userKey).(string)
    requestID := ctx.Value(requestKey).(string)
    traceID := ctx.Value(traceKey).(string)
    
    fmt.Printf("User: %s, Request: %s, Trace: %s\n", user, requestID, traceID)
}
```

### 2. Type-Safe Value Context

```go
type contextKey string

const (
    userContextKey    contextKey = "user"
    requestContextKey contextKey = "requestID"
)

type User struct {
    ID    string
    Email string
    Role  string
}

func setUser(ctx context.Context, user User) context.Context {
    return context.WithValue(ctx, userContextKey, user)
}

func getUser(ctx context.Context) (User, bool) {
    user, ok := ctx.Value(userContextKey).(User)
    return user, ok
}

func typeSafeValueExample() {
    ctx := context.Background()
    
    user := User{
        ID:    "123",
        Email: "john@example.com",
        Role:  "admin",
    }
    
    ctx = setUser(ctx, user)
    
    if retrievedUser, ok := getUser(ctx); ok {
        fmt.Printf("User: %+v\n", retrievedUser)
    }
}
```

### 3. Request Scoped Data

```go
func requestScopedDataExample() {
    // Simulate HTTP request
    ctx := context.Background()
    
    // Add request-scoped data
    ctx = addRequestData(ctx, "req-12345", "user-67890", "session-abc123")
    
    // Process request
    processRequest(ctx)
}

func addRequestData(ctx context.Context, requestID, userID, sessionID string) context.Context {
    ctx = context.WithValue(ctx, "requestID", requestID)
    ctx = context.WithValue(ctx, "userID", userID)
    ctx = context.WithValue(ctx, "sessionID", sessionID)
    return ctx
}

func processRequest(ctx context.Context) {
    // Extract request data
    requestID := ctx.Value("requestID").(string)
    userID := ctx.Value("userID").(string)
    sessionID := ctx.Value("sessionID").(string)
    
    fmt.Printf("Processing request %s for user %s (session: %s)\n", 
        requestID, userID, sessionID)
}
```

---

## 🌳 Context Hierarchies

### 1. Simple Hierarchy

```go
func simpleHierarchy() {
    // Root context
    rootCtx := context.Background()
    
    // Level 1: Add timeout
    level1Ctx, cancel1 := context.WithTimeout(rootCtx, 10*time.Second)
    defer cancel1()
    
    // Level 2: Add value
    level2Ctx := context.WithValue(level1Ctx, "level", 2)
    
    // Level 3: Add another value
    level3Ctx := context.WithValue(level2Ctx, "operation", "process")
    
    // All contexts inherit from root
    // Cancelling level1Ctx cancels level2Ctx and level3Ctx
    // But not rootCtx
}
```

### 2. Complex Hierarchy

```go
func complexHierarchy() {
    // Root context
    rootCtx := context.Background()
    
    // Create multiple branches
    branch1Ctx, cancel1 := context.WithTimeout(rootCtx, 5*time.Second)
    defer cancel1()
    
    branch2Ctx, cancel2 := context.WithTimeout(rootCtx, 8*time.Second)
    defer cancel2()
    
    // Add values to branches
    branch1Ctx = context.WithValue(branch1Ctx, "branch", "1")
    branch2Ctx = context.WithValue(branch2Ctx, "branch", "2")
    
    // Create sub-branches
    subBranch1Ctx := context.WithValue(branch1Ctx, "sub", "1.1")
    subBranch2Ctx := context.WithValue(branch2Ctx, "sub", "2.1")
    
    // Process each branch
    go processBranch(subBranch1Ctx)
    go processBranch(subBranch2Ctx)
    
    time.Sleep(6 * time.Second)
}

func processBranch(ctx context.Context) {
    for i := 0; i < 100; i++ {
        select {
        case <-ctx.Done():
            fmt.Printf("Branch %s cancelled: %v\n", 
                ctx.Value("branch"), ctx.Err())
            return
        default:
            fmt.Printf("Branch %s working... %d\n", 
                ctx.Value("branch"), i)
            time.Sleep(200 * time.Millisecond)
        }
    }
}
```

---

## 🔬 Advanced Patterns

### 1. Context Middleware

```go
func contextMiddleware() {
    // Create middleware chain
    ctx := context.Background()
    ctx = addRequestID(ctx)
    ctx = addUser(ctx, "john@example.com")
    ctx = addTimeout(ctx, 5*time.Second)
    
    // Process request
    processRequest(ctx)
}

func addRequestID(ctx context.Context) context.Context {
    requestID := generateRequestID()
    return context.WithValue(ctx, "requestID", requestID)
}

func addUser(ctx context.Context, email string) context.Context {
    return context.WithValue(ctx, "user", email)
}

func addTimeout(ctx context.Context, timeout time.Duration) context.Context {
    newCtx, _ := context.WithTimeout(ctx, timeout)
    return newCtx
}

func generateRequestID() string {
    return fmt.Sprintf("req-%d", time.Now().UnixNano())
}
```

### 2. Context Pool

```go
type ContextPool struct {
    contexts chan context.Context
    factory  func() context.Context
}

func NewContextPool(size int, factory func() context.Context) *ContextPool {
    pool := &ContextPool{
        contexts: make(chan context.Context, size),
        factory:  factory,
    }
    
    // Pre-populate pool
    for i := 0; i < size; i++ {
        pool.contexts <- factory()
    }
    
    return pool
}

func (p *ContextPool) Get() context.Context {
    select {
    case ctx := <-p.contexts:
        return ctx
    default:
        return p.factory()
    }
}

func (p *ContextPool) Put(ctx context.Context) {
    select {
    case p.contexts <- ctx:
    default:
        // Pool is full, discard context
    }
}
```

### 3. Context Chain

```go
func contextChain() {
    // Create chain of contexts
    ctx := context.Background()
    
    // Chain operations
    ctx = withRequestID(ctx)
    ctx = withUser(ctx, "john@example.com")
    ctx = withTimeout(ctx, 5*time.Second)
    ctx = withTrace(ctx, "trace-12345")
    
    // Process chain
    processChain(ctx)
}

func withRequestID(ctx context.Context) context.Context {
    return context.WithValue(ctx, "requestID", generateRequestID())
}

func withUser(ctx context.Context, user string) context.Context {
    return context.WithValue(ctx, "user", user)
}

func withTimeout(ctx context.Context, timeout time.Duration) context.Context {
    newCtx, _ := context.WithTimeout(ctx, timeout)
    return newCtx
}

func withTrace(ctx context.Context, traceID string) context.Context {
    return context.WithValue(ctx, "traceID", traceID)
}

func processChain(ctx context.Context) {
    // Extract all values
    requestID := ctx.Value("requestID")
    user := ctx.Value("user")
    traceID := ctx.Value("traceID")
    
    fmt.Printf("Processing chain: %s, %s, %s\n", requestID, user, traceID)
}
```

---

## ⚡ Performance Considerations

### 1. Context Creation Overhead

```go
func contextCreationOverhead() {
    // ❌ Bad: Creating context in hot path
    for i := 0; i < 1000000; i++ {
        ctx := context.WithValue(context.Background(), "key", i)
        processValue(ctx)
    }
    
    // ✅ Good: Create context once, reuse
    ctx := context.WithValue(context.Background(), "key", "value")
    for i := 0; i < 1000000; i++ {
        processValue(ctx)
    }
}

func processValue(ctx context.Context) {
    // Process value
    _ = ctx.Value("key")
}
```

### 2. Context Value Lookup

```go
func contextValueLookup() {
    ctx := context.Background()
    
    // Add many values
    for i := 0; i < 100; i++ {
        ctx = context.WithValue(ctx, fmt.Sprintf("key%d", i), i)
    }
    
    // Lookup performance
    start := time.Now()
    for i := 0; i < 1000; i++ {
        _ = ctx.Value("key50")
    }
    duration := time.Since(start)
    
    fmt.Printf("Lookup time: %v\n", duration)
}
```

### 3. Context Cancellation Performance

```go
func contextCancellationPerformance() {
    ctx, cancel := context.WithCancel(context.Background())
    
    // Create many goroutines
    for i := 0; i < 1000; i++ {
        go func(id int) {
            for {
                select {
                case <-ctx.Done():
                    return
                default:
                    // Do work
                    time.Sleep(1 * time.Millisecond)
                }
            }
        }(i)
    }
    
    // Measure cancellation time
    start := time.Now()
    cancel()
    duration := time.Since(start)
    
    fmt.Printf("Cancellation time: %v\n", duration)
}
```

---

## 🎯 Best Practices

### 1. Function Signatures

```go
// ✅ Good: Context as first parameter
func processRequest(ctx context.Context, data []byte) error {
    // Implementation
}

// ❌ Bad: Context not first parameter
func processRequest(data []byte, ctx context.Context) error {
    // Implementation
}

// ❌ Bad: Context in struct
type Processor struct {
    ctx context.Context
    // Other fields
}
```

### 2. Context Checking

```go
// ✅ Good: Check context in long-running operations
func longRunningOperation(ctx context.Context) error {
    for i := 0; i < 1000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do work
            time.Sleep(1 * time.Millisecond)
        }
    }
    return nil
}

// ❌ Bad: No context checking
func longRunningOperation(ctx context.Context) error {
    for i := 0; i < 1000; i++ {
        // Do work without checking context
        time.Sleep(1 * time.Millisecond)
    }
    return nil
}
```

### 3. Context Cleanup

```go
// ✅ Good: Always call cancel
func withTimeout() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // Important!
    
    // Use ctx
}

// ❌ Bad: Forgetting to call cancel
func withTimeout() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    // Forgot to call cancel!
    
    // Use ctx
}
```

### 4. Value Context Keys

```go
// ✅ Good: Custom key type
type contextKey string

const userKey contextKey = "user"

func setUser(ctx context.Context, user string) context.Context {
    return context.WithValue(ctx, userKey, user)
}

// ❌ Bad: String keys
func setUser(ctx context.Context, user string) context.Context {
    return context.WithValue(ctx, "user", user)
}
```

---

## ⚠️ Common Pitfalls

### 1. Storing Context in Structs

```go
// ❌ Bad: Storing context in struct
type Service struct {
    ctx context.Context
}

func (s *Service) Process() error {
    // Use s.ctx - this is wrong!
    return doWork(s.ctx)
}

// ✅ Good: Pass context as parameter
type Service struct {
    // No context field
}

func (s *Service) Process(ctx context.Context) error {
    return doWork(ctx)
}
```

### 2. Not Checking Context

```go
// ❌ Bad: No context checking
func processData(ctx context.Context, data []byte) error {
    // Long operation without checking context
    time.Sleep(10 * time.Second)
    return nil
}

// ✅ Good: Check context regularly
func processData(ctx context.Context, data []byte) error {
    for i := 0; i < 1000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do work
            time.Sleep(10 * time.Millisecond)
        }
    }
    return nil
}
```

### 3. Context Leaks

```go
// ❌ Bad: Context leak
func leakyFunction() {
    ctx, cancel := context.WithCancel(context.Background())
    // Forgot to call cancel!
    
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                // Work
            }
        }
    }()
}

// ✅ Good: Proper cleanup
func properFunction() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // Always call cancel
    
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                // Work
            }
        }
    }()
}
```

### 4. Wrong Context Type

```go
// ❌ Bad: Using wrong context type
func wrongContext() {
    ctx := context.TODO() // Should use Background() for new requests
    
    // Process request
    processRequest(ctx)
}

// ✅ Good: Use appropriate context type
func correctContext() {
    ctx := context.Background() // Correct for new requests
    
    // Process request
    processRequest(ctx)
}
```

---

## 🌍 Real-World Applications

### 1. HTTP Server with Context

```go
func httpServerWithContext() {
    http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
        // Get context from request
        ctx := r.Context()
        
        // Add request-scoped data
        ctx = context.WithValue(ctx, "requestID", generateRequestID())
        ctx = context.WithValue(ctx, "userID", getUserID(r))
        
        // Process request with context
        result, err := processRequest(ctx, r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        w.Write(result)
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 2. Database Operations with Context

```go
func databaseOperationsWithContext() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Database operations with context
    user, err := getUserByID(ctx, "123")
    if err != nil {
        log.Printf("Failed to get user: %v", err)
        return
    }
    
    // Update user with context
    err = updateUser(ctx, user)
    if err != nil {
        log.Printf("Failed to update user: %v", err)
        return
    }
}

func getUserByID(ctx context.Context, id string) (*User, error) {
    // Simulate database query
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    case <-time.After(2 * time.Second):
        return &User{ID: id, Name: "John"}, nil
    }
}
```

### 3. Microservices with Context

```go
func microservicesWithContext() {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // Call multiple services
    userChan := make(chan *User, 1)
    orderChan := make(chan *Order, 1)
    
    go func() {
        user, err := userService.GetUser(ctx, "123")
        if err != nil {
            log.Printf("User service error: %v", err)
            return
        }
        userChan <- user
    }()
    
    go func() {
        order, err := orderService.GetOrder(ctx, "456")
        if err != nil {
            log.Printf("Order service error: %v", err)
            return
        }
        orderChan <- order
    }()
    
    // Wait for results
    select {
    case user := <-userChan:
        fmt.Printf("User: %+v\n", user)
    case order := <-orderChan:
        fmt.Printf("Order: %+v\n", order)
    case <-ctx.Done():
        fmt.Printf("Timeout: %v\n", ctx.Err())
    }
}
```

---

## 🎓 Summary

The Context package is essential for building robust, scalable Go applications. Key takeaways:

1. **Always use context** for cancellation and timeouts
2. **Pass context as first parameter** of functions
3. **Check context regularly** in long-running operations
4. **Use appropriate context types** for different scenarios
5. **Never store context in structs** - pass it explicitly
6. **Always call cancel** to prevent leaks
7. **Use custom key types** for context values
8. **Consider performance implications** of context operations

Mastering the Context package will make you a more effective Go developer and help you build better concurrent applications! 🚀

---

## 🚀 Next Steps

1. **Practice** with the provided examples
2. **Implement** context in your own projects
3. **Experiment** with different context patterns
4. **Move to the next topic** in the curriculum
5. **Apply** context patterns to real-world scenarios

Ready to become a Context master? Let's dive into the implementation! 💪

