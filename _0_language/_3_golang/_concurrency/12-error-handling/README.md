# ⚠️ Level 3, Topic 4: Error Handling in Concurrent Code

## 🚀 Overview
Mastering error handling in concurrent code is crucial for building robust, reliable, and maintainable systems. This topic will take you from basic error handling patterns to advanced strategies that will make you an error handling expert in concurrent Go programs.

---

## 📚 Table of Contents

1. [Error Handling Fundamentals](#error-handling-fundamentals)
2. [Error Propagation Patterns](#error-propagation-patterns)
3. [Error Aggregation Strategies](#error-aggregation-strategies)
4. [Panic Recovery Patterns](#panic-recovery-patterns)
5. [Error Context and Wrapping](#error-context-and-wrapping)
6. [Graceful Degradation](#graceful-degradation)
7. [Circuit Breaker Pattern](#circuit-breaker-pattern)
8. [Retry Mechanisms](#retry-mechanisms)
9. [Error Monitoring and Logging](#error-monitoring-and-logging)
10. [Error Testing Strategies](#error-testing-strategies)
11. [Advanced Error Patterns](#advanced-error-patterns)
12. [Real-World Applications](#real-world-applications)

---

## ⚠️ Error Handling Fundamentals

### What is Error Handling in Concurrent Code?

Error handling in concurrent code involves:
- **Error propagation** across goroutines and channels
- **Error aggregation** from multiple concurrent operations
- **Panic recovery** in goroutines
- **Error context** preservation and wrapping
- **Graceful degradation** when errors occur
- **Circuit breaker** patterns for fault tolerance

### Key Principles

#### 1. Fail Fast vs. Fail Safe
```go
// Fail Fast: Stop immediately on error
func failFast() error {
    if err := validateInput(); err != nil {
        return err // Stop immediately
    }
    // Continue processing
}

// Fail Safe: Continue with degraded functionality
func failSafe() error {
    if err := optionalFeature(); err != nil {
        log.Printf("Optional feature failed: %v", err)
        // Continue with core functionality
    }
    return nil
}
```

#### 2. Error Context Preservation
```go
// Preserve error context
func processData(data []byte) error {
    if err := validateData(data); err != nil {
        return fmt.Errorf("data validation failed: %w", err)
    }
    
    if err := transformData(data); err != nil {
        return fmt.Errorf("data transformation failed: %w", err)
    }
    
    return nil
}
```

#### 3. Error Aggregation
```go
// Aggregate multiple errors
func processMultiple(items []Item) error {
    var errors []error
    
    for _, item := range items {
        if err := processItem(item); err != nil {
            errors = append(errors, err)
        }
    }
    
    if len(errors) > 0 {
        return fmt.Errorf("processing failed: %v", errors)
    }
    
    return nil
}
```

---

## 🔄 Error Propagation Patterns

### 1. Channel-Based Error Propagation

```go
// Error channel pattern
func processWithErrorChannel() {
    dataCh := make(chan int)
    errorCh := make(chan error)
    
    go func() {
        defer close(dataCh)
        defer close(errorCh)
        
        for i := 0; i < 10; i++ {
            if i == 5 {
                errorCh <- fmt.Errorf("processing failed at %d", i)
                return
            }
            dataCh <- i
        }
    }()
    
    for {
        select {
        case data, ok := <-dataCh:
            if !ok {
                return
            }
            fmt.Printf("Processed: %d\n", data)
        case err := <-errorCh:
            fmt.Printf("Error: %v\n", err)
            return
        }
    }
}
```

### 2. Result Pattern

```go
// Result pattern with error handling
type Result struct {
    Data interface{}
    Err  error
}

func processWithResult() {
    resultCh := make(chan Result)
    
    go func() {
        defer close(resultCh)
        
        for i := 0; i < 10; i++ {
            if i == 5 {
                resultCh <- Result{Err: fmt.Errorf("processing failed at %d", i)}
                return
            }
            resultCh <- Result{Data: i}
        }
    }()
    
    for result := range resultCh {
        if result.Err != nil {
            fmt.Printf("Error: %v\n", result.Err)
            return
        }
        fmt.Printf("Processed: %v\n", result.Data)
    }
}
```

### 3. Error Wrapper Pattern

```go
// Error wrapper for context preservation
type ErrorWrapper struct {
    Operation string
    Context   map[string]interface{}
    Err       error
}

func (ew ErrorWrapper) Error() string {
    return fmt.Sprintf("%s failed: %v", ew.Operation, ew.Err)
}

func (ew ErrorWrapper) Unwrap() error {
    return ew.Err
}

func processWithWrapper() error {
    if err := validateInput(); err != nil {
        return ErrorWrapper{
            Operation: "validation",
            Context:   map[string]interface{}{"input": "data"},
            Err:       err,
        }
    }
    
    return nil
}
```

---

## 🔗 Error Aggregation Strategies

### 1. Error Collection

```go
// Collect multiple errors
type ErrorCollector struct {
    errors []error
    mu     sync.Mutex
}

func (ec *ErrorCollector) Add(err error) {
    ec.mu.Lock()
    defer ec.mu.Unlock()
    ec.errors = append(ec.errors, err)
}

func (ec *ErrorCollector) Error() error {
    ec.mu.Lock()
    defer ec.mu.Unlock()
    
    if len(ec.errors) == 0 {
        return nil
    }
    
    return fmt.Errorf("multiple errors occurred: %v", ec.errors)
}

func processWithCollection() error {
    collector := &ErrorCollector{}
    
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            if err := processItem(id); err != nil {
                collector.Add(fmt.Errorf("item %d failed: %w", id, err))
            }
        }(i)
    }
    
    wg.Wait()
    return collector.Error()
}
```

### 2. Error Group Pattern

```go
// Error group for concurrent operations
type ErrorGroup struct {
    errors []error
    mu     sync.Mutex
}

func (eg *ErrorGroup) Go(fn func() error) {
    go func() {
        if err := fn(); err != nil {
            eg.mu.Lock()
            eg.errors = append(eg.errors, err)
            eg.mu.Unlock()
        }
    }()
}

func (eg *ErrorGroup) Wait() error {
    eg.mu.Lock()
    defer eg.mu.Unlock()
    
    if len(eg.errors) == 0 {
        return nil
    }
    
    return fmt.Errorf("group errors: %v", eg.errors)
}

func processWithGroup() error {
    group := &ErrorGroup{}
    
    for i := 0; i < 5; i++ {
        id := i
        group.Go(func() error {
            return processItem(id)
        })
    }
    
    time.Sleep(1 * time.Second) // Wait for completion
    return group.Wait()
}
```

### 3. Error Channel Aggregation

```go
// Aggregate errors from multiple channels
func aggregateErrors() error {
    errorChs := make([]<-chan error, 3)
    
    // Create error channels
    for i := 0; i < 3; i++ {
        ch := make(chan error, 1)
        errorChs[i] = ch
        
        go func(id int, ch chan<- error) {
            defer close(ch)
            if err := processItem(id); err != nil {
                ch <- err
            }
        }(i, ch)
    }
    
    // Aggregate errors
    var errors []error
    for _, ch := range errorChs {
        if err := <-ch; err != nil {
            errors = append(errors, err)
        }
    }
    
    if len(errors) > 0 {
        return fmt.Errorf("aggregated errors: %v", errors)
    }
    
    return nil
}
```

---

## 🛡️ Panic Recovery Patterns

### 1. Basic Panic Recovery

```go
// Basic panic recovery in goroutine
func basicPanicRecovery() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    // This will panic
    panic("something went wrong")
}
```

### 2. Goroutine Panic Recovery

```go
// Panic recovery in goroutine
func goroutinePanicRecovery() {
    ch := make(chan int)
    
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("Goroutine recovered from panic: %v\n", r)
            }
        }()
        
        // This will panic
        panic("goroutine panic")
    }()
    
    time.Sleep(100 * time.Millisecond)
    close(ch)
}
```

### 3. Panic Recovery with Error Channel

```go
// Panic recovery with error reporting
func panicRecoveryWithError() {
    errorCh := make(chan error, 1)
    
    go func() {
        defer func() {
            if r := recover(); r != nil {
                errorCh <- fmt.Errorf("panic recovered: %v", r)
            }
        }()
        
        // This will panic
        panic("something went wrong")
    }()
    
    select {
    case err := <-errorCh:
        fmt.Printf("Error from panic: %v\n", err)
    case <-time.After(1 * time.Second):
        fmt.Println("No panic occurred")
    }
}
```

### 4. Panic Recovery Middleware

```go
// Panic recovery middleware
func panicRecoveryMiddleware(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic recovered: %v", r)
        }
    }()
    
    fn()
    return nil
}

func usePanicMiddleware() {
    if err := panicRecoveryMiddleware(func() {
        panic("middleware panic")
    }); err != nil {
        fmt.Printf("Middleware caught panic: %v\n", err)
    }
}
```

---

## 🔍 Error Context and Wrapping

### 1. Error Context Preservation

```go
// Preserve error context
func preserveErrorContext() error {
    if err := step1(); err != nil {
        return fmt.Errorf("step1 failed: %w", err)
    }
    
    if err := step2(); err != nil {
        return fmt.Errorf("step2 failed: %w", err)
    }
    
    return nil
}

func step1() error {
    return fmt.Errorf("validation failed")
}

func step2() error {
    return fmt.Errorf("processing failed")
}
```

### 2. Error Wrapping with Context

```go
// Error wrapping with additional context
type ContextualError struct {
    Context map[string]interface{}
    Err     error
}

func (ce ContextualError) Error() string {
    return fmt.Sprintf("context: %v, error: %v", ce.Context, ce.Err)
}

func (ce ContextualError) Unwrap() error {
    return ce.Err
}

func processWithContext() error {
    context := map[string]interface{}{
        "operation": "data processing",
        "timestamp": time.Now(),
        "user_id":   "123",
    }
    
    if err := processData(); err != nil {
        return ContextualError{
            Context: context,
            Err:     err,
        }
    }
    
    return nil
}
```

### 3. Error Chain Tracing

```go
// Error chain tracing
func traceErrorChain() error {
    if err := level1(); err != nil {
        return fmt.Errorf("level1 failed: %w", err)
    }
    return nil
}

func level1() error {
    if err := level2(); err != nil {
        return fmt.Errorf("level2 failed: %w", err)
    }
    return nil
}

func level2() error {
    return fmt.Errorf("root cause")
}
```

---

## 🎯 Graceful Degradation

### 1. Fallback Pattern

```go
// Fallback pattern for graceful degradation
func fallbackPattern() error {
    // Try primary service
    if err := primaryService(); err != nil {
        log.Printf("Primary service failed: %v", err)
        
        // Try fallback service
        if err := fallbackService(); err != nil {
            log.Printf("Fallback service failed: %v", err)
            return fmt.Errorf("both services failed")
        }
    }
    
    return nil
}

func primaryService() error {
    return fmt.Errorf("primary service unavailable")
}

func fallbackService() error {
    return nil // Fallback succeeds
}
```

### 2. Circuit Breaker Pattern

```go
// Circuit breaker for fault tolerance
type CircuitBreaker struct {
    state       int // 0: closed, 1: open, 2: half-open
    failureCount int
    threshold   int
    timeout     time.Duration
    lastFailure time.Time
    mu          sync.Mutex
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    if cb.state == 1 { // Open
        if time.Since(cb.lastFailure) > cb.timeout {
            cb.state = 2 // Half-open
        } else {
            return fmt.Errorf("circuit breaker is open")
        }
    }
    
    err := fn()
    
    if err != nil {
        cb.failureCount++
        cb.lastFailure = time.Now()
        
        if cb.failureCount >= cb.threshold {
            cb.state = 1 // Open
        }
        return err
    }
    
    // Success
    cb.failureCount = 0
    cb.state = 0 // Closed
    return nil
}
```

### 3. Timeout Pattern

```go
// Timeout pattern for graceful degradation
func timeoutPattern() error {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    resultCh := make(chan error, 1)
    
    go func() {
        resultCh <- slowOperation()
    }()
    
    select {
    case err := <-resultCh:
        return err
    case <-ctx.Done():
        return fmt.Errorf("operation timed out: %w", ctx.Err())
    }
}

func slowOperation() error {
    time.Sleep(2 * time.Second)
    return nil
}
```

---

## 🔄 Retry Mechanisms

### 1. Simple Retry

```go
// Simple retry mechanism
func simpleRetry(fn func() error, maxRetries int) error {
    var lastErr error
    
    for i := 0; i < maxRetries; i++ {
        if err := fn(); err != nil {
            lastErr = err
            log.Printf("Attempt %d failed: %v", i+1, err)
            time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
            continue
        }
        return nil
    }
    
    return fmt.Errorf("max retries exceeded: %w", lastErr)
}
```

### 2. Exponential Backoff

```go
// Exponential backoff retry
func exponentialBackoff(fn func() error, maxRetries int) error {
    var lastErr error
    
    for i := 0; i < maxRetries; i++ {
        if err := fn(); err != nil {
            lastErr = err
            backoff := time.Duration(1<<uint(i)) * 100 * time.Millisecond
            log.Printf("Attempt %d failed, backing off for %v: %v", i+1, backoff, err)
            time.Sleep(backoff)
            continue
        }
        return nil
    }
    
    return fmt.Errorf("max retries exceeded: %w", lastErr)
}
```

### 3. Jittered Backoff

```go
// Jittered backoff to avoid thundering herd
func jitteredBackoff(fn func() error, maxRetries int) error {
    var lastErr error
    
    for i := 0; i < maxRetries; i++ {
        if err := fn(); err != nil {
            lastErr = err
            baseBackoff := time.Duration(1<<uint(i)) * 100 * time.Millisecond
            jitter := time.Duration(rand.Intn(100)) * time.Millisecond
            backoff := baseBackoff + jitter
            
            log.Printf("Attempt %d failed, backing off for %v: %v", i+1, backoff, err)
            time.Sleep(backoff)
            continue
        }
        return nil
    }
    
    return fmt.Errorf("max retries exceeded: %w", lastErr)
}
```

---

## 📊 Error Monitoring and Logging

### 1. Error Logging

```go
// Error logging with context
func logError(err error, context map[string]interface{}) {
    log.Printf("Error occurred: %v, Context: %v", err, context)
}

func processWithLogging() error {
    if err := processData(); err != nil {
        context := map[string]interface{}{
            "operation": "data processing",
            "timestamp": time.Now(),
            "user_id":   "123",
        }
        logError(err, context)
        return err
    }
    return nil
}
```

### 2. Error Metrics

```go
// Error metrics collection
type ErrorMetrics struct {
    errorCount int64
    errorTypes map[string]int64
    mu         sync.Mutex
}

func (em *ErrorMetrics) RecordError(err error) {
    em.mu.Lock()
    defer em.mu.Unlock()
    
    em.errorCount++
    errorType := fmt.Sprintf("%T", err)
    em.errorTypes[errorType]++
}

func (em *ErrorMetrics) GetStats() map[string]interface{} {
    em.mu.Lock()
    defer em.mu.Unlock()
    
    return map[string]interface{}{
        "total_errors": em.errorCount,
        "error_types":  em.errorTypes,
    }
}
```

### 3. Error Alerting

```go
// Error alerting system
type ErrorAlerter struct {
    threshold int
    count     int
    mu        sync.Mutex
}

func (ea *ErrorAlerter) RecordError(err error) {
    ea.mu.Lock()
    defer ea.mu.Unlock()
    
    ea.count++
    if ea.count >= ea.threshold {
        ea.alert(err)
        ea.count = 0 // Reset counter
    }
}

func (ea *ErrorAlerter) alert(err error) {
    log.Printf("ALERT: Error threshold exceeded: %v", err)
    // Send alert to monitoring system
}
```

---

## 🧪 Error Testing Strategies

### 1. Error Injection

```go
// Error injection for testing
type ErrorInjector struct {
    shouldFail bool
    mu         sync.Mutex
}

func (ei *ErrorInjector) SetShouldFail(shouldFail bool) {
    ei.mu.Lock()
    defer ei.mu.Unlock()
    ei.shouldFail = shouldFail
}

func (ei *ErrorInjector) Process() error {
    ei.mu.Lock()
    shouldFail := ei.shouldFail
    ei.mu.Unlock()
    
    if shouldFail {
        return fmt.Errorf("injected error")
    }
    return nil
}
```

### 2. Error Simulation

```go
// Error simulation for testing
func simulateErrors() error {
    // Simulate different error scenarios
    scenarios := []func() error{
        func() error { return fmt.Errorf("network error") },
        func() error { return fmt.Errorf("timeout error") },
        func() error { return fmt.Errorf("validation error") },
        func() error { return nil }, // Success
    }
    
    for i, scenario := range scenarios {
        if err := scenario(); err != nil {
            log.Printf("Scenario %d failed: %v", i+1, err)
        } else {
            log.Printf("Scenario %d succeeded", i+1)
        }
    }
    
    return nil
}
```

### 3. Error Recovery Testing

```go
// Test error recovery mechanisms
func testErrorRecovery() {
    // Test panic recovery
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Panic recovered: %v", r)
        }
    }()
    
    // Test error handling
    if err := processWithError(); err != nil {
        log.Printf("Error handled: %v", err)
    }
}

func processWithError() error {
    return fmt.Errorf("test error")
}
```

---

## 🚀 Advanced Error Patterns

### 1. Error Handler Chain

```go
// Error handler chain pattern
type ErrorHandler interface {
    Handle(err error) error
    SetNext(handler ErrorHandler)
}

type LoggingHandler struct {
    next ErrorHandler
}

func (lh *LoggingHandler) Handle(err error) error {
    log.Printf("Error logged: %v", err)
    if lh.next != nil {
        return lh.next.Handle(err)
    }
    return err
}

func (lh *LoggingHandler) SetNext(handler ErrorHandler) {
    lh.next = handler
}

type RetryHandler struct {
    next ErrorHandler
}

func (rh *RetryHandler) Handle(err error) error {
    // Implement retry logic
    if rh.next != nil {
        return rh.next.Handle(err)
    }
    return err
}

func (rh *RetryHandler) SetNext(handler ErrorHandler) {
    rh.next = handler
}
```

### 2. Error Recovery Strategies

```go
// Error recovery strategies
type RecoveryStrategy interface {
    Recover(err error) error
}

type RetryStrategy struct {
    maxRetries int
}

func (rs *RetryStrategy) Recover(err error) error {
    for i := 0; i < rs.maxRetries; i++ {
        if err := retryOperation(); err == nil {
            return nil
        }
        time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
    }
    return fmt.Errorf("retry failed: %w", err)
}

type FallbackStrategy struct{}

func (fs *FallbackStrategy) Recover(err error) error {
    if err := fallbackOperation(); err != nil {
        return fmt.Errorf("fallback failed: %w", err)
    }
    return nil
}
```

### 3. Error Context Propagation

```go
// Error context propagation
type ErrorContext struct {
    RequestID string
    UserID    string
    Timestamp time.Time
    Err       error
}

func (ec ErrorContext) Error() string {
    return fmt.Sprintf("[%s] %s: %v", ec.RequestID, ec.UserID, ec.Err)
}

func (ec ErrorContext) Unwrap() error {
    return ec.Err
}

func processWithContext() error {
    ctx := ErrorContext{
        RequestID: "req-123",
        UserID:    "user-456",
        Timestamp: time.Now(),
    }
    
    if err := processData(); err != nil {
        ctx.Err = err
        return ctx
    }
    
    return nil
}
```

---

## 🌍 Real-World Applications

### 1. Web Server Error Handling

```go
// Web server with comprehensive error handling
func webServerErrorHandling() {
    http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
        if err := processRequest(r); err != nil {
            log.Printf("Request failed: %v", err)
            
            // Return appropriate HTTP status
            if isTimeoutError(err) {
                http.Error(w, "Request timeout", http.StatusRequestTimeout)
            } else if isValidationError(err) {
                http.Error(w, "Invalid request", http.StatusBadRequest)
            } else {
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
            return
        }
        
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Success"))
    })
}

func processRequest(r *http.Request) error {
    // Process request with error handling
    return nil
}

func isTimeoutError(err error) bool {
    return strings.Contains(err.Error(), "timeout")
}

func isValidationError(err error) bool {
    return strings.Contains(err.Error(), "validation")
}
```

### 2. Database Error Handling

```go
// Database error handling
func databaseErrorHandling() error {
    db, err := sql.Open("postgres", "connection_string")
    if err != nil {
        return fmt.Errorf("failed to open database: %w", err)
    }
    defer db.Close()
    
    // Retry on connection errors
    return retryOnError(func() error {
        return db.Ping()
    }, 3)
}

func retryOnError(fn func() error, maxRetries int) error {
    var lastErr error
    
    for i := 0; i < maxRetries; i++ {
        if err := fn(); err != nil {
            lastErr = err
            if isRetryableError(err) {
                time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
                continue
            }
            return err
        }
        return nil
    }
    
    return fmt.Errorf("max retries exceeded: %w", lastErr)
}

func isRetryableError(err error) bool {
    return strings.Contains(err.Error(), "connection") ||
           strings.Contains(err.Error(), "timeout")
}
```

### 3. Microservice Error Handling

```go
// Microservice error handling
func microserviceErrorHandling() error {
    // Circuit breaker for external service calls
    breaker := &CircuitBreaker{
        threshold: 5,
        timeout:   1 * time.Minute,
    }
    
    // Retry with exponential backoff
    return exponentialBackoff(func() error {
        return breaker.Call(func() error {
            return callExternalService()
        })
    }, 3)
}

func callExternalService() error {
    // Simulate external service call
    return fmt.Errorf("external service unavailable")
}
```

---

## 🎓 Summary

Mastering error handling in concurrent code is essential for building robust, reliable, and maintainable systems. Key takeaways:

1. **Understand error propagation** patterns across goroutines
2. **Implement error aggregation** strategies for multiple operations
3. **Use panic recovery** patterns for fault tolerance
4. **Preserve error context** and wrap errors appropriately
5. **Implement graceful degradation** for fault tolerance
6. **Use circuit breaker** patterns for external dependencies
7. **Implement retry mechanisms** with appropriate backoff strategies
8. **Monitor and log errors** for observability
9. **Test error scenarios** thoroughly
10. **Apply patterns** to real-world applications

Error handling in concurrent code provides the foundation for building resilient systems! 🚀

---

## 🚀 Next Steps

1. **Practice** with the provided examples
2. **Experiment** with different error handling strategies
3. **Apply** patterns to your own projects
4. **Move to the next topic** in the curriculum
5. **Master** advanced error handling techniques

Ready to become an Error Handling expert? Let's dive into the implementation! 💪
