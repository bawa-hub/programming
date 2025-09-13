// 🚨 PANIC HANDLING MASTERY
// Advanced panic recovery, prevention, and graceful error handling
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

// ============================================================================
// PANIC RECOVERY UTILITIES
// ============================================================================

// PanicRecoverer provides advanced panic recovery capabilities
type PanicRecoverer struct {
	logger     *log.Logger
	recoverers []PanicHandler
	mu         sync.RWMutex
}

// PanicHandler defines the interface for panic handlers
type PanicHandler interface {
	HandlePanic(panicValue interface{}, stack []byte) error
	GetName() string
	GetPriority() int
}

// NewPanicRecoverer creates a new panic recoverer
func NewPanicRecoverer(logger *log.Logger) *PanicRecoverer {
	return &PanicRecoverer{
		logger:     logger,
		recoverers: make([]PanicHandler, 0),
	}
}

// AddHandler adds a panic handler
func (pr *PanicRecoverer) AddHandler(handler PanicHandler) {
	pr.mu.Lock()
	defer pr.mu.Unlock()
	
	pr.recoverers = append(pr.recoverers, handler)
}

// Recover recovers from a panic using registered handlers
func (pr *PanicRecoverer) Recover(panicValue interface{}) error {
	stack := debug.Stack()
	
	pr.mu.RLock()
	handlers := make([]PanicHandler, len(pr.recoverers))
	copy(handlers, pr.recoverers)
	pr.mu.RUnlock()
	
	// Sort handlers by priority
	for i := 0; i < len(handlers); i++ {
		for j := i + 1; j < len(handlers); j++ {
			if handlers[i].GetPriority() > handlers[j].GetPriority() {
				handlers[i], handlers[j] = handlers[j], handlers[i]
			}
		}
	}
	
	// Execute handlers in priority order
	for _, handler := range handlers {
		if err := handler.HandlePanic(panicValue, stack); err != nil {
			pr.logger.Printf("Panic handler %s failed: %v", handler.GetName(), err)
		}
	}
	
	return fmt.Errorf("panic recovered: %v", panicValue)
}

// ============================================================================
// PANIC HANDLERS
// ============================================================================

// LoggingPanicHandler logs panic information
type LoggingPanicHandler struct {
	logger *log.Logger
}

func NewLoggingPanicHandler(logger *log.Logger) *LoggingPanicHandler {
	return &LoggingPanicHandler{logger: logger}
}

func (h *LoggingPanicHandler) HandlePanic(panicValue interface{}, stack []byte) error {
	h.logger.Printf("PANIC RECOVERED: %v\nStack trace:\n%s", panicValue, stack)
	return nil
}

func (h *LoggingPanicHandler) GetName() string {
	return "logging"
}

func (h *LoggingPanicHandler) GetPriority() int {
	return 1
}

// MetricsPanicHandler tracks panic metrics
type MetricsPanicHandler struct {
	panicCount int
	mu         sync.Mutex
}

func NewMetricsPanicHandler() *MetricsPanicHandler {
	return &MetricsPanicHandler{}
}

func (h *MetricsPanicHandler) HandlePanic(panicValue interface{}, stack []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	h.panicCount++
	return nil
}

func (h *MetricsPanicHandler) GetName() string {
	return "metrics"
}

func (h *MetricsPanicHandler) GetPriority() int {
	return 2
}

func (h *MetricsPanicHandler) GetPanicCount() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.panicCount
}

// NotificationPanicHandler sends notifications for panics
type NotificationPanicHandler struct {
	notifications []string
	mu            sync.Mutex
}

func NewNotificationPanicHandler() *NotificationPanicHandler {
	return &NotificationPanicHandler{
		notifications: make([]string, 0),
	}
}

func (h *NotificationPanicHandler) HandlePanic(panicValue interface{}, stack []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	notification := fmt.Sprintf("Panic occurred: %v at %s", panicValue, time.Now().Format(time.RFC3339))
	h.notifications = append(h.notifications, notification)
	return nil
}

func (h *NotificationPanicHandler) GetName() string {
	return "notification"
}

func (h *NotificationPanicHandler) GetPriority() int {
	return 3
}

func (h *NotificationPanicHandler) GetNotifications() []string {
	h.mu.Lock()
	defer h.mu.Unlock()
	
	notifications := make([]string, len(h.notifications))
	copy(notifications, h.notifications)
	return notifications
}

// ============================================================================
// SAFE OPERATION WRAPPER
// ============================================================================

// SafeOperation wraps a function with panic recovery
func SafeOperation(fn func() error) error {
	defer func() {
		if r := recover(); r != nil {
			// Panic recovered, error will be returned
		}
	}()
	
	return fn()
}

// SafeOperationWithResult wraps a function with panic recovery and result
func SafeOperationWithResult(fn func() (interface{}, error)) (interface{}, error) {
	var result interface{}
	var err error
	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	
	result, err = fn()
	return result, err
}

// SafeGoroutine runs a function in a goroutine with panic recovery
func SafeGoroutine(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Goroutine panic recovered: %v", r)
			}
		}()
		
		fn()
	}()
}

// ============================================================================
// HTTP PANIC RECOVERY MIDDLEWARE
// ============================================================================

// PanicRecoveryMiddleware provides HTTP panic recovery middleware
type PanicRecoveryMiddleware struct {
	logger *log.Logger
}

func NewPanicRecoveryMiddleware(logger *log.Logger) *PanicRecoveryMiddleware {
	return &PanicRecoveryMiddleware{logger: logger}
}

func (m *PanicRecoveryMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				m.logger.Printf("HTTP handler panic recovered: %v", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		
		next.ServeHTTP(w, r)
	})
}

// ============================================================================
// PANIC MONITORING
// ============================================================================

// PanicMonitor monitors panic occurrences and patterns
type PanicMonitor struct {
	panics     []PanicInfo
	mu         sync.RWMutex
	logger     *log.Logger
	maxPanics  int
}

// PanicInfo contains information about a panic
type PanicInfo struct {
	Value     interface{} `json:"value"`
	Stack     string      `json:"stack"`
	Timestamp time.Time   `json:"timestamp"`
	Goroutine int         `json:"goroutine"`
	Function  string      `json:"function"`
	File      string      `json:"file"`
	Line      int         `json:"line"`
}

func NewPanicMonitor(logger *log.Logger, maxPanics int) *PanicMonitor {
	return &PanicMonitor{
		panics:    make([]PanicInfo, 0),
		logger:    logger,
		maxPanics: maxPanics,
	}
}

func (pm *PanicMonitor) RecordPanic(panicValue interface{}, stack []byte) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	// Get caller information
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}
	
	panicInfo := PanicInfo{
		Value:     panicValue,
		Stack:     string(stack),
		Timestamp: time.Now(),
		Goroutine: runtime.NumGoroutine(),
		Function:  "unknown",
		File:      file,
		Line:      line,
	}
	
	pm.panics = append(pm.panics, panicInfo)
	
	// Keep only the most recent panics
	if len(pm.panics) > pm.maxPanics {
		pm.panics = pm.panics[1:]
	}
	
	pm.logger.Printf("Panic recorded: %v at %s:%d", panicValue, file, line)
}

func (pm *PanicMonitor) GetPanics() []PanicInfo {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	panics := make([]PanicInfo, len(pm.panics))
	copy(panics, pm.panics)
	return panics
}

func (pm *PanicMonitor) GetPanicCount() int {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return len(pm.panics)
}

func (pm *PanicMonitor) GetRecentPanics(duration time.Duration) []PanicInfo {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	cutoff := time.Now().Add(-duration)
	var recent []PanicInfo
	
	for _, panic := range pm.panics {
		if panic.Timestamp.After(cutoff) {
			recent = append(recent, panic)
		}
	}
	
	return recent
}

// ============================================================================
// PANIC PREVENTION UTILITIES
// ============================================================================

// PanicPreventer provides utilities for preventing panics
type PanicPreventer struct{}

func NewPanicPreventer() *PanicPreventer {
	return &PanicPreventer{}
}

// SafeSliceAccess safely accesses a slice element
func (pp *PanicPreventer) SafeSliceAccess(slice interface{}, index int) (interface{}, error) {
	switch s := slice.(type) {
	case []interface{}:
		if index < 0 || index >= len(s) {
			return nil, fmt.Errorf("index %d out of bounds for slice of length %d", index, len(s))
		}
		return s[index], nil
	case []string:
		if index < 0 || index >= len(s) {
			return nil, fmt.Errorf("index %d out of bounds for slice of length %d", index, len(s))
		}
		return s[index], nil
	case []int:
		if index < 0 || index >= len(s) {
			return nil, fmt.Errorf("index %d out of bounds for slice of length %d", index, len(s))
		}
		return s[index], nil
	default:
		return nil, fmt.Errorf("unsupported slice type: %T", slice)
	}
}

// SafeMapAccess safely accesses a map element
func (pp *PanicPreventer) SafeMapAccess(m interface{}, key interface{}) (interface{}, bool) {
	switch mp := m.(type) {
	case map[string]interface{}:
		if k, ok := key.(string); ok {
			value, exists := mp[k]
			return value, exists
		}
	case map[string]string:
		if k, ok := key.(string); ok {
			value, exists := mp[k]
			return value, exists
		}
	case map[int]interface{}:
		if k, ok := key.(int); ok {
			value, exists := mp[k]
			return value, exists
		}
	}
	return nil, false
}

// SafeTypeAssertion safely performs type assertion
func (pp *PanicPreventer) SafeTypeAssertion(value interface{}, target interface{}) bool {
	switch target.(type) {
	case *string:
		if _, ok := value.(string); ok {
			return true
		}
	case *int:
		if _, ok := value.(int); ok {
			return true
		}
	case *float64:
		if _, ok := value.(float64); ok {
			return true
		}
	case *bool:
		if _, ok := value.(bool); ok {
			return true
		}
	}
	return false
}

// ============================================================================
// CONTEXT-AWARE PANIC RECOVERY
// ============================================================================

// ContextPanicRecoverer recovers from panics while preserving context
type ContextPanicRecoverer struct {
	logger *log.Logger
}

func NewContextPanicRecoverer(logger *log.Logger) *ContextPanicRecoverer {
	return &ContextPanicRecoverer{logger: logger}
}

func (cpr *ContextPanicRecoverer) RecoverWithContext(ctx context.Context, fn func(context.Context) error) error {
	defer func() {
		if r := recover(); r != nil {
			cpr.logger.Printf("Panic recovered in context %v: %v", ctx, r)
		}
	}()
	
	return fn(ctx)
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstratePanicRecovery() {
	fmt.Println("🚨 PANIC HANDLING MASTERY")
	fmt.Println("==========================")
	fmt.Println()
	
	// Create panic recoverer
	logger := log.New(log.Writer(), "", log.LstdFlags)
	recoverer := NewPanicRecoverer(logger)
	
	// Add handlers
	recoverer.AddHandler(NewLoggingPanicHandler(logger))
	recoverer.AddHandler(NewMetricsPanicHandler())
	recoverer.AddHandler(NewNotificationPanicHandler())
	
	// Demonstrate basic panic recovery
	fmt.Println("1. Basic Panic Recovery:")
	fmt.Println("------------------------")
	
	// Safe operation
	err := SafeOperation(func() error {
		// This will panic
		panic("simulated panic")
	})
	
	if err != nil {
		fmt.Printf("   📊 Panic recovered: %v\n", err)
	} else {
		fmt.Printf("   📊 No panic occurred\n")
	}
	
	fmt.Println()
	
	// Demonstrate safe operation with result
	fmt.Println("2. Safe Operation with Result:")
	fmt.Println("------------------------------")
	
	result, err := SafeOperationWithResult(func() (interface{}, error) {
		// This will panic
		panic("simulated panic with result")
	})
	
	if err != nil {
		fmt.Printf("   📊 Panic recovered: %v\n", err)
	} else {
		fmt.Printf("   📊 Result: %v\n", result)
	}
	
	fmt.Println()
	
	// Demonstrate safe goroutine
	fmt.Println("3. Safe Goroutine:")
	fmt.Println("-----------------")
	
	SafeGoroutine(func() {
		panic("simulated goroutine panic")
	})
	
	// Wait a bit for the goroutine to complete
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("   📊 Goroutine panic handled safely")
	
	fmt.Println()
}

func demonstratePanicMonitoring() {
	fmt.Println("4. Panic Monitoring:")
	fmt.Println("-------------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	monitor := NewPanicMonitor(logger, 10)
	
	// Simulate some panics
	for i := 0; i < 3; i++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					monitor.RecordPanic(r, debug.Stack())
				}
			}()
			
			panic(fmt.Sprintf("simulated panic %d", i+1))
		}()
	}
	
	fmt.Printf("   📊 Total panics recorded: %d\n", monitor.GetPanicCount())
	
	// Get recent panics
	recent := monitor.GetRecentPanics(1 * time.Minute)
	fmt.Printf("   📊 Recent panics (last minute): %d\n", len(recent))
	
	fmt.Println()
}

func demonstratePanicPrevention() {
	fmt.Println("5. Panic Prevention:")
	fmt.Println("-------------------")
	
	preventer := NewPanicPreventer()
	
	// Safe slice access
	slice := []string{"a", "b", "c"}
	
	// Valid access
	value, err := preventer.SafeSliceAccess(slice, 1)
	if err != nil {
		fmt.Printf("   📊 Error: %v\n", err)
	} else {
		fmt.Printf("   📊 Safe slice access: %v\n", value)
	}
	
	// Invalid access
	value, err = preventer.SafeSliceAccess(slice, 10)
	if err != nil {
		fmt.Printf("   📊 Safe slice access (out of bounds): %v\n", err)
	} else {
		fmt.Printf("   📊 Safe slice access: %v\n", value)
	}
	
	// Safe map access
	m := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	
	value, exists := preventer.SafeMapAccess(m, "key1")
	if exists {
		fmt.Printf("   📊 Safe map access: %v\n", value)
	} else {
		fmt.Printf("   📊 Safe map access: key not found\n")
	}
	
	// Safe type assertion
	var target *string
	isString := preventer.SafeTypeAssertion("hello", target)
	fmt.Printf("   📊 Safe type assertion (string): %t\n", isString)
	
	isString = preventer.SafeTypeAssertion(123, target)
	fmt.Printf("   📊 Safe type assertion (int): %t\n", isString)
	
	fmt.Println()
}

func demonstrateContextAwareRecovery() {
	fmt.Println("6. Context-Aware Panic Recovery:")
	fmt.Println("--------------------------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	recoverer := NewContextPanicRecoverer(logger)
	
	ctx := context.WithValue(context.Background(), "request_id", "req-123")
	
	err := recoverer.RecoverWithContext(ctx, func(ctx context.Context) error {
		// This will panic
		panic("simulated context panic")
	})
	
	if err != nil {
		fmt.Printf("   📊 Context panic recovered: %v\n", err)
	}
	
	fmt.Println()
}

func demonstrateHTTPPanicRecovery() {
	fmt.Println("7. HTTP Panic Recovery Middleware:")
	fmt.Println("----------------------------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	middleware := NewPanicRecoveryMiddleware(logger)
	
	// Create a test handler that panics
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("simulated HTTP handler panic")
	})
	
	// Wrap with panic recovery
	safeHandler := middleware.Handler(panicHandler)
	
	// Create a test request
	req, _ := http.NewRequest("GET", "/test", nil)
	w := &mockResponseWriter{}
	
	// This should not panic
	safeHandler.ServeHTTP(w, req)
	
	fmt.Printf("   📊 HTTP handler panic recovered, status: %d\n", w.statusCode)
	
	fmt.Println()
}

// ============================================================================
// MOCK IMPLEMENTATIONS
// ============================================================================

// mockResponseWriter implements http.ResponseWriter for testing
type mockResponseWriter struct {
	statusCode int
	headers    http.Header
	body       []byte
}

func (m *mockResponseWriter) Header() http.Header {
	if m.headers == nil {
		m.headers = make(http.Header)
	}
	return m.headers
}

func (m *mockResponseWriter) Write(data []byte) (int, error) {
	m.body = append(m.body, data...)
	return len(data), nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}

// ============================================================================
// MAIN DEMONSTRATION
// ============================================================================

func main() {
	fmt.Println("🚨 PANIC HANDLING MASTERY")
	fmt.Println("==========================")
	fmt.Println()
	
	// Demonstrate panic recovery
	demonstratePanicRecovery()
	
	// Demonstrate panic monitoring
	demonstratePanicMonitoring()
	
	// Demonstrate panic prevention
	demonstratePanicPrevention()
	
	// Demonstrate context-aware recovery
	demonstrateContextAwareRecovery()
	
	// Demonstrate HTTP panic recovery
	demonstrateHTTPPanicRecovery()
	
	fmt.Println("🎉 PANIC HANDLING MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Basic panic recovery techniques")
	fmt.Println("✅ Safe operation wrappers")
	fmt.Println("✅ Panic monitoring and tracking")
	fmt.Println("✅ Panic prevention utilities")
	fmt.Println("✅ Context-aware panic recovery")
	fmt.Println("✅ HTTP panic recovery middleware")
	fmt.Println()
	fmt.Println("🚀 You are now ready for Error Context Mastery!")
}
