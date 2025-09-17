package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Advanced Pattern 1: Error Handler Chain
type ErrorHandler interface {
	Handle(err error) error
	SetNext(handler ErrorHandler)
}

type LoggingHandler struct {
	next ErrorHandler
}

func (lh *LoggingHandler) Handle(err error) error {
	log.Printf("  Error logged: %v", err)
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

// Advanced Pattern 2: Error Recovery Strategies
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
	if err := advancedFallbackOperation(); err != nil {
		return fmt.Errorf("fallback failed: %w", err)
	}
	return nil
}

func retryOperation() error {
	return fmt.Errorf("retry operation failed")
}

func advancedFallbackOperation() error {
	return nil // Fallback succeeds
}

// Advanced Pattern 3: Error Context Propagation
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

// Advanced Pattern 4: Error Monitoring System
type ErrorMonitor struct {
	errors    chan error
	alerts    chan string
	metrics   *ErrorMetrics
	threshold int
	mu        sync.Mutex
	count     int
}

func NewErrorMonitor(threshold int) *ErrorMonitor {
	monitor := &ErrorMonitor{
		errors:    make(chan error, 100),
		alerts:    make(chan string, 10),
		metrics:   NewErrorMetrics(),
		threshold: threshold,
	}
	
	go monitor.run()
	return monitor
}

func (em *ErrorMonitor) run() {
	for err := range em.errors {
		em.metrics.RecordError(err)
		em.checkThreshold(err)
	}
}

func (em *ErrorMonitor) checkThreshold(err error) {
	em.mu.Lock()
	defer em.mu.Unlock()
	
	em.count++
	if em.count >= em.threshold {
		em.alerts <- fmt.Sprintf("Error threshold exceeded: %v", err)
		em.count = 0 // Reset counter
	}
}

func (em *ErrorMonitor) RecordError(err error) {
	select {
	case em.errors <- err:
	default:
		// Channel is full, drop error
	}
}

func (em *ErrorMonitor) GetAlerts() <-chan string {
	return em.alerts
}

// Advanced Pattern 5: Error Recovery Manager
type ErrorRecoveryManager struct {
	strategies map[string]RecoveryStrategy
	mu         sync.RWMutex
}

func NewErrorRecoveryManager() *ErrorRecoveryManager {
	return &ErrorRecoveryManager{
		strategies: make(map[string]RecoveryStrategy),
	}
}

func (erm *ErrorRecoveryManager) RegisterStrategy(name string, strategy RecoveryStrategy) {
	erm.mu.Lock()
	defer erm.mu.Unlock()
	erm.strategies[name] = strategy
}

func (erm *ErrorRecoveryManager) Recover(name string, err error) error {
	erm.mu.RLock()
	strategy, exists := erm.strategies[name]
	erm.mu.RUnlock()
	
	if !exists {
		return fmt.Errorf("strategy %s not found", name)
	}
	
	return strategy.Recover(err)
}

// Advanced Pattern 6: Error Circuit Breaker with Metrics
type AdvancedCircuitBreaker struct {
	state         int // 0: closed, 1: open, 2: half-open
	failureCount  int
	successCount  int
	threshold     int
	timeout       time.Duration
	lastFailure   time.Time
	metrics       *ErrorMetrics
	mu            sync.Mutex
}

func NewAdvancedCircuitBreaker(threshold int, timeout time.Duration) *AdvancedCircuitBreaker {
	return &AdvancedCircuitBreaker{
		threshold: threshold,
		timeout:   timeout,
		metrics:   NewErrorMetrics(),
	}
}

func (acb *AdvancedCircuitBreaker) Call(fn func() error) error {
	acb.mu.Lock()
	defer acb.mu.Unlock()
	
	if acb.state == 1 { // Open
		if time.Since(acb.lastFailure) > acb.timeout {
			acb.state = 2 // Half-open
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}
	
	err := fn()
	
	if err != nil {
		acb.failureCount++
		acb.lastFailure = time.Now()
		acb.metrics.RecordError(err)
		
		if acb.failureCount >= acb.threshold {
			acb.state = 1 // Open
		}
		return err
	}
	
	// Success
	acb.successCount++
	acb.failureCount = 0
	acb.state = 0 // Closed
	return nil
}

func (acb *AdvancedCircuitBreaker) GetStats() map[string]interface{} {
	acb.mu.Lock()
	defer acb.mu.Unlock()
	
	return map[string]interface{}{
		"state":         acb.state,
		"failure_count": acb.failureCount,
		"success_count": acb.successCount,
		"threshold":     acb.threshold,
	}
}

// Advanced Pattern 7: Error Rate Limiter
type ErrorRateLimiter struct {
	errors    chan error
	rate      time.Duration
	burst     int
	mu        sync.Mutex
	lastError time.Time
	count     int
}

func NewErrorRateLimiter(rate time.Duration, burst int) *ErrorRateLimiter {
	limiter := &ErrorRateLimiter{
		errors: make(chan error, burst),
		rate:   rate,
		burst:  burst,
	}
	
	go limiter.run()
	return limiter
}

func (erl *ErrorRateLimiter) run() {
	ticker := time.NewTicker(erl.rate)
	defer ticker.Stop()
	
	for range ticker.C {
		erl.mu.Lock()
		erl.count = 0
		erl.mu.Unlock()
	}
}

func (erl *ErrorRateLimiter) RecordError(err error) bool {
	erl.mu.Lock()
	defer erl.mu.Unlock()
	
	if erl.count >= erl.burst {
		return false // Rate limit exceeded
	}
	
	erl.count++
	erl.lastError = time.Now()
	
	select {
	case erl.errors <- err:
		return true
	default:
		return false
	}
}

func (erl *ErrorRateLimiter) GetErrors() <-chan error {
	return erl.errors
}

// Advanced Pattern 8: Error Correlation
type ErrorCorrelation struct {
	errors map[string][]error
	mu     sync.Mutex
}

func NewErrorCorrelation() *ErrorCorrelation {
	return &ErrorCorrelation{
		errors: make(map[string][]error),
	}
}

func (ec *ErrorCorrelation) RecordError(correlationID string, err error) {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	
	ec.errors[correlationID] = append(ec.errors[correlationID], err)
}

func (ec *ErrorCorrelation) GetErrors(correlationID string) []error {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	
	return ec.errors[correlationID]
}

func (ec *ErrorCorrelation) GetCorrelationStats() map[string]int {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	
	stats := make(map[string]int)
	for id, errors := range ec.errors {
		stats[id] = len(errors)
	}
	return stats
}

// Advanced Pattern 9: Error Recovery with Backoff
type ErrorRecoveryWithBackoff struct {
	baseDelay time.Duration
	maxDelay  time.Duration
	multiplier float64
	mu        sync.Mutex
	delays    map[string]time.Duration
}

func NewErrorRecoveryWithBackoff(baseDelay, maxDelay time.Duration, multiplier float64) *ErrorRecoveryWithBackoff {
	return &ErrorRecoveryWithBackoff{
		baseDelay:  baseDelay,
		maxDelay:   maxDelay,
		multiplier: multiplier,
		delays:     make(map[string]time.Duration),
	}
}

func (erb *ErrorRecoveryWithBackoff) Recover(operationID string, fn func() error) error {
	erb.mu.Lock()
	delay, exists := erb.delays[operationID]
	if !exists {
		delay = erb.baseDelay
	}
	erb.mu.Unlock()
	
	if err := fn(); err != nil {
		// Increase delay for next retry
		erb.mu.Lock()
		erb.delays[operationID] = time.Duration(float64(delay) * erb.multiplier)
		if erb.delays[operationID] > erb.maxDelay {
			erb.delays[operationID] = erb.maxDelay
		}
		erb.mu.Unlock()
		
		time.Sleep(delay)
		return err
	}
	
	// Reset delay on success
	erb.mu.Lock()
	erb.delays[operationID] = erb.baseDelay
	erb.mu.Unlock()
	
	return nil
}

// Advanced Pattern 10: Error Context Chain
type ErrorContextChain struct {
	contexts []map[string]interface{}
	mu       sync.Mutex
}

func NewErrorContextChain() *ErrorContextChain {
	return &ErrorContextChain{
		contexts: make([]map[string]interface{}, 0),
	}
}

func (ecc *ErrorContextChain) AddContext(ctx map[string]interface{}) {
	ecc.mu.Lock()
	defer ecc.mu.Unlock()
	ecc.contexts = append(ecc.contexts, ctx)
}

func (ecc *ErrorContextChain) WrapError(err error) error {
	ecc.mu.Lock()
	defer ecc.mu.Unlock()
	
	contextStr := ""
	for i, ctx := range ecc.contexts {
		if i > 0 {
			contextStr += " -> "
		}
		contextStr += fmt.Sprintf("%v", ctx)
	}
	
	return fmt.Errorf("[%s] %w", contextStr, err)
}

// Advanced Pattern 1: Error Handler Chain
func errorHandlerChain() {
	fmt.Println("\n1. Error Handler Chain")
	fmt.Println("=====================")
	
	// Create handler chain
	loggingHandler := &LoggingHandler{}
	retryHandler := &RetryHandler{}
	
	loggingHandler.SetNext(retryHandler)
	
	// Test handler chain
	err := fmt.Errorf("test error")
	if err := loggingHandler.Handle(err); err != nil {
		fmt.Printf("  Handler chain result: %v\n", err)
	}
	
	fmt.Println("Error handler chain completed")
}

// Advanced Pattern 2: Error Recovery Strategies
func errorRecoveryStrategies() {
	fmt.Println("\n2. Error Recovery Strategies")
	fmt.Println("============================")
	
	manager := NewErrorRecoveryManager()
	
	// Register strategies
	manager.RegisterStrategy("retry", &RetryStrategy{maxRetries: 3})
	manager.RegisterStrategy("fallback", &FallbackStrategy{})
	
	// Test strategies
	err := fmt.Errorf("operation failed")
	
	if err := manager.Recover("retry", err); err != nil {
		fmt.Printf("  Retry strategy failed: %v\n", err)
	}
	
	if err := manager.Recover("fallback", err); err != nil {
		fmt.Printf("  Fallback strategy failed: %v\n", err)
	} else {
		fmt.Println("  Fallback strategy succeeded")
	}
	
	fmt.Println("Error recovery strategies completed")
}

// Advanced Pattern 3: Error Context Propagation
func errorContextPropagation() {
	fmt.Println("\n3. Error Context Propagation")
	fmt.Println("============================")
	
	if err := processWithContextPropagation(); err != nil {
		fmt.Printf("  Contextual error: %v\n", err)
	}
	
	fmt.Println("Error context propagation completed")
}

func processWithContextPropagation() error {
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

// Advanced Pattern 4: Error Monitoring System
func errorMonitoringSystem() {
	fmt.Println("\n4. Error Monitoring System")
	fmt.Println("==========================")
	
	monitor := NewErrorMonitor(3)
	
	// Simulate errors
	for i := 0; i < 5; i++ {
		monitor.RecordError(fmt.Errorf("error %d", i))
	}
	
	// Check for alerts
	select {
	case alert := <-monitor.GetAlerts():
		fmt.Printf("  Alert: %s\n", alert)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  No alerts")
	}
	
	fmt.Println("Error monitoring system completed")
}

// Advanced Pattern 5: Advanced Circuit Breaker
func advancedCircuitBreaker() {
	fmt.Println("\n5. Advanced Circuit Breaker")
	fmt.Println("==========================")
	
	breaker := NewAdvancedCircuitBreaker(3, 1*time.Second)
	
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
	
	stats := breaker.GetStats()
	fmt.Printf("  Circuit breaker stats: %v\n", stats)
	
	fmt.Println("Advanced circuit breaker completed")
}

// Advanced Pattern 6: Error Rate Limiter
func errorRateLimiter() {
	fmt.Println("\n6. Error Rate Limiter")
	fmt.Println("====================")
	
	limiter := NewErrorRateLimiter(100*time.Millisecond, 3)
	
	// Simulate errors
	for i := 0; i < 5; i++ {
		success := limiter.RecordError(fmt.Errorf("error %d", i))
		if success {
			fmt.Printf("  Error %d recorded\n", i)
		} else {
			fmt.Printf("  Error %d rate limited\n", i)
		}
	}
	
	fmt.Println("Error rate limiter completed")
}

// Advanced Pattern 7: Error Correlation
func errorCorrelation() {
	fmt.Println("\n7. Error Correlation")
	fmt.Println("===================")
	
	correlation := NewErrorCorrelation()
	
	// Simulate correlated errors
	correlation.RecordError("req-123", fmt.Errorf("validation error"))
	correlation.RecordError("req-123", fmt.Errorf("network error"))
	correlation.RecordError("req-456", fmt.Errorf("timeout error"))
	
	errors := correlation.GetErrors("req-123")
	fmt.Printf("  Errors for req-123: %v\n", errors)
	
	stats := correlation.GetCorrelationStats()
	fmt.Printf("  Correlation stats: %v\n", stats)
	
	fmt.Println("Error correlation completed")
}

// Advanced Pattern 8: Error Recovery with Backoff
func errorRecoveryWithBackoff() {
	fmt.Println("\n8. Error Recovery with Backoff")
	fmt.Println("=============================")
	
	recovery := NewErrorRecoveryWithBackoff(100*time.Millisecond, 1*time.Second, 2.0)
	
	// Test recovery with backoff
	for i := 0; i < 3; i++ {
		err := recovery.Recover("operation-1", func() error {
			return fmt.Errorf("operation failed")
		})
		
		if err != nil {
			fmt.Printf("  Recovery attempt %d failed: %v\n", i+1, err)
		}
	}
	
	fmt.Println("Error recovery with backoff completed")
}

// Advanced Pattern 9: Error Context Chain
func errorContextChain() {
	fmt.Println("\n9. Error Context Chain")
	fmt.Println("=====================")
	
	chain := NewErrorContextChain()
	
	// Add contexts
	chain.AddContext(map[string]interface{}{"service": "user-service"})
	chain.AddContext(map[string]interface{}{"operation": "create-user"})
	chain.AddContext(map[string]interface{}{"user_id": "123"})
	
	// Wrap error with context chain
	err := fmt.Errorf("database error")
	wrappedErr := chain.WrapError(err)
	fmt.Printf("  Wrapped error: %v\n", wrappedErr)
	
	fmt.Println("Error context chain completed")
}

// Advanced Pattern 10: Web Server Error Handling
func webServerErrorHandling() {
	fmt.Println("\n10. Web Server Error Handling")
	fmt.Println("============================")
	
	// Simulate web server error handling
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		if err := processRequest(r); err != nil {
			log.Printf("  Request failed: %v", err)
			
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
	
	fmt.Println("Web server error handling completed")
}

func processRequest(r *http.Request) error {
	// Simulate request processing
	return nil
}

func isTimeoutError(err error) bool {
	return strings.Contains(err.Error(), "timeout")
}

func isValidationError(err error) bool {
	return strings.Contains(err.Error(), "validation")
}

// Advanced Pattern 11: Database Error Handling
func databaseErrorHandling() {
	fmt.Println("\n11. Database Error Handling")
	fmt.Println("==========================")
	
	// Simulate database error handling
	if err := databaseOperation(); err != nil {
		fmt.Printf("  Database operation failed: %v\n", err)
	}
	
	fmt.Println("Database error handling completed")
}

func databaseOperation() error {
	// Simulate database operation
	return fmt.Errorf("connection timeout")
}

// Advanced Pattern 12: Microservice Error Handling
func microserviceErrorHandling() {
	fmt.Println("\n12. Microservice Error Handling")
	fmt.Println("==============================")
	
	// Circuit breaker for external service calls
	breaker := NewAdvancedCircuitBreaker(5, 1*time.Minute)
	
	// Retry with exponential backoff
	if err := advancedRetryWithBackoff(func() error {
		return breaker.Call(func() error {
			return callExternalService()
		})
	}, 3); err != nil {
		fmt.Printf("  Microservice call failed: %v\n", err)
	}
	
	fmt.Println("Microservice error handling completed")
}

func callExternalService() error {
	// Simulate external service call
	return fmt.Errorf("external service unavailable")
}

func advancedRetryWithBackoff(fn func() error, maxRetries int) error {
	var lastErr error
	
	for i := 0; i < maxRetries; i++ {
		if err := fn(); err != nil {
			lastErr = err
			backoff := time.Duration(1<<uint(i)) * 100 * time.Millisecond
			time.Sleep(backoff)
			continue
		}
		return nil
	}
	
	return fmt.Errorf("max retries exceeded: %w", lastErr)
}

// Run all advanced patterns
func runAdvancedPatterns() {
	fmt.Println("🚀 Advanced Error Handling Patterns")
	fmt.Println("===================================")
	
	errorHandlerChain()
	errorRecoveryStrategies()
	errorContextPropagation()
	errorMonitoringSystem()
	advancedCircuitBreaker()
	errorRateLimiter()
	errorCorrelation()
	errorRecoveryWithBackoff()
	errorContextChain()
	webServerErrorHandling()
	databaseErrorHandling()
	microserviceErrorHandling()
	
	fmt.Println("\n🎉 All advanced patterns completed!")
	fmt.Println("Ready to move to the next topic!")
}
