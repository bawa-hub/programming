// 🔍 ERROR CONTEXT MASTERY
// Advanced error context, tracing, and correlation techniques
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

// ============================================================================
// CONTEXTUAL ERROR TYPES
// ============================================================================

// ContextualError represents an error with rich context
type ContextualError struct {
	Message   string                 `json:"message"`
	Context   map[string]interface{} `json:"context"`
	Err       error                  `json:"-"`
	Timestamp time.Time              `json:"timestamp"`
	TraceID   string                 `json:"trace_id,omitempty"`
	SpanID    string                 `json:"span_id,omitempty"`
}

func (e *ContextualError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *ContextualError) Unwrap() error {
	return e.Err
}

func (e *ContextualError) GetContext() map[string]interface{} {
	return e.Context
}

func (e *ContextualError) GetTraceID() string {
	return e.TraceID
}

func (e *ContextualError) GetSpanID() string {
	return e.SpanID
}

// ============================================================================
// ERROR CONTEXT BUILDER
// ============================================================================

// ErrorContextBuilder builds rich error context
type ErrorContextBuilder struct {
	context map[string]interface{}
}

func NewErrorContextBuilder() *ErrorContextBuilder {
	return &ErrorContextBuilder{
		context: make(map[string]interface{}),
	}
}

func (b *ErrorContextBuilder) Add(key string, value interface{}) *ErrorContextBuilder {
	b.context[key] = value
	return b
}

func (b *ErrorContextBuilder) AddRequest(req *RequestContext) *ErrorContextBuilder {
	b.context["request"] = req
	return b
}

func (b *ErrorContextBuilder) AddUser(user *UserContext) *ErrorContextBuilder {
	b.context["user"] = user
	return b
}

func (b *ErrorContextBuilder) AddService(service *ServiceContext) *ErrorContextBuilder {
	b.context["service"] = service
	return b
}

func (b *ErrorContextBuilder) AddTiming(timing *TimingContext) *ErrorContextBuilder {
	b.context["timing"] = timing
	return b
}

func (b *ErrorContextBuilder) AddEnvironment(env *EnvironmentContext) *ErrorContextBuilder {
	b.context["environment"] = env
	return b
}

func (b *ErrorContextBuilder) AddTrace(traceID, spanID string) *ErrorContextBuilder {
	b.context["trace_id"] = traceID
	b.context["span_id"] = spanID
	return b
}

func (b *ErrorContextBuilder) Build() map[string]interface{} {
	// Add default context
	if _, exists := b.context["timestamp"]; !exists {
		b.context["timestamp"] = time.Now()
	}
	
	if _, exists := b.context["goroutine"]; !exists {
		b.context["goroutine"] = runtime.NumGoroutine()
	}
	
	return b.context
}

func (b *ErrorContextBuilder) Clear() *ErrorContextBuilder {
	b.context = make(map[string]interface{})
	return b
}

// ============================================================================
// CONTEXT TYPES
// ============================================================================

// RequestContext represents HTTP request context
type RequestContext struct {
	ID        string            `json:"id"`
	Method    string            `json:"method"`
	URL       string            `json:"url"`
	Headers   map[string]string `json:"headers"`
	UserAgent string            `json:"user_agent"`
	IP        string            `json:"ip"`
	Timestamp time.Time         `json:"timestamp"`
}

// UserContext represents user context
type UserContext struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	TenantID string `json:"tenant_id"`
}

// ServiceContext represents service context
type ServiceContext struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Region  string `json:"region"`
	Pod     string `json:"pod"`
}

// TimingContext represents timing context
type TimingContext struct {
	StartTime    time.Time     `json:"start_time"`
	EndTime      time.Time     `json:"end_time"`
	Duration     time.Duration `json:"duration"`
	DatabaseTime time.Duration `json:"database_time,omitempty"`
	APITime      time.Duration `json:"api_time,omitempty"`
}

// EnvironmentContext represents environment context
type EnvironmentContext struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Version   string `json:"version"`
	BuildID   string `json:"build_id"`
	CommitSHA string `json:"commit_sha"`
}

// ============================================================================
// ERROR CONTEXT MANAGER
// ============================================================================

// ErrorContextManager manages error context across the application
type ErrorContextManager struct {
	contexts map[string]map[string]interface{}
	mu       sync.RWMutex
	logger   *log.Logger
}

func NewErrorContextManager(logger *log.Logger) *ErrorContextManager {
	return &ErrorContextManager{
		contexts: make(map[string]map[string]interface{}),
		logger:   logger,
	}
}

func (ecm *ErrorContextManager) SetContext(key string, context map[string]interface{}) {
	ecm.mu.Lock()
	defer ecm.mu.Unlock()
	
	ecm.contexts[key] = context
}

func (ecm *ErrorContextManager) GetContext(key string) map[string]interface{} {
	ecm.mu.RLock()
	defer ecm.mu.RUnlock()
	
	if context, exists := ecm.contexts[key]; exists {
		return context
	}
	return nil
}

func (ecm *ErrorContextManager) MergeContext(key string, additionalContext map[string]interface{}) {
	ecm.mu.Lock()
	defer ecm.mu.Unlock()
	
	if context, exists := ecm.contexts[key]; exists {
		for k, v := range additionalContext {
			context[k] = v
		}
	} else {
		ecm.contexts[key] = additionalContext
	}
}

func (ecm *ErrorContextManager) ClearContext(key string) {
	ecm.mu.Lock()
	defer ecm.mu.Unlock()
	
	delete(ecm.contexts, key)
}

func (ecm *ErrorContextManager) GetAllContexts() map[string]map[string]interface{} {
	ecm.mu.RLock()
	defer ecm.mu.RUnlock()
	
	contexts := make(map[string]map[string]interface{})
	for k, v := range ecm.contexts {
		contexts[k] = v
	}
	return contexts
}

// ============================================================================
// ERROR CONTEXT PROPAGATION
// ============================================================================

// ErrorContextPropagator propagates context through error chains
type ErrorContextPropagator struct {
	manager *ErrorContextManager
	logger  *log.Logger
}

func NewErrorContextPropagator(manager *ErrorContextManager, logger *log.Logger) *ErrorContextPropagator {
	return &ErrorContextPropagator{
		manager: manager,
		logger:  logger,
	}
}

func (ecp *ErrorContextPropagator) WrapError(err error, message string, contextKey string) error {
	if err == nil {
		return nil
	}
	
	context := ecp.manager.GetContext(contextKey)
	if context == nil {
		context = make(map[string]interface{})
	}
	
	return &ContextualError{
		Message:   message,
		Context:   context,
		Err:       err,
		Timestamp: time.Now(),
	}
}

func (ecp *ErrorContextPropagator) WrapErrorWithTrace(err error, message string, contextKey string, traceID, spanID string) error {
	if err == nil {
		return nil
	}
	
	context := ecp.manager.GetContext(contextKey)
	if context == nil {
		context = make(map[string]interface{})
	}
	
	context["trace_id"] = traceID
	context["span_id"] = spanID
	
	return &ContextualError{
		Message:   message,
		Context:   context,
		Err:       err,
		Timestamp: time.Now(),
		TraceID:   traceID,
		SpanID:    spanID,
	}
}

// ============================================================================
// ERROR CONTEXT SERIALIZATION
// ============================================================================

// ErrorContextSerializer serializes error context for logging
type ErrorContextSerializer struct{}

func NewErrorContextSerializer() *ErrorContextSerializer {
	return &ErrorContextSerializer{}
}

func (ecs *ErrorContextSerializer) ToJSON(err error) ([]byte, error) {
	if contextualErr, ok := err.(*ContextualError); ok {
		return json.Marshal(contextualErr)
	}
	
	// Fallback for non-contextual errors
	return json.Marshal(map[string]interface{}{
		"message": err.Error(),
		"type":    "non_contextual_error",
	})
}

func (ecs *ErrorContextSerializer) ToMap(err error) map[string]interface{} {
	if contextualErr, ok := err.(*ContextualError); ok {
		return map[string]interface{}{
			"message":   contextualErr.Message,
			"context":   contextualErr.Context,
			"timestamp": contextualErr.Timestamp,
			"trace_id":  contextualErr.TraceID,
			"span_id":   contextualErr.SpanID,
			"error":     contextualErr.Err.Error(),
		}
	}
	
	return map[string]interface{}{
		"message": err.Error(),
		"type":    "non_contextual_error",
	}
}

// ============================================================================
// ERROR CONTEXT FILTERING
// ============================================================================

// ErrorContextFilter filters error context for different audiences
type ErrorContextFilter struct{}

func NewErrorContextFilter() *ErrorContextFilter {
	return &ErrorContextFilter{}
}

func (ecf *ErrorContextFilter) FilterForLogging(context map[string]interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})
	
	for k, v := range context {
		// Include all context for logging
		filtered[k] = v
	}
	
	return filtered
}

func (ecf *ErrorContextFilter) FilterForAPI(context map[string]interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})
	
	// Only include safe context for API responses
	safeKeys := []string{"request_id", "trace_id", "span_id", "timestamp", "service", "environment"}
	
	for _, key := range safeKeys {
		if value, exists := context[key]; exists {
			filtered[key] = value
		}
	}
	
	return filtered
}

func (ecf *ErrorContextFilter) FilterForUser(context map[string]interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})
	
	// Only include user-friendly context
	userKeys := []string{"request_id", "timestamp", "service"}
	
	for _, key := range userKeys {
		if value, exists := context[key]; exists {
			filtered[key] = value
		}
	}
	
	return filtered
}

// ============================================================================
// ERROR CONTEXT CORRELATION
// ============================================================================

// ErrorContextCorrelator correlates errors across services
type ErrorContextCorrelator struct {
	correlations map[string][]string
	mu           sync.RWMutex
	logger       *log.Logger
}

func NewErrorContextCorrelator(logger *log.Logger) *ErrorContextCorrelator {
	return &ErrorContextCorrelator{
		correlations: make(map[string][]string),
		logger:       logger,
	}
}

func (ecc *ErrorContextCorrelator) CorrelateError(traceID string, errorID string) {
	ecc.mu.Lock()
	defer ecc.mu.Unlock()
	
	if _, exists := ecc.correlations[traceID]; !exists {
		ecc.correlations[traceID] = make([]string, 0)
	}
	
	ecc.correlations[traceID] = append(ecc.correlations[traceID], errorID)
}

func (ecc *ErrorContextCorrelator) GetCorrelatedErrors(traceID string) []string {
	ecc.mu.RLock()
	defer ecc.mu.RUnlock()
	
	if errors, exists := ecc.correlations[traceID]; exists {
		return errors
	}
	return nil
}

func (ecc *ErrorContextCorrelator) GetAllCorrelations() map[string][]string {
	ecc.mu.RLock()
	defer ecc.mu.RUnlock()
	
	correlations := make(map[string][]string)
	for k, v := range ecc.correlations {
		correlations[k] = v
	}
	return correlations
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateErrorContext() {
	fmt.Println("🔍 ERROR CONTEXT MASTERY")
	fmt.Println("========================")
	fmt.Println()
	
	// Create context builder
	builder := NewErrorContextBuilder()
	
	// Build rich context
	context := builder.
		AddRequest(&RequestContext{
			ID:        "req-123",
			Method:    "GET",
			URL:       "/api/v1/users/123",
			Headers:   map[string]string{"Authorization": "Bearer token123"},
			UserAgent: "Mozilla/5.0",
			IP:        "192.168.1.1",
			Timestamp: time.Now(),
		}).
		AddUser(&UserContext{
			ID:       "user-456",
			Email:    "user@example.com",
			Role:     "admin",
			TenantID: "tenant-789",
		}).
		AddService(&ServiceContext{
			Name:    "user-service",
			Version: "1.2.3",
			Region:  "us-west-2",
			Pod:     "user-service-abc123",
		}).
		AddTiming(&TimingContext{
			StartTime:    time.Now().Add(-100 * time.Millisecond),
			EndTime:      time.Now(),
			Duration:     100 * time.Millisecond,
			DatabaseTime: 50 * time.Millisecond,
			APITime:      30 * time.Millisecond,
		}).
		AddEnvironment(&EnvironmentContext{
			Name:      "production",
			Region:    "us-west-2",
			Version:   "1.2.3",
			BuildID:   "build-456",
			CommitSHA: "abc123def456",
		}).
		AddTrace("trace-789", "span-101").
		Build()
	
	fmt.Println("1. Error Context Building:")
	fmt.Println("-------------------------")
	fmt.Printf("   📊 Built context with %d fields\n", len(context))
	
	// Create contextual error
	baseErr := fmt.Errorf("database connection failed")
	contextualErr := &ContextualError{
		Message:   "failed to get user",
		Context:   context,
		Err:       baseErr,
		Timestamp: time.Now(),
		TraceID:   "trace-789",
		SpanID:    "span-101",
	}
	
	fmt.Printf("   📊 Contextual Error: %v\n", contextualErr)
	fmt.Printf("   📊 Trace ID: %s\n", contextualErr.GetTraceID())
	fmt.Printf("   📊 Span ID: %s\n", contextualErr.GetSpanID())
	
	fmt.Println()
	
	// Demonstrate context serialization
	fmt.Println("2. Error Context Serialization:")
	fmt.Println("-------------------------------")
	
	serializer := NewErrorContextSerializer()
	
	// Serialize to JSON
	jsonData, _ := serializer.ToJSON(contextualErr)
	fmt.Printf("   📊 JSON Length: %d bytes\n", len(jsonData))
	
	// Serialize to map
	errorMap := serializer.ToMap(contextualErr)
	fmt.Printf("   📊 Map Keys: %d\n", len(errorMap))
	
	fmt.Println()
	
	// Demonstrate context filtering
	fmt.Println("3. Error Context Filtering:")
	fmt.Println("---------------------------")
	
	filter := NewErrorContextFilter()
	
	// Filter for logging
	loggingContext := filter.FilterForLogging(context)
	fmt.Printf("   📊 Logging Context Keys: %d\n", len(loggingContext))
	
	// Filter for API
	apiContext := filter.FilterForAPI(context)
	fmt.Printf("   📊 API Context Keys: %d\n", len(apiContext))
	
	// Filter for user
	userContext := filter.FilterForUser(context)
	fmt.Printf("   📊 User Context Keys: %d\n", len(userContext))
	
	fmt.Println()
}

func demonstrateContextPropagation() {
	fmt.Println("4. Error Context Propagation:")
	fmt.Println("-----------------------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	manager := NewErrorContextManager(logger)
	propagator := NewErrorContextPropagator(manager, logger)
	
	// Set context
	contextKey := "request-123"
	context := map[string]interface{}{
		"request_id": "req-123",
		"user_id":    "user-456",
		"service":    "user-service",
	}
	manager.SetContext(contextKey, context)
	
	// Create error with context
	baseErr := fmt.Errorf("database query failed")
	wrappedErr := propagator.WrapError(baseErr, "failed to get user", contextKey)
	
	fmt.Printf("   📊 Wrapped Error: %v\n", wrappedErr)
	
	if contextualErr, ok := wrappedErr.(*ContextualError); ok {
		fmt.Printf("   📊 Context Keys: %d\n", len(contextualErr.GetContext()))
	}
	
	fmt.Println()
}

func demonstrateContextCorrelation() {
	fmt.Println("5. Error Context Correlation:")
	fmt.Println("-----------------------------")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	correlator := NewErrorContextCorrelator(logger)
	
	// Correlate errors
	traceID := "trace-123"
	correlator.CorrelateError(traceID, "error-1")
	correlator.CorrelateError(traceID, "error-2")
	correlator.CorrelateError(traceID, "error-3")
	
	// Get correlated errors
	correlatedErrors := correlator.GetCorrelatedErrors(traceID)
	fmt.Printf("   📊 Correlated Errors for %s: %v\n", traceID, correlatedErrors)
	
	// Show all correlations
	allCorrelations := correlator.GetAllCorrelations()
	fmt.Printf("   📊 Total Correlations: %d\n", len(allCorrelations))
	
	fmt.Println()
}

func demonstrateRealWorldScenarios() {
	fmt.Println("6. Real-World Scenarios:")
	fmt.Println("------------------------")
	
	// Scenario 1: Microservices error propagation
	fmt.Println("   📊 Scenario 1: Microservices Error Propagation")
	
	logger := log.New(log.Writer(), "", log.LstdFlags)
	manager := NewErrorContextManager(logger)
	propagator := NewErrorContextPropagator(manager, logger)
	
	// Simulate service chain: API Gateway -> User Service -> Database
	traceID := "trace-456"
	spanID := "span-789"
	
	// Set context at API Gateway
	apiContext := map[string]interface{}{
		"request_id": "req-456",
		"method":     "GET",
		"url":        "/api/v1/users/123",
		"user_id":    "user-789",
		"service":    "api-gateway",
	}
	manager.SetContext("api-gateway", apiContext)
	
	// Propagate to User Service
	userServiceContext := map[string]interface{}{
		"service":    "user-service",
		"operation":  "get_user",
		"user_id":    "user-789",
		"trace_id":   traceID,
		"span_id":    spanID,
	}
	manager.MergeContext("api-gateway", userServiceContext)
	
	// Create error with full context
	dbErr := fmt.Errorf("connection timeout")
	userErr := propagator.WrapErrorWithTrace(dbErr, "failed to get user from database", "api-gateway", traceID, spanID)
	
	fmt.Printf("      📊 User Service Error: %v\n", userErr)
	
	// Scenario 2: Error context for debugging
	fmt.Println("   📊 Scenario 2: Error Context for Debugging")
	
	// Add debugging context
	debugContext := map[string]interface{}{
		"query":      "SELECT * FROM users WHERE id = ?",
		"parameters": []interface{}{"user-789"},
		"connection": "db-pool-1",
		"timeout":    "30s",
	}
	manager.MergeContext("api-gateway", debugContext)
	
	// Create final error with all context
	finalErr := propagator.WrapErrorWithTrace(userErr, "API request failed", "api-gateway", traceID, spanID)
	
	fmt.Printf("      📊 Final Error: %v\n", finalErr)
	
	if contextualErr, ok := finalErr.(*ContextualError); ok {
		fmt.Printf("      📊 Total Context Keys: %d\n", len(contextualErr.GetContext()))
	}
	
	fmt.Println()
}

// ============================================================================
// MAIN DEMONSTRATION
// ============================================================================

func main() {
	fmt.Println("🔍 ERROR CONTEXT MASTERY")
	fmt.Println("========================")
	fmt.Println()
	
	// Demonstrate error context building
	demonstrateErrorContext()
	
	// Demonstrate context propagation
	demonstrateContextPropagation()
	
	// Demonstrate context correlation
	demonstrateContextCorrelation()
	
	// Demonstrate real-world scenarios
	demonstrateRealWorldScenarios()
	
	fmt.Println("🎉 ERROR CONTEXT MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Error context building and enrichment")
	fmt.Println("✅ Context propagation through error chains")
	fmt.Println("✅ Context serialization and filtering")
	fmt.Println("✅ Error correlation across services")
	fmt.Println("✅ Real-world context scenarios")
	fmt.Println()
	fmt.Println("🚀 You are now ready for Structured Logging Mastery!")
}
