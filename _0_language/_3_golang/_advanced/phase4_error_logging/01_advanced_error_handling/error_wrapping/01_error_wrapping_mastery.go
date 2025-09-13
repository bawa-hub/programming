// 🔗 ERROR WRAPPING & UNWRAPPING MASTERY
// Advanced error wrapping, unwrapping, and chain traversal techniques
package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// ============================================================================
// ERROR WRAPPING UTILITIES
// ============================================================================

// ErrorWrapper provides advanced error wrapping capabilities
type ErrorWrapper struct {
	context map[string]interface{}
}

func NewErrorWrapper() *ErrorWrapper {
	return &ErrorWrapper{
		context: make(map[string]interface{}),
	}
}

// Wrap wraps an error with additional context
func (ew *ErrorWrapper) Wrap(err error, message string, context ...map[string]interface{}) error {
	if err == nil {
		return nil
	}
	
	// Add context to the error
	errorContext := make(map[string]interface{})
	for k, v := range ew.context {
		errorContext[k] = v
	}
	
	// Add additional context
	for _, ctx := range context {
		for k, v := range ctx {
			errorContext[k] = v
		}
	}
	
	// Create wrapped error with context
	return &WrappedError{
		Message: message,
		Context: errorContext,
		Err:     err,
		Time:    time.Now(),
	}
}

// WrapWithCode wraps an error with a specific error code
func (ew *ErrorWrapper) WrapWithCode(err error, message, code string, context ...map[string]interface{}) error {
	if err == nil {
		return nil
	}
	
	errorContext := make(map[string]interface{})
	errorContext["code"] = code
	
	// Add existing context
	for k, v := range ew.context {
		errorContext[k] = v
	}
	
	// Add additional context
	for _, ctx := range context {
		for k, v := range ctx {
			errorContext[k] = v
		}
	}
	
	return &WrappedError{
		Message: message,
		Context: errorContext,
		Err:     err,
		Time:    time.Now(),
	}
}

// AddContext adds context to the wrapper
func (ew *ErrorWrapper) AddContext(key string, value interface{}) *ErrorWrapper {
	ew.context[key] = value
	return ew
}

// ClearContext clears all context
func (ew *ErrorWrapper) ClearContext() *ErrorWrapper {
	ew.context = make(map[string]interface{})
	return ew
}

// ============================================================================
// WRAPPED ERROR TYPE
// ============================================================================

// WrappedError represents an error with additional context
type WrappedError struct {
	Message string                 `json:"message"`
	Context map[string]interface{} `json:"context"`
	Err     error                  `json:"-"`
	Time    time.Time              `json:"time"`
}

func (e *WrappedError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *WrappedError) Unwrap() error {
	return e.Err
}

func (e *WrappedError) GetContext() map[string]interface{} {
	return e.Context
}

func (e *WrappedError) GetCode() string {
	if code, ok := e.Context["code"].(string); ok {
		return code
	}
	return ""
}

func (e *WrappedError) GetTime() time.Time {
	return e.Time
}

// ============================================================================
// ERROR CHAIN TRAVERSAL
// ============================================================================

// ErrorChainTraverser provides utilities for traversing error chains
type ErrorChainTraverser struct{}

func NewErrorChainTraverser() *ErrorChainTraverser {
	return &ErrorChainTraverser{}
}

// FindRootCause finds the root cause of an error chain
func (ect *ErrorChainTraverser) FindRootCause(err error) error {
	if err == nil {
		return nil
	}
	
	current := err
	for {
		unwrapped := errors.Unwrap(current)
		if unwrapped == nil {
			break
		}
		current = unwrapped
	}
	
	return current
}

// GetErrorChain returns the complete error chain
func (ect *ErrorChainTraverser) GetErrorChain(err error) []error {
	if err == nil {
		return nil
	}
	
	var chain []error
	current := err
	
	for current != nil {
		chain = append(chain, current)
		current = errors.Unwrap(current)
	}
	
	return chain
}

// FindErrorType finds the first error of a specific type in the chain
func (ect *ErrorChainTraverser) FindErrorType(err error, target interface{}) bool {
	return errors.As(err, target)
}

// FindErrorCode finds the first error with a specific code
func (ect *ErrorChainTraverser) FindErrorCode(err error, code string) bool {
	chain := ect.GetErrorChain(err)
	
	for _, e := range chain {
		if wrapped, ok := e.(*WrappedError); ok {
			if wrapped.GetCode() == code {
				return true
			}
		}
	}
	
	return false
}

// GetErrorContexts returns all context from the error chain
func (ect *ErrorChainTraverser) GetErrorContexts(err error) []map[string]interface{} {
	var contexts []map[string]interface{}
	chain := ect.GetErrorChain(err)
	
	for _, e := range chain {
		if wrapped, ok := e.(*WrappedError); ok {
			contexts = append(contexts, wrapped.GetContext())
		}
	}
	
	return contexts
}

// ============================================================================
// ERROR TRANSFORMATION
// ============================================================================

// ErrorTransformer provides utilities for transforming errors
type ErrorTransformer struct{}

func NewErrorTransformer() *ErrorTransformer {
	return &ErrorTransformer{}
}

// Transform transforms an error while preserving the chain
func (et *ErrorTransformer) Transform(err error, transformer func(error) error) error {
	if err == nil {
		return nil
	}
	
	// Transform the root cause
	rootCause := errors.Unwrap(err)
	if rootCause == nil {
		rootCause = err
	}
	
	transformed := transformer(rootCause)
	if transformed == nil {
		return err
	}
	
	// Rebuild the chain with transformed root cause
	return &WrappedError{
		Message: "transformed error",
		Context: map[string]interface{}{"transformed": true},
		Err:     transformed,
		Time:    time.Now(),
	}
}

// Mask masks sensitive information in error messages
func (et *ErrorTransformer) Mask(err error, sensitiveFields []string) error {
	if err == nil {
		return nil
	}
	
	chain := NewErrorChainTraverser().GetErrorChain(err)
	maskedChain := make([]error, len(chain))
	
	for i, e := range chain {
		if wrapped, ok := e.(*WrappedError); ok {
			maskedContext := make(map[string]interface{})
			for k, v := range wrapped.Context {
				if contains(sensitiveFields, k) {
					maskedContext[k] = "***MASKED***"
				} else {
					maskedContext[k] = v
				}
			}
			
			maskedChain[i] = &WrappedError{
				Message: wrapped.Message,
				Context: maskedContext,
				Err:     wrapped.Err,
				Time:    wrapped.Time,
			}
		} else {
			maskedChain[i] = e
		}
	}
	
	// Rebuild the chain
	if len(maskedChain) == 0 {
		return err
	}
	
	result := maskedChain[0]
	for i := 1; i < len(maskedChain); i++ {
		result = &WrappedError{
			Message: "masked error",
			Context: map[string]interface{}{"masked": true},
			Err:     result,
			Time:    time.Now(),
		}
	}
	
	return result
}

// ============================================================================
// ERROR AGGREGATION
// ============================================================================

// ErrorAggregator aggregates multiple errors into a single error
type ErrorAggregator struct {
	errors []error
}

func NewErrorAggregator() *ErrorAggregator {
	return &ErrorAggregator{
		errors: make([]error, 0),
	}
}

func (ea *ErrorAggregator) Add(err error) {
	if err != nil {
		ea.errors = append(ea.errors, err)
	}
}

func (ea *ErrorAggregator) HasErrors() bool {
	return len(ea.errors) > 0
}

func (ea *ErrorAggregator) GetErrors() []error {
	return ea.errors
}

func (ea *ErrorAggregator) ToError() error {
	if len(ea.errors) == 0 {
		return nil
	}
	
	if len(ea.errors) == 1 {
		return ea.errors[0]
	}
	
	return &AggregatedError{
		Errors: ea.errors,
		Time:   time.Now(),
	}
}

// AggregatedError represents multiple errors
type AggregatedError struct {
	Errors []error   `json:"errors"`
	Time   time.Time `json:"time"`
}

func (e *AggregatedError) Error() string {
	if len(e.Errors) == 0 {
		return "no errors"
	}
	
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}
	
	var messages []string
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}
	
	return fmt.Sprintf("multiple errors: %s", strings.Join(messages, "; "))
}

func (e *AggregatedError) Unwrap() []error {
	return e.Errors
}

// ============================================================================
// ERROR FILTERING
// ============================================================================

// ErrorFilter provides utilities for filtering errors
type ErrorFilter struct{}

func NewErrorFilter() *ErrorFilter {
	return &ErrorFilter{}
}

// FilterByType filters errors by type
func (ef *ErrorFilter) FilterByType(err error, target interface{}) []error {
	var filtered []error
	chain := NewErrorChainTraverser().GetErrorChain(err)
	
	for _, e := range chain {
		if errors.As(e, target) {
			filtered = append(filtered, e)
		}
	}
	
	return filtered
}

// FilterByCode filters errors by code
func (ef *ErrorFilter) FilterByCode(err error, code string) []error {
	var filtered []error
	chain := NewErrorChainTraverser().GetErrorChain(err)
	
	for _, e := range chain {
		if wrapped, ok := e.(*WrappedError); ok {
			if wrapped.GetCode() == code {
				filtered = append(filtered, e)
			}
		}
	}
	
	return filtered
}

// FilterBySeverity filters errors by severity
func (ef *ErrorFilter) FilterBySeverity(err error, severity string) []error {
	var filtered []error
	chain := NewErrorChainTraverser().GetErrorChain(err)
	
	for _, e := range chain {
		if wrapped, ok := e.(*WrappedError); ok {
			if sev, ok := wrapped.Context["severity"].(string); ok && sev == severity {
				filtered = append(filtered, e)
			}
		}
	}
	
	return filtered
}

// ============================================================================
// ERROR LOGGING INTEGRATION
// ============================================================================

// ErrorLogger provides structured logging for errors
type ErrorLogger struct {
	logger *log.Logger
}

func NewErrorLogger(logger *log.Logger) *ErrorLogger {
	return &ErrorLogger{logger: logger}
}

// LogError logs an error with full context
func (el *ErrorLogger) LogError(err error, level string, additionalContext ...map[string]interface{}) {
	if err == nil {
		return
	}
	
	chain := NewErrorChainTraverser().GetErrorChain(err)
	contexts := NewErrorChainTraverser().GetErrorContexts(err)
	
	// Merge all context
	mergedContext := make(map[string]interface{})
	for _, ctx := range contexts {
		for k, v := range ctx {
			mergedContext[k] = v
		}
	}
	
	// Add additional context
	for _, ctx := range additionalContext {
		for k, v := range ctx {
			mergedContext[k] = v
		}
	}
	
	// Log the error
	el.logger.Printf("[%s] Error: %v | Context: %+v | Chain Length: %d", 
		level, err, mergedContext, len(chain))
}

// LogErrorChain logs the complete error chain
func (el *ErrorLogger) LogErrorChain(err error, level string) {
	if err == nil {
		return
	}
	
	chain := NewErrorChainTraverser().GetErrorChain(err)
	
	el.logger.Printf("[%s] Error Chain (length: %d):", level, len(chain))
	for i, e := range chain {
		el.logger.Printf("  [%d] %v", i, e)
		if wrapped, ok := e.(*WrappedError); ok {
			el.logger.Printf("      Context: %+v", wrapped.GetContext())
		}
	}
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateErrorWrapping() {
	fmt.Println("🔗 ERROR WRAPPING & UNWRAPPING MASTERY")
	fmt.Println("=======================================")
	fmt.Println()
	
	// Create error wrapper
	wrapper := NewErrorWrapper()
	
	// Demonstrate basic error wrapping
	fmt.Println("1. Basic Error Wrapping:")
	fmt.Println("------------------------")
	
	// Create a base error
	baseErr := errors.New("database connection failed")
	
	// Wrap with context
	wrappedErr := wrapper.Wrap(baseErr, "failed to initialize user service", map[string]interface{}{
		"service": "user-service",
		"version": "1.0.0",
	})
	
	fmt.Printf("   📊 Base Error: %v\n", baseErr)
	fmt.Printf("   📊 Wrapped Error: %v\n", wrappedErr)
	
	// Wrap with code
	codedErr := wrapper.WrapWithCode(wrappedErr, "user service startup failed", "STARTUP_ERROR", map[string]interface{}{
		"retry_count": 3,
		"timeout":     "30s",
	})
	
	fmt.Printf("   📊 Coded Error: %v\n", codedErr)
	fmt.Printf("   📊 Error Code: %s\n", codedErr.(*WrappedError).GetCode())
	
	fmt.Println()
	
	// Demonstrate error chain traversal
	fmt.Println("2. Error Chain Traversal:")
	fmt.Println("-------------------------")
	
	traverser := NewErrorChainTraverser()
	
	// Get error chain
	chain := traverser.GetErrorChain(codedErr)
	fmt.Printf("   📊 Error Chain Length: %d\n", len(chain))
	
	for i, err := range chain {
		fmt.Printf("   📊 Chain[%d]: %v\n", i, err)
	}
	
	// Find root cause
	rootCause := traverser.FindRootCause(codedErr)
	fmt.Printf("   📊 Root Cause: %v\n", rootCause)
	
	// Find error code
	hasCode := traverser.FindErrorCode(codedErr, "STARTUP_ERROR")
	fmt.Printf("   📊 Has STARTUP_ERROR code: %t\n", hasCode)
	
	fmt.Println()
	
	// Demonstrate error transformation
	fmt.Println("3. Error Transformation:")
	fmt.Println("-----------------------")
	
	transformer := NewErrorTransformer()
	
	// Transform error
	transformedErr := transformer.Transform(codedErr, func(err error) error {
		return fmt.Errorf("transformed: %v", err)
	})
	
	fmt.Printf("   📊 Transformed Error: %v\n", transformedErr)
	
	// Mask sensitive information
	maskedErr := transformer.Mask(codedErr, []string{"service", "version"})
	fmt.Printf("   📊 Masked Error: %v\n", maskedErr)
	
	fmt.Println()
	
	// Demonstrate error aggregation
	fmt.Println("4. Error Aggregation:")
	fmt.Println("--------------------")
	
	aggregator := NewErrorAggregator()
	aggregator.Add(errors.New("error 1"))
	aggregator.Add(errors.New("error 2"))
	aggregator.Add(errors.New("error 3"))
	
	aggregatedErr := aggregator.ToError()
	fmt.Printf("   📊 Aggregated Error: %v\n", aggregatedErr)
	
	fmt.Println()
	
	// Demonstrate error filtering
	fmt.Println("5. Error Filtering:")
	fmt.Println("------------------")
	
	filter := NewErrorFilter()
	
	// Filter by code
	filteredByCode := filter.FilterByCode(codedErr, "STARTUP_ERROR")
	fmt.Printf("   📊 Filtered by STARTUP_ERROR: %d errors\n", len(filteredByCode))
	
	// Filter by type
	var wrappedError *WrappedError
	filteredByType := filter.FilterByType(codedErr, &wrappedError)
	fmt.Printf("   📊 Filtered by WrappedError type: %d errors\n", len(filteredByType))
	
	fmt.Println()
}

func demonstrateErrorLogging() {
	fmt.Println("📝 ERROR LOGGING INTEGRATION:")
	fmt.Println("=============================")
	
	// Create error logger
	logger := log.New(log.Writer(), "", log.LstdFlags)
	errorLogger := NewErrorLogger(logger)
	
	// Create a complex error chain
	wrapper := NewErrorWrapper()
	baseErr := errors.New("file not found")
	wrappedErr := wrapper.Wrap(baseErr, "failed to load configuration", map[string]interface{}{
		"file":    "config.yaml",
		"service": "api-gateway",
	})
	codedErr := wrapper.WrapWithCode(wrappedErr, "service initialization failed", "INIT_ERROR", map[string]interface{}{
		"retry_count": 5,
		"timeout":     "60s",
	})
	
	// Log error with context
	fmt.Println("   📊 Logging error with context:")
	errorLogger.LogError(codedErr, "ERROR", map[string]interface{}{
		"request_id": "req-123",
		"user_id":    "user-456",
	})
	
	// Log complete error chain
	fmt.Println("   📊 Logging complete error chain:")
	errorLogger.LogErrorChain(codedErr, "DEBUG")
	
	fmt.Println()
}

func demonstrateRealWorldScenarios() {
	fmt.Println("🌍 REAL-WORLD SCENARIOS:")
	fmt.Println("========================")
	
	// Scenario 1: Microservices error propagation
	fmt.Println("1. Microservices Error Propagation:")
	fmt.Println("-----------------------------------")
	
	wrapper := NewErrorWrapper()
	
	// Simulate service chain: API Gateway -> User Service -> Database
	dbErr := errors.New("connection timeout")
	userServiceErr := wrapper.Wrap(dbErr, "failed to get user from database", map[string]interface{}{
		"service": "user-service",
		"table":   "users",
		"query":   "SELECT * FROM users WHERE id = ?",
	})
	apiGatewayErr := wrapper.Wrap(userServiceErr, "failed to process user request", map[string]interface{}{
		"service":     "api-gateway",
		"endpoint":    "/api/v1/users/123",
		"method":      "GET",
		"request_id":  "req-789",
		"user_agent":  "Mozilla/5.0",
	})
	
	fmt.Printf("   📊 API Gateway Error: %v\n", apiGatewayErr)
	
	// Extract request ID from error chain
	traverser := NewErrorChainTraverser()
	contexts := traverser.GetErrorContexts(apiGatewayErr)
	for _, ctx := range contexts {
		if requestID, ok := ctx["request_id"].(string); ok {
			fmt.Printf("   📊 Request ID: %s\n", requestID)
			break
		}
	}
	
	fmt.Println()
	
	// Scenario 2: Error recovery based on error chain
	fmt.Println("2. Error Recovery Based on Error Chain:")
	fmt.Println("---------------------------------------")
	
	// Check if error is recoverable
	recoverable := false
	chain := traverser.GetErrorChain(apiGatewayErr)
	for _, err := range chain {
		if wrapped, ok := err.(*WrappedError); ok {
			if code := wrapped.GetCode(); code == "TIMEOUT" || code == "CONNECTION_ERROR" {
				recoverable = true
				break
			}
		}
	}
	
	fmt.Printf("   📊 Error is recoverable: %t\n", recoverable)
	
	if recoverable {
		fmt.Println("   📊 Recovery strategy: retry with exponential backoff")
	}
	
	fmt.Println()
}

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ============================================================================
// MAIN DEMONSTRATION
// ============================================================================

func main() {
	fmt.Println("🔗 ERROR WRAPPING & UNWRAPPING MASTERY")
	fmt.Println("=======================================")
	fmt.Println()
	
	// Demonstrate error wrapping
	demonstrateErrorWrapping()
	
	// Demonstrate error logging
	demonstrateErrorLogging()
	
	// Demonstrate real-world scenarios
	demonstrateRealWorldScenarios()
	
	fmt.Println("🎉 ERROR WRAPPING & UNWRAPPING MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Error wrapping with context")
	fmt.Println("✅ Error chain traversal and analysis")
	fmt.Println("✅ Error transformation and masking")
	fmt.Println("✅ Error aggregation and filtering")
	fmt.Println("✅ Error logging integration")
	fmt.Println("✅ Real-world error handling scenarios")
	fmt.Println()
	fmt.Println("🚀 You are now ready for Error Recovery Strategies!")
}
