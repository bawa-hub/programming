package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Example 1: Basic Error Propagation
func basicErrorPropagation() {
	fmt.Println("\n1. Basic Error Propagation")
	fmt.Println("=========================")
	
	// Fail Fast vs. Fail Safe
	fmt.Println("  Fail Fast example:")
	if err := failFast(); err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	fmt.Println("  Fail Safe example:")
	if err := failSafe(); err != nil {
		fmt.Printf("    Error: %v\n", err)
	}
	
	fmt.Println("  Basic error propagation completed")
}

func failFast() error {
	if err := validateInput(); err != nil {
		return err // Stop immediately
	}
	return nil
}

func failSafe() error {
	if err := optionalFeature(); err != nil {
		log.Printf("    Optional feature failed: %v", err)
		// Continue with core functionality
	}
	return nil
}

func validateInput() error {
	return fmt.Errorf("validation failed")
}

func optionalFeature() error {
	return fmt.Errorf("optional feature unavailable")
}

// Example 2: Error Context Preservation
func errorContextPreservation() {
	fmt.Println("\n2. Error Context Preservation")
	fmt.Println("============================")
	
	if err := processData(); err != nil {
		fmt.Printf("  Error with context: %v\n", err)
	}
	
	fmt.Println("  Error context preservation completed")
}

func processData() error {
	if err := validateData(); err != nil {
		return fmt.Errorf("data validation failed: %w", err)
	}
	
	if err := transformData(); err != nil {
		return fmt.Errorf("data transformation failed: %w", err)
	}
	
	return nil
}

func validateData() error {
	return fmt.Errorf("invalid data format")
}

func transformData() error {
	return fmt.Errorf("transformation failed")
}

// Example 3: Error Aggregation
func errorAggregation() {
	fmt.Println("\n3. Error Aggregation")
	fmt.Println("===================")
	
	items := []string{"item1", "item2", "item3", "item4", "item5"}
	if err := processMultipleItems(items); err != nil {
		fmt.Printf("  Aggregated error: %v\n", err)
	}
	
	fmt.Println("  Error aggregation completed")
}

func processMultipleItems(items []string) error {
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

func processItem(item string) error {
	if item == "item3" {
		return fmt.Errorf("item %s failed", item)
	}
	return nil
}

// Example 4: Channel-Based Error Propagation
func channelBasedErrorPropagation() {
	fmt.Println("\n4. Channel-Based Error Propagation")
	fmt.Println("==================================")
	
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
			fmt.Printf("  Processed: %d\n", data)
		case err := <-errorCh:
			fmt.Printf("  Error: %v\n", err)
			return
		}
	}
}

// Example 5: Result Pattern
func resultPattern() {
	fmt.Println("\n5. Result Pattern")
	fmt.Println("================")
	
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
			fmt.Printf("  Error: %v\n", result.Err)
			return
		}
		fmt.Printf("  Processed: %v\n", result.Data)
	}
}

type Result struct {
	Data interface{}
	Err  error
}

// Example 6: Error Wrapper Pattern
func errorWrapperPattern() {
	fmt.Println("\n6. Error Wrapper Pattern")
	fmt.Println("=======================")
	
	if err := processWithWrapper(); err != nil {
		fmt.Printf("  Wrapped error: %v\n", err)
	}
	
	fmt.Println("  Error wrapper pattern completed")
}

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
	context := map[string]interface{}{
		"input": "data",
	}
	
	if err := validateInput(); err != nil {
		return ErrorWrapper{
			Operation: "validation",
			Context:   context,
			Err:       err,
		}
	}
	
	return nil
}

// Example 7: Error Collection
func errorCollection() {
	fmt.Println("\n7. Error Collection")
	fmt.Println("==================")
	
	collector := &ErrorCollector{}
	
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := processItem(fmt.Sprintf("item%d", id)); err != nil {
				collector.Add(fmt.Errorf("item %d failed: %w", id, err))
			}
		}(i)
	}
	
	wg.Wait()
	
	if err := collector.Error(); err != nil {
		fmt.Printf("  Collected errors: %v\n", err)
	}
	
	fmt.Println("  Error collection completed")
}

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

// Example 8: Error Group Pattern
func errorGroupPattern() {
	fmt.Println("\n8. Error Group Pattern")
	fmt.Println("=====================")
	
	group := &ErrorGroup{}
	
	for i := 0; i < 5; i++ {
		id := i
		group.Go(func() error {
			return processItem(fmt.Sprintf("item%d", id))
		})
	}
	
	time.Sleep(1 * time.Second) // Wait for completion
	
	if err := group.Wait(); err != nil {
		fmt.Printf("  Group errors: %v\n", err)
	}
	
	fmt.Println("  Error group pattern completed")
}

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

// Example 9: Panic Recovery
func panicRecovery() {
	fmt.Println("\n9. Panic Recovery")
	fmt.Println("================")
	
	// Basic panic recovery
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Recovered from panic: %v\n", r)
		}
	}()
	
	// This will panic
	panic("something went wrong")
}

// Example 10: Goroutine Panic Recovery
func goroutinePanicRecovery() {
	fmt.Println("\n10. Goroutine Panic Recovery")
	fmt.Println("===========================")
	
	ch := make(chan int)
	
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Goroutine recovered from panic: %v\n", r)
			}
		}()
		
		// This will panic
		panic("goroutine panic")
	}()
	
	time.Sleep(100 * time.Millisecond)
	close(ch)
	
	fmt.Println("  Goroutine panic recovery completed")
}

// Example 11: Panic Recovery with Error Channel
func panicRecoveryWithError() {
	fmt.Println("\n11. Panic Recovery with Error Channel")
	fmt.Println("====================================")
	
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
		fmt.Printf("  Error from panic: %v\n", err)
	case <-time.After(1 * time.Second):
		fmt.Println("  No panic occurred")
	}
	
	fmt.Println("  Panic recovery with error channel completed")
}

// Example 12: Panic Recovery Middleware
func panicRecoveryMiddleware() {
	fmt.Println("\n12. Panic Recovery Middleware")
	fmt.Println("============================")
	
	if err := panicRecoveryMiddlewareImpl(func() {
		panic("middleware panic")
	}); err != nil {
		fmt.Printf("  Middleware caught panic: %v\n", err)
	}
	
	fmt.Println("  Panic recovery middleware completed")
}

func panicRecoveryMiddlewareImpl(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	
	fn()
	return nil
}

// Example 13: Fallback Pattern
func fallbackPattern() {
	fmt.Println("\n13. Fallback Pattern")
	fmt.Println("===================")
	
	if err := basicFallbackPatternImpl(); err != nil {
		fmt.Printf("  Fallback failed: %v\n", err)
	} else {
		fmt.Println("  Fallback succeeded")
	}
	
	fmt.Println("  Fallback pattern completed")
}

func basicFallbackPatternImpl() error {
	// Try primary service
	if err := primaryService(); err != nil {
		log.Printf("  Primary service failed: %v", err)
		
		// Try fallback service
		if err := fallbackService(); err != nil {
			log.Printf("  Fallback service failed: %v", err)
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

// Example 14: Circuit Breaker Pattern
func circuitBreakerPattern() {
	fmt.Println("\n14. Circuit Breaker Pattern")
	fmt.Println("==========================")
	
	breaker := &CircuitBreaker{
		threshold: 3,
		timeout:   1 * time.Second,
	}
	
	// Test circuit breaker
	for i := 0; i < 5; i++ {
		err := breaker.Call(func() error {
			if i < 3 {
				return fmt.Errorf("simulated error")
			}
			return nil
		})
		
		if err != nil {
			fmt.Printf("  Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("  Call %d succeeded\n", i)
		}
	}
	
	fmt.Println("  Circuit breaker pattern completed")
}

type CircuitBreaker struct {
	state         int // 0: closed, 1: open, 2: half-open
	failureCount  int
	threshold     int
	timeout       time.Duration
	lastFailure   time.Time
	mu            sync.Mutex
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

// Example 15: Timeout Pattern
func timeoutPattern() {
	fmt.Println("\n15. Timeout Pattern")
	fmt.Println("==================")
	
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	resultCh := make(chan error, 1)
	
	go func() {
		resultCh <- basicSlowOperation()
	}()
	
	select {
	case err := <-resultCh:
		if err != nil {
			fmt.Printf("  Operation failed: %v\n", err)
		} else {
			fmt.Println("  Operation succeeded")
		}
	case <-ctx.Done():
		fmt.Printf("  Operation timed out: %v\n", ctx.Err())
	}
	
	fmt.Println("  Timeout pattern completed")
}

func basicSlowOperation() error {
	time.Sleep(2 * time.Second)
	return nil
}

// Example 16: Simple Retry
func simpleRetry() {
	fmt.Println("\n16. Simple Retry")
	fmt.Println("===============")
	
	if err := simpleRetryImpl(func() error {
		return fmt.Errorf("operation failed")
	}, 3); err != nil {
		fmt.Printf("  Retry failed: %v\n", err)
	}
	
	fmt.Println("  Simple retry completed")
}

func simpleRetryImpl(fn func() error, maxRetries int) error {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		if err := fn(); err != nil {
			lastErr = err
			log.Printf("  Attempt %d failed: %v", i+1, err)
			time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
			continue
		}
		return nil
	}
	
	return fmt.Errorf("max retries exceeded: %w", lastErr)
}

// Example 17: Exponential Backoff
func exponentialBackoff() {
	fmt.Println("\n17. Exponential Backoff")
	fmt.Println("======================")
	
	if err := exponentialBackoffImpl(func() error {
		return fmt.Errorf("operation failed")
	}, 3); err != nil {
		fmt.Printf("  Exponential backoff failed: %v\n", err)
	}
	
	fmt.Println("  Exponential backoff completed")
}

func exponentialBackoffImpl(fn func() error, maxRetries int) error {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		if err := fn(); err != nil {
			lastErr = err
			backoff := time.Duration(1<<uint(i)) * 100 * time.Millisecond
			log.Printf("  Attempt %d failed, backing off for %v: %v", i+1, backoff, err)
			time.Sleep(backoff)
			continue
		}
		return nil
	}
	
	return fmt.Errorf("max retries exceeded: %w", lastErr)
}

// Example 18: Error Logging
func errorLogging() {
	fmt.Println("\n18. Error Logging")
	fmt.Println("================")
	
	if err := processWithLogging(); err != nil {
		fmt.Printf("  Process failed: %v\n", err)
	}
	
	fmt.Println("  Error logging completed")
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

func logError(err error, context map[string]interface{}) {
	log.Printf("  Error occurred: %v, Context: %v", err, context)
}

// Example 19: Error Metrics
func errorMetrics() {
	fmt.Println("\n19. Error Metrics")
	fmt.Println("================")
	
	metrics := NewErrorMetrics()
	
	// Simulate some errors
	metrics.RecordError(fmt.Errorf("validation error"))
	metrics.RecordError(fmt.Errorf("network error"))
	metrics.RecordError(fmt.Errorf("validation error"))
	
	stats := metrics.GetStats()
	fmt.Printf("  Error metrics: %v\n", stats)
	
	fmt.Println("  Error metrics completed")
}

type ErrorMetrics struct {
	errorCount int64
	errorTypes map[string]int64
	mu         sync.Mutex
}

func NewErrorMetrics() *ErrorMetrics {
	return &ErrorMetrics{
		errorTypes: make(map[string]int64),
	}
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

// Example 20: Error Testing
func errorTesting() {
	fmt.Println("\n20. Error Testing")
	fmt.Println("================")
	
	// Test error injection
	injector := &ErrorInjector{}
	
	// Test normal operation
	injector.SetShouldFail(false)
	if err := injector.Process(); err != nil {
		fmt.Printf("  Normal operation failed: %v\n", err)
	} else {
		fmt.Println("  Normal operation succeeded")
	}
	
	// Test error injection
	injector.SetShouldFail(true)
	if err := injector.Process(); err != nil {
		fmt.Printf("  Injected error: %v\n", err)
	}
	
	fmt.Println("  Error testing completed")
}

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

// Run all basic examples
func runBasicExamples() {
	fmt.Println("⚠️ Error Handling in Concurrent Code Examples")
	fmt.Println("=============================================")
	
	basicErrorPropagation()
	errorContextPreservation()
	errorAggregation()
	channelBasedErrorPropagation()
	resultPattern()
	errorWrapperPattern()
	errorCollection()
	errorGroupPattern()
	panicRecovery()
	goroutinePanicRecovery()
	panicRecoveryWithError()
	panicRecoveryMiddleware()
	fallbackPattern()
	circuitBreakerPattern()
	timeoutPattern()
	simpleRetry()
	exponentialBackoff()
	errorLogging()
	errorMetrics()
	errorTesting()
}
