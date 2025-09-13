package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// 🎯 CONTEXT BEST PRACTICES
// Understanding context best practices and performance considerations

func main() {
	fmt.Println("🎯 CONTEXT BEST PRACTICES")
	fmt.Println("=========================")

	// 1. Context Value Storage Guidelines
	fmt.Println("\n1. Context Value Storage Guidelines:")
	contextValueStorageGuidelines()

	// 2. Performance Considerations
	fmt.Println("\n2. Performance Considerations:")
	performanceConsiderations()

	// 3. Testing with Contexts
	fmt.Println("\n3. Testing with Contexts:")
	testingWithContexts()

	// 4. Context vs Other Patterns
	fmt.Println("\n4. Context vs Other Patterns:")
	contextVsOtherPatterns()

	// 5. Context Trees and Inheritance
	fmt.Println("\n5. Context Trees and Inheritance:")
	contextTreesAndInheritance()

	// 6. Error Handling with Context
	fmt.Println("\n6. Error Handling with Context:")
	errorHandlingWithContext()

	// 7. Context in HTTP Handlers
	fmt.Println("\n7. Context in HTTP Handlers:")
	contextInHTTPHandlers()

	// 8. Advanced Context Patterns
	fmt.Println("\n8. Advanced Context Patterns:")
	advancedContextPatterns()
}

// CONTEXT VALUE STORAGE GUIDELINES: Understanding value storage best practices
func contextValueStorageGuidelines() {
	fmt.Println("Understanding context value storage guidelines...")
	
	// 1. Use typed keys
	fmt.Println("  📝 Best Practice 1: Use typed keys")
	useTypedKeys()
	
	// 2. Avoid storing large values
	fmt.Println("  📝 Best Practice 2: Avoid storing large values")
	avoidStoringLargeValues()
	
	// 3. Use context for request-scoped data
	fmt.Println("  📝 Best Practice 3: Use context for request-scoped data")
	useContextForRequestScopedData()
	
	// 4. Don't store optional data in context
	fmt.Println("  📝 Best Practice 4: Don't store optional data in context")
	dontStoreOptionalDataInContext()
}

func useTypedKeys() {
	// Good: Use typed keys
	type userKey struct{}
	type requestIDKey struct{}
	
	ctx := context.Background()
	ctx = context.WithValue(ctx, userKey{}, "john_doe")
	ctx = context.WithValue(ctx, requestIDKey{}, "req-123")
	
	user := ctx.Value(userKey{}).(string)
	requestID := ctx.Value(requestIDKey{}).(string)
	
	fmt.Printf("    ✅ User: %s, RequestID: %s\n", user, requestID)
}

func avoidStoringLargeValues() {
	// Bad: Storing large values
	largeData := make([]byte, 1024*1024) // 1MB
	_ = context.WithValue(context.Background(), "large_data", largeData)
	
	// Good: Store reference or ID
	ctx := context.WithValue(context.Background(), "data_id", "large_data_123")
	_ = ctx
	
	fmt.Printf("    ✅ Stored data ID instead of large data\n")
}

func useContextForRequestScopedData() {
	// Good: Use context for request-scoped data
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", "123")
	ctx = context.WithValue(ctx, "request_id", "req-456")
	
	// Pass context through function calls
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("user_id").(string)
	requestID := ctx.Value("request_id").(string)
	
	fmt.Printf("    ✅ Processing request %s for user %s\n", requestID, userID)
}

func dontStoreOptionalDataInContext() {
	// Bad: Storing optional data
	// ctx := context.WithValue(context.Background(), "optional_config", "some_value")
	
	// Good: Use function parameters for optional data
	processWithOptionalConfig(context.Background(), "some_value")
}

func processWithOptionalConfig(ctx context.Context, config string) {
	fmt.Printf("    ✅ Using function parameter for optional config: %s\n", config)
}

// PERFORMANCE CONSIDERATIONS: Understanding performance implications
func performanceConsiderations() {
	fmt.Println("Understanding performance considerations...")
	
	// 1. Context creation overhead
	fmt.Println("  📝 Performance 1: Context creation overhead")
	contextCreationOverhead()
	
	// 2. Value lookup performance
	fmt.Println("  📝 Performance 2: Value lookup performance")
	valueLookupPerformance()
	
	// 3. Context cancellation performance
	fmt.Println("  📝 Performance 3: Context cancellation performance")
	contextCancellationPerformance()
	
	// 4. Memory usage
	fmt.Println("  📝 Performance 4: Memory usage")
	memoryUsage()
}

func contextCreationOverhead() {
	// Test context creation overhead
	start := time.Now()
	
	for i := 0; i < 1000000; i++ {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "key", "value")
		ctx, _ = context.WithTimeout(ctx, time.Second)
		_ = ctx
	}
	
	duration := time.Since(start)
	fmt.Printf("    📊 Context creation time: %v\n", duration)
}

func valueLookupPerformance() {
	ctx := context.Background()
	for i := 0; i < 1000; i++ {
		ctx = context.WithValue(ctx, fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	
	// Test value lookup performance
	start := time.Now()
	
	for i := 0; i < 1000000; i++ {
		_ = ctx.Value("key500")
	}
	
	duration := time.Since(start)
	fmt.Printf("    📊 Value lookup time: %v\n", duration)
}

func contextCancellationPerformance() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	
	// Test cancellation performance
	start := time.Now()
	
	for i := 0; i < 1000000; i++ {
		select {
		case <-ctx.Done():
			break
		default:
		}
	}
	
	duration := time.Since(start)
	fmt.Printf("    📊 Cancellation check time: %v\n", duration)
}

func memoryUsage() {
	// Test memory usage with many contexts
	contexts := make([]context.Context, 10000)
	
	for i := 0; i < 10000; i++ {
		ctx := context.Background()
		for j := 0; j < 10; j++ {
			ctx = context.WithValue(ctx, fmt.Sprintf("key%d", j), fmt.Sprintf("value%d", j))
		}
		contexts[i] = ctx
	}
	
	fmt.Printf("    📊 Created %d contexts with values\n", len(contexts))
}

// TESTING WITH CONTEXTS: Understanding testing patterns
func testingWithContexts() {
	fmt.Println("Understanding testing with contexts...")
	
	// 1. Test with timeout
	fmt.Println("  📝 Test 1: Test with timeout")
	testWithTimeout()
	
	// 2. Test with cancellation
	fmt.Println("  📝 Test 2: Test with cancellation")
	testWithCancellation()
	
	// 3. Test with values
	fmt.Println("  📝 Test 3: Test with values")
	testWithValues()
	
	// 4. Test context propagation
	fmt.Println("  📝 Test 4: Test context propagation")
	testContextPropagation()
}

func testWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	// Test function that respects timeout
	err := longRunningFunction(ctx)
	if err != nil {
		fmt.Printf("    ✅ Timeout test passed: %v\n", err)
	} else {
		fmt.Printf("    ❌ Timeout test failed\n")
	}
}

func testWithCancellation() {
	ctx, cancel := context.WithCancel(context.Background())
	
	// Cancel after short delay
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	
	err := longRunningFunction(ctx)
	if err != nil {
		fmt.Printf("    ✅ Cancellation test passed: %v\n", err)
	} else {
		fmt.Printf("    ❌ Cancellation test failed\n")
	}
}

func testWithValues() {
	ctx := context.WithValue(context.Background(), "test_key", "test_value")
	
	value := ctx.Value("test_key").(string)
	if value == "test_value" {
		fmt.Printf("    ✅ Value test passed: %s\n", value)
	} else {
		fmt.Printf("    ❌ Value test failed\n")
	}
}

func testContextPropagation() {
	ctx := context.WithValue(context.Background(), "propagated_key", "propagated_value")
	
	// Test context propagation through function calls
	result := propagateContext(ctx)
	if result == "propagated_value" {
		fmt.Printf("    ✅ Context propagation test passed: %s\n", result)
	} else {
		fmt.Printf("    ❌ Context propagation test failed\n")
	}
}

func longRunningFunction(ctx context.Context) error {
	select {
	case <-time.After(200 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func propagateContext(ctx context.Context) string {
	return ctx.Value("propagated_key").(string)
}

// CONTEXT VS OTHER PATTERNS: Understanding when to use context
func contextVsOtherPatterns() {
	fmt.Println("Understanding context vs other patterns...")
	
	// 1. Context vs channels
	fmt.Println("  📝 Pattern 1: Context vs channels")
	contextVsChannels()
	
	// 2. Context vs global variables
	fmt.Println("  📝 Pattern 2: Context vs global variables")
	contextVsGlobalVariables()
	
	// 3. Context vs function parameters
	fmt.Println("  📝 Pattern 3: Context vs function parameters")
	contextVsFunctionParameters()
}

func contextVsChannels() {
	// Use context for cancellation
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	
	// Use channels for data flow
	dataCh := make(chan string, 1)
	
	go func() {
		select {
		case dataCh <- "data":
		case <-ctx.Done():
			return
		}
	}()
	
	select {
	case data := <-dataCh:
		fmt.Printf("    ✅ Received data: %s\n", data)
	case <-ctx.Done():
		fmt.Printf("    ✅ Context cancelled: %v\n", ctx.Err())
	}
}

func contextVsGlobalVariables() {
	// Bad: Using global variables
	// globalUserID = "123"
	
	// Good: Using context
	ctx := context.WithValue(context.Background(), "user_id", "123")
	processUser(ctx)
}

func processUser(ctx context.Context) {
	userID := ctx.Value("user_id").(string)
	fmt.Printf("    ✅ Processing user: %s\n", userID)
}

func contextVsFunctionParameters() {
	// Use context for cancellation and request-scoped data
	ctx := context.WithValue(context.Background(), "request_id", "req-123")
	
	// Use function parameters for business logic
	processBusinessLogic(ctx, "business_data", 42)
}

func processBusinessLogic(ctx context.Context, data string, count int) {
	requestID := ctx.Value("request_id").(string)
	fmt.Printf("    ✅ Processing %s (count: %d) for request %s\n", data, count, requestID)
}

// CONTEXT TREES AND INHERITANCE: Understanding context trees
func contextTreesAndInheritance() {
	fmt.Println("Understanding context trees and inheritance...")
	
	// Create context tree
	rootCtx := context.Background()
	childCtx := context.WithValue(rootCtx, "level", "child")
	grandchildCtx := context.WithValue(childCtx, "level", "grandchild")
	
	// Test inheritance
	fmt.Printf("  📊 Root context level: %v\n", rootCtx.Value("level"))
	fmt.Printf("  📊 Child context level: %v\n", childCtx.Value("level"))
	fmt.Printf("  📊 Grandchild context level: %v\n", grandchildCtx.Value("level"))
	
	// Test cancellation inheritance
	testCancellationInheritance()
}

func testCancellationInheritance() {
	rootCtx, rootCancel := context.WithCancel(context.Background())
	childCtx, childCancel := context.WithCancel(rootCtx)
	
	// Cancel child context
	childCancel()
	
	// Check if child is cancelled
	if childCtx.Err() != nil {
		fmt.Printf("  📊 Child context cancelled: %v\n", childCtx.Err())
	}
	
	// Check if root is still active
	if rootCtx.Err() == nil {
		fmt.Printf("  📊 Root context still active\n")
	}
	
	// Cancel root context
	rootCancel()
	
	// Check if root is cancelled
	if rootCtx.Err() != nil {
		fmt.Printf("  📊 Root context cancelled: %v\n", rootCtx.Err())
	}
}

// ERROR HANDLING WITH CONTEXT: Understanding error handling
func errorHandlingWithContext() {
	fmt.Println("Understanding error handling with context...")
	
	// 1. Context cancellation errors
	fmt.Println("  📝 Error 1: Context cancellation errors")
	contextCancellationErrors()
	
	// 2. Context timeout errors
	fmt.Println("  📝 Error 2: Context timeout errors")
	contextTimeoutErrors()
	
	// 3. Error propagation
	fmt.Println("  📝 Error 3: Error propagation")
	errorPropagation()
}

func contextCancellationErrors() {
	ctx, cancel := context.WithCancel(context.Background())
	
	// Cancel context
	cancel()
	
	// Check error
	if err := ctx.Err(); err != nil {
		fmt.Printf("    ✅ Context cancelled: %v\n", err)
	}
}

func contextTimeoutErrors() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	
	// Wait for timeout
	time.Sleep(2 * time.Millisecond)
	
	// Check error
	if err := ctx.Err(); err != nil {
		fmt.Printf("    ✅ Context timed out: %v\n", err)
	}
}

func errorPropagation() {
	ctx := context.Background()
	
	// Simulate error propagation
	err := processWithError(ctx)
	if err != nil {
		fmt.Printf("    ✅ Error propagated: %v\n", err)
	}
}

func processWithError(ctx context.Context) error {
	// Simulate error
	return fmt.Errorf("processing error")
}

// CONTEXT IN HTTP HANDLERS: Understanding HTTP context usage
func contextInHTTPHandlers() {
	fmt.Println("Understanding context in HTTP handlers...")
	
	// 1. Request context
	fmt.Println("  📝 HTTP 1: Request context")
	requestContext()
	
	// 2. Timeout handling
	fmt.Println("  📝 HTTP 2: Timeout handling")
	timeoutHandling()
	
	// 3. Middleware context
	fmt.Println("  📝 HTTP 3: Middleware context")
	middlewareContext()
}

func requestContext() {
	// Simulate HTTP request context
	req, _ := http.NewRequest("GET", "/api/users", nil)
	ctx := req.Context()
	
	// Add request-scoped data
	ctx = context.WithValue(ctx, "user_id", "123")
	ctx = context.WithValue(ctx, "request_id", "req-456")
	
	// Process request
	processHTTPRequest(ctx)
}

func processHTTPRequest(ctx context.Context) {
	userID := ctx.Value("user_id").(string)
	requestID := ctx.Value("request_id").(string)
	
	fmt.Printf("    ✅ Processing HTTP request %s for user %s\n", requestID, userID)
}

func timeoutHandling() {
	// Simulate HTTP request with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	
	// Process with timeout
	err := processWithTimeout(ctx)
	if err != nil {
		fmt.Printf("    ✅ Timeout handled: %v\n", err)
	}
}

func processWithTimeout(ctx context.Context) error {
	select {
	case <-time.After(200 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func middlewareContext() {
	// Simulate middleware adding context
	ctx := context.Background()
	ctx = addUserMiddleware(ctx, "john_doe")
	ctx = addRequestIDMiddleware(ctx, "req-789")
	
	// Process with middleware context
	processWithMiddleware(ctx)
}

func addUserMiddleware(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, "user_id", userID)
}

func addRequestIDMiddleware(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, "request_id", requestID)
}

func processWithMiddleware(ctx context.Context) {
	userID := ctx.Value("user_id").(string)
	requestID := ctx.Value("request_id").(string)
	
	fmt.Printf("    ✅ Processing with middleware: user=%s, request=%s\n", userID, requestID)
}

// ADVANCED CONTEXT PATTERNS: Understanding advanced patterns
func advancedContextPatterns() {
	fmt.Println("Understanding advanced context patterns...")
	
	// 1. Context with deadline
	fmt.Println("  📝 Pattern 1: Context with deadline")
	contextWithDeadline()
	
	// 2. Context with multiple timeouts
	fmt.Println("  📝 Pattern 2: Context with multiple timeouts")
	contextWithMultipleTimeouts()
	
	// 3. Context with retry logic
	fmt.Println("  📝 Pattern 3: Context with retry logic")
	contextWithRetryLogic()
	
	// 4. Context with circuit breaker
	fmt.Println("  📝 Pattern 4: Context with circuit breaker")
	contextWithCircuitBreaker()
}

func contextWithDeadline() {
	deadline := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	// Process with deadline
	err := processWithDeadline(ctx)
	if err != nil {
		fmt.Printf("    ✅ Deadline handled: %v\n", err)
	}
}

func processWithDeadline(ctx context.Context) error {
	select {
	case <-time.After(200 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func contextWithMultipleTimeouts() {
	// Create context with multiple timeout layers
	ctx1, cancel1 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel1()
	
	ctx2, cancel2 := context.WithTimeout(ctx1, 100*time.Millisecond)
	defer cancel2()
	
	// Process with multiple timeouts
	err := processWithMultipleTimeouts(ctx2)
	if err != nil {
		fmt.Printf("    ✅ Multiple timeouts handled: %v\n", err)
	}
}

func processWithMultipleTimeouts(ctx context.Context) error {
	select {
	case <-time.After(150 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func contextWithRetryLogic() {
	ctx := context.Background()
	
	// Retry with exponential backoff
	err := retryWithContext(ctx, 3)
	if err != nil {
		fmt.Printf("    ✅ Retry logic handled: %v\n", err)
	}
}

func retryWithContext(ctx context.Context, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		err := processWithRetry(ctx)
		if err == nil {
			return nil
		}
		
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		
		// Exponential backoff
		backoff := time.Duration(i+1) * 100 * time.Millisecond
		time.Sleep(backoff)
	}
	
	return fmt.Errorf("max retries exceeded")
}

func processWithRetry(ctx context.Context) error {
	// Simulate random failure
	if time.Now().UnixNano()%2 == 0 {
		return fmt.Errorf("random failure")
	}
	return nil
}

func contextWithCircuitBreaker() {
	ctx := context.Background()
	
	// Simulate circuit breaker pattern
	err := processWithCircuitBreaker(ctx)
	if err != nil {
		fmt.Printf("    ✅ Circuit breaker handled: %v\n", err)
	}
}

func processWithCircuitBreaker(ctx context.Context) error {
	// Simulate circuit breaker logic
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Process normally
		return nil
	}
}
