// 📊 LOG LEVELS MASTERY
// Advanced structured logging with proper levels and performance optimization
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================================
// LOG LEVEL DEFINITIONS
// ============================================================================

type LogLevel int

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	switch l {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func (l LogLevel) Priority() int {
	return int(l)
}

// ============================================================================
// STRUCTURED LOGGER
// ============================================================================

type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Caller    string                 `json:"caller,omitempty"`
	TraceID   string                 `json:"trace_id,omitempty"`
	SpanID    string                 `json:"span_id,omitempty"`
}

type StructuredLogger struct {
	mu           sync.RWMutex
	level        LogLevel
	output       io.Writer
	fields       map[string]interface{}
	callerDepth  int
	performance  *PerformanceMetrics
}

type PerformanceMetrics struct {
	TotalLogs     int64
	FilteredLogs  int64
	AverageTime   time.Duration
	MaxTime       time.Duration
	MinTime       time.Duration
}

func NewStructuredLogger(level LogLevel, output io.Writer) *StructuredLogger {
	return &StructuredLogger{
		level:       level,
		output:      output,
		fields:      make(map[string]interface{}),
		callerDepth: 2,
		performance: &PerformanceMetrics{},
	}
}

func (l *StructuredLogger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *StructuredLogger) GetLevel() LogLevel {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.level
}

func (l *StructuredLogger) WithFields(fields map[string]interface{}) *StructuredLogger {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	newFields := make(map[string]interface{})
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}
	
	return &StructuredLogger{
		level:       l.level,
		output:      l.output,
		fields:      newFields,
		callerDepth: l.callerDepth,
		performance: l.performance,
	}
}

func (l *StructuredLogger) log(level LogLevel, message string, fields map[string]interface{}) {
	if level.Priority() < l.level.Priority() {
		atomic.AddInt64(&l.performance.FilteredLogs, 1)
		return
	}
	
	start := time.Now()
	
	// Get caller information
	_, file, line, ok := runtime.Caller(l.callerDepth)
	caller := ""
	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}
	
	// Merge fields
	allFields := make(map[string]interface{})
	for k, v := range l.fields {
		allFields[k] = v
	}
	for k, v := range fields {
		allFields[k] = v
	}
	
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Fields:    allFields,
		Caller:    caller,
	}
	
	// Add trace/span IDs if present
	if traceID, ok := allFields["trace_id"].(string); ok {
		entry.TraceID = traceID
	}
	if spanID, ok := allFields["span_id"].(string); ok {
		entry.SpanID = spanID
	}
	
	// Serialize to JSON
	jsonData, err := json.Marshal(entry)
	if err != nil {
		// Fallback to simple logging
		fmt.Fprintf(l.output, "[%s] %s: %s\n", level.String(), time.Now().Format(time.RFC3339), message)
		return
	}
	
	// Write to output
	fmt.Fprintln(l.output, string(jsonData))
	
	// Update performance metrics
	duration := time.Since(start)
	atomic.AddInt64(&l.performance.TotalLogs, 1)
	
	// Update timing statistics
	l.updateTiming(duration)
}

func (l *StructuredLogger) updateTiming(duration time.Duration) {
	// Simple moving average for average time
	currentAvg := atomic.LoadInt64((*int64)(&l.performance.AverageTime))
	newAvg := (currentAvg + int64(duration)) / 2
	atomic.StoreInt64((*int64)(&l.performance.AverageTime), newAvg)
	
	// Update max time
	for {
		currentMax := atomic.LoadInt64((*int64)(&l.performance.MaxTime))
		if int64(duration) <= currentMax {
			break
		}
		if atomic.CompareAndSwapInt64((*int64)(&l.performance.MaxTime), currentMax, int64(duration)) {
			break
		}
	}
	
	// Update min time
	for {
		currentMin := atomic.LoadInt64((*int64)(&l.performance.MinTime))
		if currentMin == 0 || int64(duration) >= currentMin {
			break
		}
		if atomic.CompareAndSwapInt64((*int64)(&l.performance.MinTime), currentMin, int64(duration)) {
			break
		}
	}
}

// Logging methods
func (l *StructuredLogger) Trace(message string, fields ...map[string]interface{}) {
	l.log(TRACE, message, mergeFields(fields...))
}

func (l *StructuredLogger) Debug(message string, fields ...map[string]interface{}) {
	l.log(DEBUG, message, mergeFields(fields...))
}

func (l *StructuredLogger) Info(message string, fields ...map[string]interface{}) {
	l.log(INFO, message, mergeFields(fields...))
}

func (l *StructuredLogger) Warn(message string, fields ...map[string]interface{}) {
	l.log(WARN, message, mergeFields(fields...))
}

func (l *StructuredLogger) Error(message string, fields ...map[string]interface{}) {
	l.log(ERROR, message, mergeFields(fields...))
}

func (l *StructuredLogger) Fatal(message string, fields ...map[string]interface{}) {
	l.log(FATAL, message, mergeFields(fields...))
	os.Exit(1)
}

func mergeFields(fields ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, fieldMap := range fields {
		for k, v := range fieldMap {
			result[k] = v
		}
	}
	return result
}

// ============================================================================
// LOG ROUTER
// ============================================================================

type LogRouter struct {
	routes map[LogLevel][]io.Writer
	mu     sync.RWMutex
}

func NewLogRouter() *LogRouter {
	return &LogRouter{
		routes: make(map[LogLevel][]io.Writer),
	}
}

func (r *LogRouter) AddRoute(level LogLevel, writer io.Writer) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.routes[level] = append(r.routes[level], writer)
}

func (r *LogRouter) Route(level LogLevel, data []byte) {
	r.mu.RLock()
	writers := r.routes[level]
	r.mu.RUnlock()
	
	for _, writer := range writers {
		writer.Write(data)
	}
}

// ============================================================================
// PERFORMANCE MONITORING
// ============================================================================

func (l *StructuredLogger) GetPerformanceMetrics() *PerformanceMetrics {
	return &PerformanceMetrics{
		TotalLogs:    atomic.LoadInt64(&l.performance.TotalLogs),
		FilteredLogs: atomic.LoadInt64(&l.performance.FilteredLogs),
		AverageTime:  time.Duration(atomic.LoadInt64((*int64)(&l.performance.AverageTime))),
		MaxTime:      time.Duration(atomic.LoadInt64((*int64)(&l.performance.MaxTime))),
		MinTime:      time.Duration(atomic.LoadInt64((*int64)(&l.performance.MinTime))),
	}
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateLogLevels() {
	fmt.Println("\n=== Log Level Hierarchy ===")
	
	logger := NewStructuredLogger(TRACE, os.Stdout)
	
	// Test all log levels
	logger.Trace("This is a trace message", map[string]interface{}{
		"component": "auth",
		"action":    "login_attempt",
	})
	
	logger.Debug("Debug information", map[string]interface{}{
		"user_id":   123,
		"session_id": "sess-456",
	})
	
	logger.Info("User logged in successfully", map[string]interface{}{
		"user_id":   123,
		"ip_address": "192.168.1.1",
	})
	
	logger.Warn("High memory usage detected", map[string]interface{}{
		"memory_usage": "85%",
		"threshold":    "80%",
	})
	
	logger.Error("Database connection failed", map[string]interface{}{
		"error":      "connection timeout",
		"retry_count": 3,
	})
}

func demonstrateDynamicLevels() {
	fmt.Println("\n=== Dynamic Log Level Changes ===")
	
	logger := NewStructuredLogger(INFO, os.Stdout)
	
	// Log at different levels
	logger.Debug("This debug message won't be shown")
	logger.Info("This info message will be shown")
	
	// Change level to DEBUG
	logger.SetLevel(DEBUG)
	logger.Debug("Now this debug message will be shown")
	
	// Change level to ERROR
	logger.SetLevel(ERROR)
	logger.Warn("This warning won't be shown")
	logger.Error("This error will be shown")
}

func demonstrateStructuredLogging() {
	fmt.Println("\n=== Structured Logging ===")
	
	logger := NewStructuredLogger(INFO, os.Stdout)
	
	// Add context fields
	contextLogger := logger.WithFields(map[string]interface{}{
		"service":    "user-service",
		"version":    "1.2.3",
		"trace_id":   "trace-789",
		"span_id":    "span-101",
	})
	
	contextLogger.Info("Processing user request", map[string]interface{}{
		"user_id":    456,
		"endpoint":   "/api/users",
		"method":     "GET",
		"duration_ms": 150,
	})
}

func demonstrateLogRouting() {
	fmt.Println("\n=== Log Routing ===")
	
	// Create different output files for different levels
	errorFile, _ := os.Create("errors.log")
	warnFile, _ := os.Create("warnings.log")
	infoFile, _ := os.Create("info.log")
	defer errorFile.Close()
	defer warnFile.Close()
	defer infoFile.Close()
	
	router := NewLogRouter()
	router.AddRoute(ERROR, errorFile)
	router.AddRoute(WARN, warnFile)
	router.AddRoute(INFO, infoFile)
	
	// Create logger that routes to different files
	logger := NewStructuredLogger(INFO, &routingWriter{router: router})
	
	logger.Info("This goes to info.log")
	logger.Warn("This goes to warnings.log")
	logger.Error("This goes to errors.log")
	
	fmt.Println("   📊 Logs routed to different files based on level")
}

type routingWriter struct {
	router *LogRouter
}

func (w *routingWriter) Write(data []byte) (int, error) {
	// Parse log level from JSON data
	var entry LogEntry
	if err := json.Unmarshal(data, &entry); err == nil {
		w.router.Route(entry.Level, data)
	}
	return len(data), nil
}

func demonstratePerformanceOptimization() {
	fmt.Println("\n=== Performance Optimization ===")
	
	logger := NewStructuredLogger(INFO, os.Stdout)
	
	// Simulate high-frequency logging
	start := time.Now()
	for i := 0; i < 1000; i++ {
		logger.Info("High frequency log message", map[string]interface{}{
			"iteration": i,
			"timestamp": time.Now().UnixNano(),
		})
	}
	duration := time.Since(start)
	
	metrics := logger.GetPerformanceMetrics()
	
	fmt.Printf("   📊 Total logs: %d\n", metrics.TotalLogs)
	fmt.Printf("   📊 Filtered logs: %d\n", metrics.FilteredLogs)
	fmt.Printf("   📊 Total time: %v\n", duration)
	fmt.Printf("   📊 Average time per log: %v\n", metrics.AverageTime)
	fmt.Printf("   📊 Max time per log: %v\n", metrics.MaxTime)
	fmt.Printf("   📊 Min time per log: %v\n", metrics.MinTime)
}

func demonstrateLazyEvaluation() {
	fmt.Println("\n=== Lazy Evaluation ===")
	
	logger := NewStructuredLogger(DEBUG, os.Stdout)
	
	// Expensive operation that should only run if DEBUG is enabled
	expensiveOperation := func() map[string]interface{} {
		fmt.Println("   🔍 Expensive operation executed!")
		return map[string]interface{}{
			"expensive_data": "This took a long time to compute",
			"computation_time": "5 seconds",
		}
	}
	
	// Only execute expensive operation if DEBUG level is enabled
	if logger.GetLevel() <= DEBUG {
		logger.Debug("Debug with expensive data", expensiveOperation())
	}
	
	// Change level to INFO - expensive operation won't run
	logger.SetLevel(INFO)
	if logger.GetLevel() <= DEBUG {
		logger.Debug("This won't execute expensive operation", expensiveOperation())
	}
}

func main() {
	fmt.Println("📊 LOG LEVELS MASTERY")
	fmt.Println("======================")
	
	demonstrateLogLevels()
	demonstrateDynamicLevels()
	demonstrateStructuredLogging()
	demonstrateLogRouting()
	demonstratePerformanceOptimization()
	demonstrateLazyEvaluation()
	
	fmt.Println("\n🎉 LOG LEVELS MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Complete log level hierarchy")
	fmt.Println("✅ Dynamic log level configuration")
	fmt.Println("✅ Structured JSON logging")
	fmt.Println("✅ Log routing and filtering")
	fmt.Println("✅ Performance optimization")
	fmt.Println("✅ Lazy evaluation patterns")
	
	fmt.Println("\n🚀 You are now ready for Contextual Logging Mastery!")
}
