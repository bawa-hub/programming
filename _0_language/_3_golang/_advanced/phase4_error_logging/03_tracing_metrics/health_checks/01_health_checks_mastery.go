// 🏥 HEALTH CHECKS MASTERY
// Comprehensive health checking patterns for production systems
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================================
// HEALTH STATUS TYPES
// ============================================================================

type HealthStatus string

const (
	StatusUp       HealthStatus = "UP"
	StatusDown     HealthStatus = "DOWN"
	StatusDegraded HealthStatus = "DEGRADED"
	StatusUnknown  HealthStatus = "UNKNOWN"
)

type HealthCheck struct {
	Name        string                 `json:"name"`
	Status      HealthStatus           `json:"status"`
	Message     string                 `json:"message,omitempty"`
	Duration    time.Duration          `json:"duration"`
	Details     map[string]interface{} `json:"details,omitempty"`
	LastChecked time.Time              `json:"last_checked"`
}

type HealthResponse struct {
	Status    HealthStatus `json:"status"`
	Timestamp time.Time    `json:"timestamp"`
	Checks    []HealthCheck `json:"checks,omitempty"`
	Version   string       `json:"version,omitempty"`
	Uptime    time.Duration `json:"uptime,omitempty"`
}

// ============================================================================
// BASIC HEALTH CHECKER
// ============================================================================

type BasicHealthChecker struct {
	startTime time.Time
	version   string
	mu        sync.RWMutex
	checks    map[string]HealthCheck
}

func NewBasicHealthChecker(version string) *BasicHealthChecker {
	return &BasicHealthChecker{
		startTime: time.Now(),
		version:   version,
		checks:    make(map[string]HealthCheck),
	}
}

func (bhc *BasicHealthChecker) AddCheck(name string, checkFunc func() HealthCheck) {
	bhc.mu.Lock()
	defer bhc.mu.Unlock()
	bhc.checks[name] = checkFunc()
}

func (bhc *BasicHealthChecker) GetHealth() HealthResponse {
	bhc.mu.RLock()
	defer bhc.mu.RUnlock()
	
	checks := make([]HealthCheck, 0, len(bhc.checks))
	overallStatus := StatusUp
	
	for _, check := range bhc.checks {
		checks = append(checks, check)
		if check.Status == StatusDown {
			overallStatus = StatusDown
		} else if check.Status == StatusDegraded && overallStatus != StatusDown {
			overallStatus = StatusDegraded
		}
	}
	
	return HealthResponse{
		Status:    overallStatus,
		Timestamp: time.Now(),
		Checks:    checks,
		Version:   bhc.version,
		Uptime:    time.Since(bhc.startTime),
	}
}

func (bhc *BasicHealthChecker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	health := bhc.GetHealth()
	
	w.Header().Set("Content-Type", "application/json")
	
	statusCode := http.StatusOK
	if health.Status == StatusDown {
		statusCode = http.StatusServiceUnavailable
	} else if health.Status == StatusDegraded {
		statusCode = http.StatusOK // Still OK but degraded
	}
	
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(health)
}

// ============================================================================
// DEPENDENCY HEALTH CHECKER
// ============================================================================

type DependencyChecker struct {
	name     string
	checkFunc func() HealthCheck
	interval time.Duration
	mu       sync.RWMutex
	lastCheck HealthCheck
	ticker   *time.Ticker
	stopCh   chan struct{}
}

func NewDependencyChecker(name string, checkFunc func() HealthCheck, interval time.Duration) *DependencyChecker {
	dc := &DependencyChecker{
		name:      name,
		checkFunc: checkFunc,
		interval:  interval,
		stopCh:    make(chan struct{}),
	}
	
	// Initial check
	dc.lastCheck = checkFunc()
	
	// Start background checking
	dc.ticker = time.NewTicker(interval)
	go dc.run()
	
	return dc
}

func (dc *DependencyChecker) run() {
	for {
		select {
		case <-dc.ticker.C:
			check := dc.checkFunc()
			dc.mu.Lock()
			dc.lastCheck = check
			dc.mu.Unlock()
		case <-dc.stopCh:
			dc.ticker.Stop()
			return
		}
	}
}

func (dc *DependencyChecker) GetHealth() HealthCheck {
	dc.mu.RLock()
	defer dc.mu.RUnlock()
	return dc.lastCheck
}

func (dc *DependencyChecker) Stop() {
	close(dc.stopCh)
}

// ============================================================================
// CIRCUIT BREAKER
// ============================================================================

type CircuitState int

const (
	CircuitClosed CircuitState = iota
	CircuitOpen
	CircuitHalfOpen
)

type CircuitBreaker struct {
	name           string
	maxFailures    int
	resetTimeout   time.Duration
	state          CircuitState
	failureCount   int32
	lastFailure    time.Time
	mu             sync.RWMutex
	onStateChange  func(string, CircuitState)
}

func NewCircuitBreaker(name string, maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		name:         name,
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
		state:        CircuitClosed,
	}
}

func (cb *CircuitBreaker) SetStateChangeCallback(callback func(string, CircuitState)) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.onStateChange = callback
}

func (cb *CircuitBreaker) Execute(operation func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	// Check if circuit should be reset
	if cb.state == CircuitOpen && time.Since(cb.lastFailure) > cb.resetTimeout {
		cb.state = CircuitHalfOpen
		cb.failureCount = 0
		if cb.onStateChange != nil {
			cb.onStateChange(cb.name, cb.state)
		}
	}
	
	// Reject if circuit is open
	if cb.state == CircuitOpen {
		return fmt.Errorf("circuit breaker %s is open", cb.name)
	}
	
	// Execute operation
	err := operation()
	
	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()
		
		// Open circuit if max failures reached
		if cb.failureCount >= int32(cb.maxFailures) {
			oldState := cb.state
			cb.state = CircuitOpen
			if cb.onStateChange != nil && oldState != CircuitOpen {
				cb.onStateChange(cb.name, cb.state)
			}
		}
		return err
	}
	
	// Reset on success
	if cb.state == CircuitHalfOpen {
		cb.state = CircuitClosed
		cb.failureCount = 0
		if cb.onStateChange != nil {
			cb.onStateChange(cb.name, cb.state)
		}
	}
	
	return nil
}

func (cb *CircuitBreaker) GetState() CircuitState {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

func (cb *CircuitBreaker) GetHealth() HealthCheck {
	state := cb.GetState()
	
	var status HealthStatus
	var message string
	
	switch state {
	case CircuitClosed:
		status = StatusUp
		message = "Circuit is closed - normal operation"
	case CircuitOpen:
		status = StatusDown
		message = "Circuit is open - failing fast"
	case CircuitHalfOpen:
		status = StatusDegraded
		message = "Circuit is half-open - testing recovery"
	}
	
	return HealthCheck{
		Name:        cb.name,
		Status:      status,
		Message:     message,
		Duration:    0,
		Details: map[string]interface{}{
			"state":         state,
			"failure_count": atomic.LoadInt32(&cb.failureCount),
			"last_failure":  cb.lastFailure,
		},
		LastChecked: time.Now(),
	}
}

// ============================================================================
// KUBERNETES PROBES
// ============================================================================

type KubernetesProbes struct {
	readinessChecks map[string]func() bool
	livenessChecks  map[string]func() bool
	mu              sync.RWMutex
}

func NewKubernetesProbes() *KubernetesProbes {
	return &KubernetesProbes{
		readinessChecks: make(map[string]func() bool),
		livenessChecks:  make(map[string]func() bool),
	}
}

func (kp *KubernetesProbes) AddReadinessCheck(name string, check func() bool) {
	kp.mu.Lock()
	defer kp.mu.Unlock()
	kp.readinessChecks[name] = check
}

func (kp *KubernetesProbes) AddLivenessCheck(name string, check func() bool) {
	kp.mu.Lock()
	defer kp.mu.Unlock()
	kp.livenessChecks[name] = check
}

func (kp *KubernetesProbes) ReadinessProbe(w http.ResponseWriter, r *http.Request) {
	kp.mu.RLock()
	defer kp.mu.RUnlock()
	
	allReady := true
	results := make(map[string]bool)
	
	for name, check := range kp.readinessChecks {
		result := check()
		results[name] = result
		if !result {
			allReady = false
		}
	}
	
	if allReady {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "Not Ready",
			"checks": results,
		})
	}
}

func (kp *KubernetesProbes) LivenessProbe(w http.ResponseWriter, r *http.Request) {
	kp.mu.RLock()
	defer kp.mu.RUnlock()
	
	allAlive := true
	results := make(map[string]bool)
	
	for name, check := range kp.livenessChecks {
		result := check()
		results[name] = result
		if !result {
			allAlive = false
		}
	}
	
	if allAlive {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "Not Alive",
			"checks": results,
		})
	}
}

// ============================================================================
// HEALTH CHECK AGGREGATOR
// ============================================================================

type HealthAggregator struct {
	checkers map[string]HealthChecker
	mu       sync.RWMutex
}

type HealthChecker interface {
	GetHealth() HealthCheck
}

func NewHealthAggregator() *HealthAggregator {
	return &HealthAggregator{
		checkers: make(map[string]HealthChecker),
	}
}

func (ha *HealthAggregator) AddChecker(name string, checker HealthChecker) {
	ha.mu.Lock()
	defer ha.mu.Unlock()
	ha.checkers[name] = checker
}

func (ha *HealthAggregator) GetOverallHealth() HealthResponse {
	ha.mu.RLock()
	defer ha.mu.RUnlock()
	
	checks := make([]HealthCheck, 0, len(ha.checkers))
	overallStatus := StatusUp
	
	for name, checker := range ha.checkers {
		check := checker.GetHealth()
		check.Name = name
		checks = append(checks, check)
		
		if check.Status == StatusDown {
			overallStatus = StatusDown
		} else if check.Status == StatusDegraded && overallStatus != StatusDown {
			overallStatus = StatusDegraded
		}
	}
	
	return HealthResponse{
		Status:    overallStatus,
		Timestamp: time.Now(),
		Checks:    checks,
	}
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateBasicHealthChecks() {
	fmt.Println("\n=== Basic Health Checks ===")
	
	checker := NewBasicHealthChecker("1.0.0")
	
	// Add some basic checks
	checker.AddCheck("database", func() HealthCheck {
		start := time.Now()
		// Simulate database check
		time.Sleep(10 * time.Millisecond)
		return HealthCheck{
			Status:   StatusUp,
			Message:  "Database connection successful",
			Duration: time.Since(start),
		}
	})
	
	checker.AddCheck("memory", func() HealthCheck {
		start := time.Now()
		// Simulate memory check
		time.Sleep(5 * time.Millisecond)
		return HealthCheck{
			Status:   StatusUp,
			Message:  "Memory usage normal",
			Duration: time.Since(start),
			Details: map[string]interface{}{
				"used_mb": 512,
				"total_mb": 1024,
			},
		}
	})
	
	health := checker.GetHealth()
	fmt.Printf("   📊 Overall Status: %s\n", health.Status)
	fmt.Printf("   📊 Uptime: %v\n", health.Uptime)
	fmt.Printf("   📊 Checks: %d\n", len(health.Checks))
}

func demonstrateDependencyHealthChecks() {
	fmt.Println("\n=== Dependency Health Checks ===")
	
	// Simulate external service check
	externalServiceCheck := func() HealthCheck {
		start := time.Now()
		// Simulate network call
		time.Sleep(20 * time.Millisecond)
		
		// Randomly fail 20% of the time
		if rand.Float64() < 0.2 {
			return HealthCheck{
				Status:   StatusDown,
				Message:  "External service unavailable",
				Duration: time.Since(start),
			}
		}
		
		return HealthCheck{
			Status:   StatusUp,
			Message:  "External service responding",
			Duration: time.Since(start),
		}
	}
	
	checker := NewDependencyChecker("external-api", externalServiceCheck, 5*time.Second)
	defer checker.Stop()
	
	// Check a few times
	for i := 0; i < 3; i++ {
		health := checker.GetHealth()
		fmt.Printf("   📊 %s: %s (%v)\n", health.Name, health.Status, health.Duration)
		time.Sleep(2 * time.Second)
	}
}

func demonstrateCircuitBreaker() {
	fmt.Println("\n=== Circuit Breaker ===")
	
	cb := NewCircuitBreaker("api-service", 3, 10*time.Second)
	
	// Set up state change callback
	cb.SetStateChangeCallback(func(name string, state CircuitState) {
		stateNames := map[CircuitState]string{
			CircuitClosed:   "CLOSED",
			CircuitOpen:     "OPEN",
			CircuitHalfOpen: "HALF-OPEN",
		}
		fmt.Printf("   🔄 Circuit %s changed to %s\n", name, stateNames[state])
	})
	
	// Simulate operations that fail
	for i := 0; i < 5; i++ {
		err := cb.Execute(func() error {
			// Simulate operation that fails
			if rand.Float64() < 0.8 {
				return fmt.Errorf("operation failed")
			}
			return nil
		})
		
		health := cb.GetHealth()
		fmt.Printf("   📊 Operation %d: %v, Circuit: %s\n", i+1, err, health.Status)
	}
}

func demonstrateKubernetesProbes() {
	fmt.Println("\n=== Kubernetes Probes ===")
	
	probes := NewKubernetesProbes()
	
	// Add readiness checks
	probes.AddReadinessCheck("database", func() bool {
		// Simulate database readiness
		return rand.Float64() > 0.3
	})
	
	probes.AddReadinessCheck("cache", func() bool {
		// Simulate cache readiness
		return rand.Float64() > 0.1
	})
	
	// Add liveness checks
	probes.AddLivenessCheck("main-process", func() bool {
		// Main process should always be alive
		return true
	})
	
	probes.AddLivenessCheck("memory", func() bool {
		// Simulate memory check
		return rand.Float64() > 0.05
	})
	
	// Test probes
	fmt.Println("   📊 Testing Readiness Probe:")
	probes.ReadinessProbe(&mockResponseWriter{}, &http.Request{})
	
	fmt.Println("   📊 Testing Liveness Probe:")
	probes.LivenessProbe(&mockResponseWriter{}, &http.Request{})
}

func demonstrateHealthAggregation() {
	fmt.Println("\n=== Health Aggregation ===")
	
	aggregator := NewHealthAggregator()
	
	// Add various checkers
	aggregator.AddChecker("database", &mockHealthChecker{StatusUp, "Database OK"})
	aggregator.AddChecker("cache", &mockHealthChecker{StatusUp, "Cache OK"})
	aggregator.AddChecker("external-api", &mockHealthChecker{StatusDegraded, "API slow"})
	
	health := aggregator.GetOverallHealth()
	fmt.Printf("   📊 Overall Status: %s\n", health.Status)
	fmt.Printf("   📊 Total Checks: %d\n", len(health.Checks))
	
	for _, check := range health.Checks {
		fmt.Printf("   📊 %s: %s - %s\n", check.Name, check.Status, check.Message)
	}
}

func demonstrateHealthMiddleware() {
	fmt.Println("\n=== Health Check Middleware ===")
	
	checker := NewBasicHealthChecker("1.0.0")
	checker.AddCheck("app", func() HealthCheck {
		return HealthCheck{
			Status:  StatusUp,
			Message: "Application running",
		}
	})
	
	// Create a simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	
	// Add health check middleware
	healthMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/health" {
				checker.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
	
	wrappedHandler := healthMiddleware(handler)
	
	// Test the middleware
	req, _ := http.NewRequest("GET", "/health", nil)
	w := &mockResponseWriter{}
	wrappedHandler.ServeHTTP(w, req)
	
	fmt.Println("   📊 Health endpoint accessible via middleware")
}

// ============================================================================
// MOCK IMPLEMENTATIONS
// ============================================================================

type mockHealthChecker struct {
	status  HealthStatus
	message string
}

func (mhc *mockHealthChecker) GetHealth() HealthCheck {
	return HealthCheck{
		Status:  mhc.status,
		Message: mhc.message,
	}
}

type mockResponseWriter struct{}

func (mrw *mockResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (mrw *mockResponseWriter) Write(data []byte) (int, error) {
	fmt.Printf("   📤 Response: %s\n", string(data))
	return len(data), nil
}

func (mrw *mockResponseWriter) WriteHeader(code int) {
	fmt.Printf("   📤 Status Code: %d\n", code)
}

func main() {
	fmt.Println("🏥 HEALTH CHECKS MASTERY")
	fmt.Println("========================")
	
	demonstrateBasicHealthChecks()
	demonstrateDependencyHealthChecks()
	demonstrateCircuitBreaker()
	demonstrateKubernetesProbes()
	demonstrateHealthAggregation()
	demonstrateHealthMiddleware()
	
	fmt.Println("\n🎉 HEALTH CHECKS MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Basic health checking patterns")
	fmt.Println("✅ Dependency health monitoring")
	fmt.Println("✅ Circuit breaker patterns")
	fmt.Println("✅ Kubernetes readiness/liveness probes")
	fmt.Println("✅ Health check aggregation")
	fmt.Println("✅ Health check middleware")
	
	fmt.Println("\n🚀 You are now ready for Monitoring Patterns Mastery!")
}
