// 🔍 CONTEXTUAL LOGGING MASTERY
// Advanced contextual logging with request tracing and user context
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

// ============================================================================
// CONTEXT TYPES
// ============================================================================

type RequestContext struct {
	RequestID    string
	UserID       string
	SessionID    string
	IPAddress    string
	UserAgent    string
	TraceID      string
	SpanID       string
	ParentSpanID string
	StartTime    time.Time
	EndTime      time.Time
}

type UserContext struct {
	UserID       string
	Username     string
	Email        string
	Roles        []string
	Permissions  []string
	Organization string
	Department   string
}

type ServiceContext struct {
	ServiceName    string
	ServiceVersion string
	Environment    string
	Region         string
	InstanceID     string
	DeploymentID   string
}

type PerformanceContext struct {
	Duration      time.Duration
	MemoryUsage   int64
	CPUUsage      float64
	CacheHits     int64
	CacheMisses   int64
	DBQueries     int64
	HTTPRequests  int64
}

// ============================================================================
// CONTEXTUAL LOGGER
// ============================================================================

type ContextualLogger struct {
	baseLogger     *StructuredLogger
	context        map[string]interface{}
	mu             sync.RWMutex
	requestContext *RequestContext
	userContext    *UserContext
	serviceContext *ServiceContext
}

type StructuredLogger struct {
	level  LogLevel
	output io.Writer
}

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

type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Context   map[string]interface{} `json:"context,omitempty"`
}

func NewStructuredLogger(level LogLevel, output io.Writer) *StructuredLogger {
	return &StructuredLogger{
		level:  level,
		output: output,
	}
}

func (sl *StructuredLogger) log(level LogLevel, message string, fields map[string]interface{}) {
	if level.Priority() < sl.level.Priority() {
		return
	}
	
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Fields:    fields,
	}
	
	jsonData, _ := json.Marshal(entry)
	fmt.Fprintln(sl.output, string(jsonData))
}

func NewContextualLogger(baseLogger *StructuredLogger) *ContextualLogger {
	return &ContextualLogger{
		baseLogger: baseLogger,
		context:    make(map[string]interface{}),
	}
}

func (cl *ContextualLogger) WithRequestContext(reqCtx *RequestContext) *ContextualLogger {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	
	newLogger := &ContextualLogger{
		baseLogger:     cl.baseLogger,
		context:        make(map[string]interface{}),
		requestContext: reqCtx,
		userContext:    cl.userContext,
		serviceContext: cl.serviceContext,
	}
	
	// Copy existing context
	for k, v := range cl.context {
		newLogger.context[k] = v
	}
	
	// Add request context
	if reqCtx != nil {
		newLogger.context["request_id"] = reqCtx.RequestID
		newLogger.context["trace_id"] = reqCtx.TraceID
		newLogger.context["span_id"] = reqCtx.SpanID
		newLogger.context["user_id"] = reqCtx.UserID
		newLogger.context["session_id"] = reqCtx.SessionID
		newLogger.context["ip_address"] = reqCtx.IPAddress
		newLogger.context["user_agent"] = reqCtx.UserAgent
	}
	
	return newLogger
}

func (cl *ContextualLogger) WithUserContext(userCtx *UserContext) *ContextualLogger {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	
	newLogger := &ContextualLogger{
		baseLogger:     cl.baseLogger,
		context:        make(map[string]interface{}),
		requestContext: cl.requestContext,
		userContext:    userCtx,
		serviceContext: cl.serviceContext,
	}
	
	// Copy existing context
	for k, v := range cl.context {
		newLogger.context[k] = v
	}
	
	// Add user context
	if userCtx != nil {
		newLogger.context["user_id"] = userCtx.UserID
		newLogger.context["username"] = userCtx.Username
		newLogger.context["email"] = userCtx.Email
		newLogger.context["roles"] = userCtx.Roles
		newLogger.context["permissions"] = userCtx.Permissions
		newLogger.context["organization"] = userCtx.Organization
		newLogger.context["department"] = userCtx.Department
	}
	
	return newLogger
}

func (cl *ContextualLogger) WithServiceContext(serviceCtx *ServiceContext) *ContextualLogger {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	
	newLogger := &ContextualLogger{
		baseLogger:     cl.baseLogger,
		context:        make(map[string]interface{}),
		requestContext: cl.requestContext,
		userContext:    cl.userContext,
		serviceContext: serviceCtx,
	}
	
	// Copy existing context
	for k, v := range cl.context {
		newLogger.context[k] = v
	}
	
	// Add service context
	if serviceCtx != nil {
		newLogger.context["service_name"] = serviceCtx.ServiceName
		newLogger.context["service_version"] = serviceCtx.ServiceVersion
		newLogger.context["environment"] = serviceCtx.Environment
		newLogger.context["region"] = serviceCtx.Region
		newLogger.context["instance_id"] = serviceCtx.InstanceID
		newLogger.context["deployment_id"] = serviceCtx.DeploymentID
	}
	
	return newLogger
}

func (cl *ContextualLogger) WithFields(fields map[string]interface{}) *ContextualLogger {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	
	newLogger := &ContextualLogger{
		baseLogger:     cl.baseLogger,
		context:        make(map[string]interface{}),
		requestContext: cl.requestContext,
		userContext:    cl.userContext,
		serviceContext: cl.serviceContext,
	}
	
	// Copy existing context
	for k, v := range cl.context {
		newLogger.context[k] = v
	}
	
	// Add new fields
	for k, v := range fields {
		newLogger.context[k] = v
	}
	
	return newLogger
}

func (cl *ContextualLogger) log(level LogLevel, message string, fields map[string]interface{}) {
	// Merge context with fields
	allFields := make(map[string]interface{})
	for k, v := range cl.context {
		allFields[k] = v
	}
	for k, v := range fields {
		allFields[k] = v
	}
	
	cl.baseLogger.log(level, message, allFields)
}

// Logging methods
func (cl *ContextualLogger) Trace(message string, fields ...map[string]interface{}) {
	cl.log(TRACE, message, mergeFields(fields...))
}

func (cl *ContextualLogger) Debug(message string, fields ...map[string]interface{}) {
	cl.log(DEBUG, message, mergeFields(fields...))
}

func (cl *ContextualLogger) Info(message string, fields ...map[string]interface{}) {
	cl.log(INFO, message, mergeFields(fields...))
}

func (cl *ContextualLogger) Warn(message string, fields ...map[string]interface{}) {
	cl.log(WARN, message, mergeFields(fields...))
}

func (cl *ContextualLogger) Error(message string, fields ...map[string]interface{}) {
	cl.log(ERROR, message, mergeFields(fields...))
}

func (cl *ContextualLogger) Fatal(message string, fields ...map[string]interface{}) {
	cl.log(FATAL, message, mergeFields(fields...))
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
// CONTEXT MANAGER
// ============================================================================

type ContextManager struct {
	mu      sync.RWMutex
	contexts map[string]*RequestContext
}

func NewContextManager() *ContextManager {
	return &ContextManager{
		contexts: make(map[string]*RequestContext),
	}
}

func (cm *ContextManager) StartRequest(requestID, userID, sessionID, ipAddress, userAgent string) *RequestContext {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	ctx := &RequestContext{
		RequestID: requestID,
		UserID:    userID,
		SessionID: sessionID,
		IPAddress: ipAddress,
		UserAgent: userAgent,
		TraceID:   generateTraceID(),
		SpanID:    generateSpanID(),
		StartTime: time.Now(),
	}
	
	cm.contexts[requestID] = ctx
	return ctx
}

func (cm *ContextManager) EndRequest(requestID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	if ctx, exists := cm.contexts[requestID]; exists {
		ctx.EndTime = time.Now()
		delete(cm.contexts, requestID)
	}
}

func (cm *ContextManager) GetContext(requestID string) *RequestContext {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.contexts[requestID]
}

func generateTraceID() string {
	return fmt.Sprintf("trace-%d", time.Now().UnixNano())
}

func generateSpanID() string {
	return fmt.Sprintf("span-%d", time.Now().UnixNano())
}

// ============================================================================
// PERFORMANCE TRACKER
// ============================================================================

type PerformanceTracker struct {
	startTime    time.Time
	context      *PerformanceContext
	mu           sync.Mutex
}

func NewPerformanceTracker() *PerformanceTracker {
	return &PerformanceTracker{
		startTime: time.Now(),
		context:   &PerformanceContext{},
	}
}

func (pt *PerformanceTracker) Start() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.startTime = time.Now()
}

func (pt *PerformanceTracker) Stop() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.context.Duration = time.Since(pt.startTime)
}

func (pt *PerformanceTracker) AddCacheHit() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.context.CacheHits++
}

func (pt *PerformanceTracker) AddCacheMiss() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.context.CacheMisses++
}

func (pt *PerformanceTracker) AddDBQuery() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.context.DBQueries++
}

func (pt *PerformanceTracker) AddHTTPRequest() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.context.HTTPRequests++
}

func (pt *PerformanceTracker) GetContext() *PerformanceContext {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	return pt.context
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateRequestContext() {
	fmt.Println("\n=== Request Context ===")
	
	baseLogger := NewStructuredLogger(INFO, os.Stdout)
	logger := NewContextualLogger(baseLogger)
	
	// Create request context
	reqCtx := &RequestContext{
		RequestID: "req-123",
		UserID:    "user-456",
		SessionID: "sess-789",
		IPAddress: "192.168.1.100",
		UserAgent: "Mozilla/5.0...",
		TraceID:   "trace-abc",
		SpanID:    "span-def",
		StartTime: time.Now(),
	}
	
	// Create contextual logger with request context
	requestLogger := logger.WithRequestContext(reqCtx)
	
	requestLogger.Info("Processing user request", map[string]interface{}{
		"endpoint": "/api/users/profile",
		"method":   "GET",
	})
	
	requestLogger.Info("Database query executed", map[string]interface{}{
		"query":     "SELECT * FROM users WHERE id = ?",
		"duration_ms": 45,
	})
}

func demonstrateUserContext() {
	fmt.Println("\n=== User Context ===")
	
	baseLogger := NewStructuredLogger(INFO, os.Stdout)
	logger := NewContextualLogger(baseLogger)
	
	// Create user context
	userCtx := &UserContext{
		UserID:       "user-789",
		Username:     "john.doe",
		Email:        "john.doe@example.com",
		Roles:        []string{"admin", "user"},
		Permissions:  []string{"read", "write", "delete"},
		Organization: "Acme Corp",
		Department:   "Engineering",
	}
	
	// Create contextual logger with user context
	userLogger := logger.WithUserContext(userCtx)
	
	userLogger.Info("User performed action", map[string]interface{}{
		"action": "file_upload",
		"file_name": "document.pdf",
		"file_size": 1024000,
	})
	
	userLogger.Warn("Permission denied", map[string]interface{}{
		"required_permission": "admin",
		"user_permissions":    userCtx.Permissions,
	})
}

func demonstrateServiceContext() {
	fmt.Println("\n=== Service Context ===")
	
	baseLogger := NewStructuredLogger(INFO, os.Stdout)
	logger := NewContextualLogger(baseLogger)
	
	// Create service context
	serviceCtx := &ServiceContext{
		ServiceName:    "user-service",
		ServiceVersion: "2.1.0",
		Environment:    "production",
		Region:         "us-west-2",
		InstanceID:     "i-1234567890abcdef0",
		DeploymentID:   "deploy-20231201-001",
	}
	
	// Create contextual logger with service context
	serviceLogger := logger.WithServiceContext(serviceCtx)
	
	serviceLogger.Info("Service started", map[string]interface{}{
		"port": 8080,
		"pid":  12345,
	})
	
	serviceLogger.Info("Health check passed", map[string]interface{}{
		"status": "healthy",
		"uptime": "2h 30m",
	})
}

func demonstrateContextInheritance() {
	fmt.Println("\n=== Context Inheritance ===")
	
	baseLogger := NewStructuredLogger(INFO, os.Stdout)
	logger := NewContextualLogger(baseLogger)
	
	// Add base context
	baseLogger := logger.WithFields(map[string]interface{}{
		"component": "api-gateway",
		"version":   "1.0.0",
	})
	
	// Add request context
	reqCtx := &RequestContext{
		RequestID: "req-456",
		UserID:    "user-789",
		TraceID:   "trace-xyz",
		SpanID:    "span-abc",
	}
	requestLogger := baseLogger.WithRequestContext(reqCtx)
	
	// Add additional fields
	finalLogger := requestLogger.WithFields(map[string]interface{}{
		"operation": "user_authentication",
		"duration_ms": 150,
	})
	
	finalLogger.Info("Authentication completed", map[string]interface{}{
		"success": true,
		"method":  "oauth2",
	})
}

func demonstratePerformanceContext() {
	fmt.Println("\n=== Performance Context ===")
	
	baseLogger := NewStructuredLogger(INFO, os.Stdout)
	logger := NewContextualLogger(baseLogger)
	
	// Create performance tracker
	tracker := NewPerformanceTracker()
	tracker.Start()
	
	// Simulate some operations
	tracker.AddCacheHit()
	tracker.AddCacheMiss()
	tracker.AddDBQuery()
	tracker.AddHTTPRequest()
	
	tracker.Stop()
	
	// Add performance context to logger
	perfLogger := logger.WithFields(map[string]interface{}{
		"performance": tracker.GetContext(),
	})
	
	perfLogger.Info("Request completed", map[string]interface{}{
		"total_duration": tracker.GetContext().Duration,
		"cache_hits":     tracker.GetContext().CacheHits,
		"cache_misses":   tracker.GetContext().CacheMisses,
		"db_queries":     tracker.GetContext().DBQueries,
	})
}

func demonstrateContextManager() {
	fmt.Println("\n=== Context Manager ===")
	
	baseLogger := NewStructuredLogger(INFO, os.Stdout)
	logger := NewContextualLogger(baseLogger)
	contextManager := NewContextManager()
	
	// Start a request
	reqCtx := contextManager.StartRequest("req-999", "user-123", "sess-456", "192.168.1.200", "Chrome/91.0")
	requestLogger := logger.WithRequestContext(reqCtx)
	
	requestLogger.Info("Request started")
	
	// Simulate request processing
	time.Sleep(10 * time.Millisecond)
	
	requestLogger.Info("Request processing completed")
	
	// End the request
	contextManager.EndRequest("req-999")
	
	fmt.Println("   📊 Request context managed and cleaned up")
}

func main() {
	fmt.Println("🔍 CONTEXTUAL LOGGING MASTERY")
	fmt.Println("=============================")
	
	demonstrateRequestContext()
	demonstrateUserContext()
	demonstrateServiceContext()
	demonstrateContextInheritance()
	demonstratePerformanceContext()
	demonstrateContextManager()
	
	fmt.Println("\n🎉 CONTEXTUAL LOGGING MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Request-scoped contextual logging")
	fmt.Println("✅ User and session context")
	fmt.Println("✅ Service identification patterns")
	fmt.Println("✅ Context inheritance and composition")
	fmt.Println("✅ Performance context tracking")
	fmt.Println("✅ Context lifecycle management")
	
	fmt.Println("\n🚀 You are now ready for Log Aggregation Mastery!")
}
