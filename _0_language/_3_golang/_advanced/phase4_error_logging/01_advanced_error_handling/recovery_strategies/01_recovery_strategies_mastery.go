// 🔄 ERROR RECOVERY STRATEGIES MASTERY
// Advanced error recovery, retry strategies, and resilience patterns
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

// ============================================================================
// RETRY STRATEGIES
// ============================================================================

// RetryStrategy defines the interface for retry strategies
type RetryStrategy interface {
	ShouldRetry(attempt int, err error) bool
	GetDelay(attempt int) time.Duration
	GetMaxAttempts() int
}

// ExponentialBackoffStrategy implements exponential backoff retry
type ExponentialBackoffStrategy struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
	Multiplier  float64
	Jitter      bool
}

func NewExponentialBackoffStrategy(maxAttempts int, baseDelay, maxDelay time.Duration, multiplier float64, jitter bool) *ExponentialBackoffStrategy {
	return &ExponentialBackoffStrategy{
		MaxAttempts: maxAttempts,
		BaseDelay:   baseDelay,
		MaxDelay:    maxDelay,
		Multiplier:  multiplier,
		Jitter:      jitter,
	}
}

func (s *ExponentialBackoffStrategy) ShouldRetry(attempt int, err error) bool {
	return attempt < s.MaxAttempts && isRetryableError(err)
}

func (s *ExponentialBackoffStrategy) GetDelay(attempt int) time.Duration {
	delay := time.Duration(float64(s.BaseDelay) * math.Pow(s.Multiplier, float64(attempt)))
	
	if delay > s.MaxDelay {
		delay = s.MaxDelay
	}
	
	if s.Jitter {
		// Add jitter to prevent thundering herd
		jitterAmount := time.Duration(rand.Float64() * float64(delay) * 0.1)
		delay += jitterAmount
	}
	
	return delay
}

func (s *ExponentialBackoffStrategy) GetMaxAttempts() int {
	return s.MaxAttempts
}

// LinearBackoffStrategy implements linear backoff retry
type LinearBackoffStrategy struct {
	MaxAttempts int
	Delay       time.Duration
}

func NewLinearBackoffStrategy(maxAttempts int, delay time.Duration) *LinearBackoffStrategy {
	return &LinearBackoffStrategy{
		MaxAttempts: maxAttempts,
		Delay:       delay,
	}
}

func (s *LinearBackoffStrategy) ShouldRetry(attempt int, err error) bool {
	return attempt < s.MaxAttempts && isRetryableError(err)
}

func (s *LinearBackoffStrategy) GetDelay(attempt int) time.Duration {
	return s.Delay
}

func (s *LinearBackoffStrategy) GetMaxAttempts() int {
	return s.MaxAttempts
}

// FixedDelayStrategy implements fixed delay retry
type FixedDelayStrategy struct {
	MaxAttempts int
	Delay       time.Duration
}

func NewFixedDelayStrategy(maxAttempts int, delay time.Duration) *FixedDelayStrategy {
	return &FixedDelayStrategy{
		MaxAttempts: maxAttempts,
		Delay:       delay,
	}
}

func (s *FixedDelayStrategy) ShouldRetry(attempt int, err error) bool {
	return attempt < s.MaxAttempts && isRetryableError(err)
}

func (s *FixedDelayStrategy) GetDelay(attempt int) time.Duration {
	return s.Delay
}

func (s *FixedDelayStrategy) GetMaxAttempts() int {
	return s.MaxAttempts
}

// ============================================================================
// RETRY EXECUTOR
// ============================================================================

// RetryExecutor executes operations with retry strategies
type RetryExecutor struct {
	strategy RetryStrategy
	logger   *log.Logger
}

func NewRetryExecutor(strategy RetryStrategy, logger *log.Logger) *RetryExecutor {
	return &RetryExecutor{
		strategy: strategy,
		logger:   logger,
	}
}

func (re *RetryExecutor) Execute(operation func() error) error {
	var lastErr error
	
	for attempt := 0; attempt < re.strategy.GetMaxAttempts(); attempt++ {
		err := operation()
		if err == nil {
			if attempt > 0 {
				re.logger.Printf("Operation succeeded after %d retries", attempt)
			}
			return nil
		}
		
		lastErr = err
		
		if !re.strategy.ShouldRetry(attempt, err) {
			break
		}
		
		if attempt < re.strategy.GetMaxAttempts()-1 {
			delay := re.strategy.GetDelay(attempt)
			re.logger.Printf("Operation failed (attempt %d/%d): %v, retrying in %v", 
				attempt+1, re.strategy.GetMaxAttempts(), err, delay)
			time.Sleep(delay)
		}
	}
	
	return fmt.Errorf("operation failed after %d attempts: %w", re.strategy.GetMaxAttempts(), lastErr)
}

func (re *RetryExecutor) ExecuteWithContext(ctx context.Context, operation func(context.Context) error) error {
	var lastErr error
	
	for attempt := 0; attempt < re.strategy.GetMaxAttempts(); attempt++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		
		err := operation(ctx)
		if err == nil {
			if attempt > 0 {
				re.logger.Printf("Operation succeeded after %d retries", attempt)
			}
			return nil
		}
		
		lastErr = err
		
		if !re.strategy.ShouldRetry(attempt, err) {
			break
		}
		
		if attempt < re.strategy.GetMaxAttempts()-1 {
			delay := re.strategy.GetDelay(attempt)
			re.logger.Printf("Operation failed (attempt %d/%d): %v, retrying in %v", 
				attempt+1, re.strategy.GetMaxAttempts(), err, delay)
			
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
	
	return fmt.Errorf("operation failed after %d attempts: %w", re.strategy.GetMaxAttempts(), lastErr)
}

// ============================================================================
// CIRCUIT BREAKER
// ============================================================================

// CircuitBreakerState represents the state of a circuit breaker
type CircuitBreakerState int

const (
	StateClosed CircuitBreakerState = iota
	StateOpen
	StateHalfOpen
)

func (s CircuitBreakerState) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateOpen:
		return "open"
	case StateHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	name         string
	maxFailures  int
	timeout      time.Duration
	state        CircuitBreakerState
	failures     int
	lastFailTime time.Time
	successes    int
	mu           sync.RWMutex
	logger       *log.Logger
}

func NewCircuitBreaker(name string, maxFailures int, timeout time.Duration, logger *log.Logger) *CircuitBreaker {
	return &CircuitBreaker{
		name:        name,
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       StateClosed,
		logger:      logger,
	}
}

func (cb *CircuitBreaker) Call(operation func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	// Check if circuit should be opened
	if cb.state == StateOpen {
		if time.Since(cb.lastFailTime) > cb.timeout {
			cb.state = StateHalfOpen
			cb.successes = 0
			cb.logger.Printf("Circuit breaker %s: Moving to half-open state", cb.name)
		} else {
			return fmt.Errorf("circuit breaker %s is open", cb.name)
		}
	}
	
	// Execute the operation
	err := operation()
	
	if err != nil {
		cb.failures++
		cb.lastFailTime = time.Now()
		
		if cb.failures >= cb.maxFailures {
			cb.state = StateOpen
			cb.logger.Printf("Circuit breaker %s: Opened due to %d failures", cb.name, cb.failures)
		}
		
		return err
	}
	
	// Success - reset failure count
	cb.failures = 0
	cb.successes++
	
	if cb.state == StateHalfOpen {
		if cb.successes >= 3 { // Require 3 successes to close
			cb.state = StateClosed
			cb.logger.Printf("Circuit breaker %s: Closed after %d successes", cb.name, cb.successes)
		}
	}
	
	return nil
}

func (cb *CircuitBreaker) GetState() CircuitBreakerState {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

func (cb *CircuitBreaker) GetStats() map[string]interface{} {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	
	return map[string]interface{}{
		"name":      cb.name,
		"state":     cb.state.String(),
		"failures":  cb.failures,
		"successes": cb.successes,
		"last_fail": cb.lastFailTime,
	}
}

// ============================================================================
// FALLBACK MECHANISMS
// ============================================================================

// FallbackProvider defines the interface for fallback mechanisms
type FallbackProvider interface {
	Execute() (interface{}, error)
	IsAvailable() bool
	GetPriority() int
}

// FallbackManager manages multiple fallback providers
type FallbackManager struct {
	providers []FallbackProvider
	logger    *log.Logger
}

func NewFallbackManager(logger *log.Logger) *FallbackManager {
	return &FallbackManager{
		providers: make([]FallbackProvider, 0),
		logger:    logger,
	}
}

func (fm *FallbackManager) AddProvider(provider FallbackProvider) {
	fm.providers = append(fm.providers, provider)
}

func (fm *FallbackManager) ExecuteWithFallback(primary func() (interface{}, error)) (interface{}, error) {
	// Try primary operation first
	result, err := primary()
	if err == nil {
		return result, nil
	}
	
	fm.logger.Printf("Primary operation failed: %v, trying fallbacks", err)
	
	// Try fallback providers in order of priority
	for _, provider := range fm.providers {
		if !provider.IsAvailable() {
			continue
		}
		
		fm.logger.Printf("Trying fallback provider with priority %d", provider.GetPriority())
		result, err := provider.Execute()
		if err == nil {
			fm.logger.Printf("Fallback provider succeeded")
			return result, nil
		}
		
		fm.logger.Printf("Fallback provider failed: %v", err)
	}
	
	return nil, fmt.Errorf("all operations failed, including fallbacks")
}

// ============================================================================
// HEALTH CHECKS
// ============================================================================

// HealthCheck defines the interface for health checks
type HealthCheck interface {
	Name() string
	Check() error
	GetTimeout() time.Duration
}

// HealthChecker manages multiple health checks
type HealthChecker struct {
	checks map[string]HealthCheck
	logger *log.Logger
}

func NewHealthChecker(logger *log.Logger) *HealthChecker {
	return &HealthChecker{
		checks: make(map[string]HealthCheck),
		logger: logger,
	}
}

func (hc *HealthChecker) AddCheck(check HealthCheck) {
	hc.checks[check.Name()] = check
}

func (hc *HealthChecker) CheckAll() map[string]error {
	results := make(map[string]error)
	
	for name, check := range hc.checks {
		ctx, cancel := context.WithTimeout(context.Background(), check.GetTimeout())
		defer cancel()
		
		done := make(chan error, 1)
		go func() {
			done <- check.Check()
		}()
		
		select {
		case err := <-done:
			results[name] = err
		case <-ctx.Done():
			results[name] = fmt.Errorf("health check timeout")
		}
	}
	
	return results
}

func (hc *HealthChecker) IsHealthy() bool {
	results := hc.CheckAll()
	for _, err := range results {
		if err != nil {
			return false
		}
	}
	return true
}

// ============================================================================
// LOAD SHEDDING
// ============================================================================

// LoadShedder implements load shedding to prevent system overload
type LoadShedder struct {
	maxConcurrency int
	currentLoad    int
	mu             sync.RWMutex
	logger         *log.Logger
}

func NewLoadShedder(maxConcurrency int, logger *log.Logger) *LoadShedder {
	return &LoadShedder{
		maxConcurrency: maxConcurrency,
		logger:         logger,
	}
}

func (ls *LoadShedder) Acquire() bool {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	
	if ls.currentLoad >= ls.maxConcurrency {
		ls.logger.Printf("Load shedding: rejecting request (current: %d, max: %d)", 
			ls.currentLoad, ls.maxConcurrency)
		return false
	}
	
	ls.currentLoad++
	return true
}

func (ls *LoadShedder) Release() {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	
	if ls.currentLoad > 0 {
		ls.currentLoad--
	}
}

func (ls *LoadShedder) GetCurrentLoad() int {
	ls.mu.RLock()
	defer ls.mu.RUnlock()
	return ls.currentLoad
}

// ============================================================================
// BULKHEAD ISOLATION
// ============================================================================

// BulkheadPool represents an isolated resource pool
type BulkheadPool struct {
	name        string
	maxSize     int
	currentSize int
	mu          sync.RWMutex
	logger      *log.Logger
}

func NewBulkheadPool(name string, maxSize int, logger *log.Logger) *BulkheadPool {
	return &BulkheadPool{
		name:    name,
		maxSize: maxSize,
		logger:  logger,
	}
}

func (bp *BulkheadPool) Acquire() bool {
	bp.mu.Lock()
	defer bp.mu.Unlock()
	
	if bp.currentSize >= bp.maxSize {
		bp.logger.Printf("Bulkhead pool %s: rejecting request (current: %d, max: %d)", 
			bp.name, bp.currentSize, bp.maxSize)
		return false
	}
	
	bp.currentSize++
	return true
}

func (bp *BulkheadPool) Release() {
	bp.mu.Lock()
	defer bp.mu.Unlock()
	
	if bp.currentSize > 0 {
		bp.currentSize--
	}
}

func (bp *BulkheadPool) GetCurrentSize() int {
	bp.mu.RLock()
	defer bp.mu.RUnlock()
	return bp.currentSize
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateRetryStrategies() {
	fmt.Println("🔄 ERROR RECOVERY STRATEGIES MASTERY")
	fmt.Println("=====================================")
	fmt.Println()
	
	// Create logger
	logger := log.New(log.Writer(), "", log.LstdFlags)
	
	// Demonstrate exponential backoff
	fmt.Println("1. Exponential Backoff Strategy:")
	fmt.Println("--------------------------------")
	
	expStrategy := NewExponentialBackoffStrategy(5, 100*time.Millisecond, 5*time.Second, 2.0, true)
	expExecutor := NewRetryExecutor(expStrategy, logger)
	
	// Simulate failing operation
	attemptCount := 0
	err := expExecutor.Execute(func() error {
		attemptCount++
		if attemptCount < 3 {
			return errors.New("simulated failure")
		}
		return nil
	})
	
	if err != nil {
		fmt.Printf("   📊 Exponential backoff failed: %v\n", err)
	} else {
		fmt.Printf("   📊 Exponential backoff succeeded after %d attempts\n", attemptCount)
	}
	
	fmt.Println()
	
	// Demonstrate linear backoff
	fmt.Println("2. Linear Backoff Strategy:")
	fmt.Println("---------------------------")
	
	linearStrategy := NewLinearBackoffStrategy(3, 500*time.Millisecond)
	linearExecutor := NewRetryExecutor(linearStrategy, logger)
	
	attemptCount = 0
	err = linearExecutor.Execute(func() error {
		attemptCount++
		if attemptCount < 2 {
			return errors.New("simulated failure")
		}
		return nil
	})
	
	if err != nil {
		fmt.Printf("   📊 Linear backoff failed: %v\n", err)
	} else {
		fmt.Printf("   📊 Linear backoff succeeded after %d attempts\n", attemptCount)
	}
	
	fmt.Println()
}

func demonstrateCircuitBreaker() {
	fmt.Println("3. Circuit Breaker Pattern:")
	fmt.Println("---------------------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	cb := NewCircuitBreaker("test-service", 3, 10*time.Second, logger)
	
	// Simulate failures
	for i := 0; i < 5; i++ {
		err := cb.Call(func() error {
			if i < 3 {
				return errors.New("simulated failure")
			}
			return nil
		})
		
		state := cb.GetState()
		fmt.Printf("   📊 Attempt %d: State=%s, Error=%v\n", i+1, state, err)
	}
	
	// Show stats
	stats := cb.GetStats()
	fmt.Printf("   📊 Circuit Breaker Stats: %+v\n", stats)
	
	fmt.Println()
}

func demonstrateFallbackMechanisms() {
	fmt.Println("4. Fallback Mechanisms:")
	fmt.Println("-----------------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	fm := NewFallbackManager(logger)
	
	// Add fallback providers
	fm.AddProvider(&MockFallbackProvider{name: "cache", priority: 1, available: true})
	fm.AddProvider(&MockFallbackProvider{name: "database", priority: 2, available: true})
	fm.AddProvider(&MockFallbackProvider{name: "file", priority: 3, available: false})
	
	// Try primary operation (will fail)
	result, err := fm.ExecuteWithFallback(func() (interface{}, error) {
		return nil, errors.New("primary operation failed")
	})
	
	if err != nil {
		fmt.Printf("   📊 All operations failed: %v\n", err)
	} else {
		fmt.Printf("   📊 Fallback succeeded: %v\n", result)
	}
	
	fmt.Println()
}

func demonstrateHealthChecks() {
	fmt.Println("5. Health Checks:")
	fmt.Println("-----------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	hc := NewHealthChecker(logger)
	
	// Add health checks
	hc.AddCheck(&MockHealthCheck{name: "database", timeout: 5 * time.Second, healthy: true})
	hc.AddCheck(&MockHealthCheck{name: "redis", timeout: 3 * time.Second, healthy: false})
	hc.AddCheck(&MockHealthCheck{name: "api", timeout: 2 * time.Second, healthy: true})
	
	// Check all
	results := hc.CheckAll()
	for name, err := range results {
		if err != nil {
			fmt.Printf("   📊 %s: ❌ %v\n", name, err)
		} else {
			fmt.Printf("   📊 %s: ✅ healthy\n", name)
		}
	}
	
	fmt.Printf("   📊 Overall health: %t\n", hc.IsHealthy())
	
	fmt.Println()
}

func demonstrateLoadShedding() {
	fmt.Println("6. Load Shedding:")
	fmt.Println("-----------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	ls := NewLoadShedder(3, logger)
	
	// Simulate requests
	for i := 0; i < 5; i++ {
		if ls.Acquire() {
			fmt.Printf("   📊 Request %d: ✅ accepted (load: %d)\n", i+1, ls.GetCurrentLoad())
			// Simulate processing
			time.Sleep(100 * time.Millisecond)
			ls.Release()
		} else {
			fmt.Printf("   📊 Request %d: ❌ rejected (load: %d)\n", i+1, ls.GetCurrentLoad())
		}
	}
	
	fmt.Println()
}

func demonstrateBulkheadIsolation() {
	fmt.Println("7. Bulkhead Isolation:")
	fmt.Println("----------------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	
	// Create different bulkhead pools
	dbPool := NewBulkheadPool("database", 2, logger)
	apiPool := NewBulkheadPool("api", 3, logger)
	cachePool := NewBulkheadPool("cache", 1, logger)
	
	// Simulate requests to different pools
	pools := []*BulkheadPool{dbPool, apiPool, cachePool}
	
	for i := 0; i < 5; i++ {
		for j, pool := range pools {
			if pool.Acquire() {
				fmt.Printf("   📊 Pool %d: ✅ request %d accepted (size: %d)\n", j+1, i+1, pool.GetCurrentSize())
				// Simulate processing
				time.Sleep(50 * time.Millisecond)
				pool.Release()
			} else {
				fmt.Printf("   📊 Pool %d: ❌ request %d rejected (size: %d)\n", j+1, i+1, pool.GetCurrentSize())
			}
		}
	}
	
	fmt.Println()
}

// ============================================================================
// MOCK IMPLEMENTATIONS
// ============================================================================

// MockFallbackProvider implements FallbackProvider for testing
type MockFallbackProvider struct {
	name     string
	priority int
	available bool
}

func (m *MockFallbackProvider) Execute() (interface{}, error) {
	if !m.available {
		return nil, errors.New("provider not available")
	}
	return fmt.Sprintf("result from %s", m.name), nil
}

func (m *MockFallbackProvider) IsAvailable() bool {
	return m.available
}

func (m *MockFallbackProvider) GetPriority() int {
	return m.priority
}

// MockHealthCheck implements HealthCheck for testing
type MockHealthCheck struct {
	name    string
	timeout time.Duration
	healthy bool
}

func (m *MockHealthCheck) Name() string {
	return m.name
}

func (m *MockHealthCheck) Check() error {
	if !m.healthy {
		return errors.New("health check failed")
	}
	return nil
}

func (m *MockHealthCheck) GetTimeout() time.Duration {
	return m.timeout
}

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

func isRetryableError(err error) bool {
	// Simple retryable error detection
	if err == nil {
		return false
	}
	
	// Check for common retryable error patterns
	errorStr := err.Error()
	retryablePatterns := []string{
		"timeout",
		"connection",
		"network",
		"temporary",
		"busy",
		"unavailable",
	}
	
	for _, pattern := range retryablePatterns {
		if contains(errorStr, pattern) {
			return true
		}
	}
	
	return false
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}

// ============================================================================
// MAIN DEMONSTRATION
// ============================================================================

func main() {
	fmt.Println("🔄 ERROR RECOVERY STRATEGIES MASTERY")
	fmt.Println("=====================================")
	fmt.Println()
	
	// Demonstrate retry strategies
	demonstrateRetryStrategies()
	
	// Demonstrate circuit breaker
	demonstrateCircuitBreaker()
	
	// Demonstrate fallback mechanisms
	demonstrateFallbackMechanisms()
	
	// Demonstrate health checks
	demonstrateHealthChecks()
	
	// Demonstrate load shedding
	demonstrateLoadShedding()
	
	// Demonstrate bulkhead isolation
	demonstrateBulkheadIsolation()
	
	fmt.Println("🎉 ERROR RECOVERY STRATEGIES MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Retry strategies with exponential backoff")
	fmt.Println("✅ Circuit breaker pattern implementation")
	fmt.Println("✅ Fallback mechanisms and providers")
	fmt.Println("✅ Health checks and monitoring")
	fmt.Println("✅ Load shedding and resource management")
	fmt.Println("✅ Bulkhead isolation patterns")
	fmt.Println()
	fmt.Println("🚀 You are now ready for Panic Handling Mastery!")
}
