// 🚨 CUSTOM ERROR TYPES MASTERY
// Advanced error handling with custom error types and hierarchies
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ============================================================================
// BASIC CUSTOM ERROR TYPES
// ============================================================================

// ValidationError represents a validation failure
type ValidationError struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s (value: %v, code: %s)", 
		e.Field, e.Message, e.Value, e.Code)
}

// BusinessError represents a business logic error
type BusinessError struct {
	Operation string    `json:"operation"`
	Resource  string    `json:"resource"`
	Message   string    `json:"message"`
	Code      string    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BusinessError) Error() string {
	return fmt.Sprintf("business error in %s operation on %s: %s (code: %s)", 
		e.Operation, e.Resource, e.Message, e.Code)
}

// SystemError represents a system-level error
type SystemError struct {
	Component string    `json:"component"`
	Operation string    `json:"operation"`
	Message   string    `json:"message"`
	Code      string    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Severity  Severity  `json:"severity"`
}

func (e SystemError) Error() string {
	return fmt.Sprintf("system error in %s.%s: %s (code: %s, severity: %s)", 
		e.Component, e.Operation, e.Message, e.Code, e.Severity)
}

// ============================================================================
// ERROR HIERARCHIES
// ============================================================================

// Severity represents error severity levels
type Severity string

const (
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

// DatabaseError represents a database-related error
type DatabaseError struct {
	Operation string    `json:"operation"`
	Table     string    `json:"table"`
	Query     string    `json:"query,omitempty"`
	Message   string    `json:"message"`
	Code      string    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Err       error     `json:"-"`
}

func (e DatabaseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("database error in %s operation on table %s: %s (code: %s) - %v", 
			e.Operation, e.Table, e.Message, e.Code, e.Err)
	}
	return fmt.Sprintf("database error in %s operation on table %s: %s (code: %s)", 
		e.Operation, e.Table, e.Message, e.Code)
}

func (e DatabaseError) Unwrap() error {
	return e.Err
}

// ConnectionError represents a database connection error
type ConnectionError struct {
	DatabaseError
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Timeout int    `json:"timeout"`
}

func (e ConnectionError) Error() string {
	return fmt.Sprintf("connection error to %s:%d (timeout: %ds) in %s operation: %s (code: %s)", 
		e.Host, e.Port, e.Timeout, e.Operation, e.Message, e.Code)
}

// QueryError represents a database query error
type QueryError struct {
	DatabaseError
	SQL         string `json:"sql"`
	Parameters  []interface{} `json:"parameters,omitempty"`
	RowCount    int    `json:"row_count,omitempty"`
}

func (e QueryError) Error() string {
	return fmt.Sprintf("query error in %s operation on table %s: %s (code: %s) - SQL: %s", 
		e.Operation, e.Table, e.Message, e.Code, e.SQL)
}

// TransactionError represents a database transaction error
type TransactionError struct {
	DatabaseError
	TransactionID string `json:"transaction_id"`
	IsolationLevel string `json:"isolation_level,omitempty"`
}

func (e TransactionError) Error() string {
	return fmt.Sprintf("transaction error (ID: %s) in %s operation: %s (code: %s)", 
		e.TransactionID, e.Operation, e.Message, e.Code)
}

// ============================================================================
// ERROR INTERFACES
// ============================================================================

// CodedError interface for errors with codes
type CodedError interface {
	error
	Code() string
	Severity() Severity
}

// RecoverableError interface for errors that can be recovered
type RecoverableError interface {
	error
	CanRecover() bool
	RecoveryStrategy() string
	RetryAfter() time.Duration
}

// TraceableError interface for errors with tracing information
type TraceableError interface {
	error
	TraceID() string
	SpanID() string
	Operation() string
}

// ============================================================================
// ADVANCED ERROR TYPES
// ============================================================================

// APIError represents an API-related error
type APIError struct {
	StatusCode    int               `json:"status_code"`
	ErrorCode     string            `json:"error_code"`
	Message       string            `json:"message"`
	Details       map[string]string `json:"details,omitempty"`
	RequestID     string            `json:"request_id,omitempty"`
	Timestamp     time.Time         `json:"timestamp"`
	CorrelationID string            `json:"correlation_id,omitempty"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error %d (%s): %s (request: %s)", 
		e.StatusCode, e.ErrorCode, e.Message, e.RequestID)
}

func (e APIError) Code() string {
	return e.ErrorCode
}

func (e APIError) Severity() Severity {
	switch e.StatusCode {
	case 400, 401, 403:
		return SeverityMedium
	case 404:
		return SeverityLow
	case 500, 502, 503, 504:
		return SeverityHigh
	default:
		return SeverityMedium
	}
}

// NetworkError represents a network-related error
type NetworkError struct {
	Operation   string        `json:"operation"`
	URL         string        `json:"url"`
	Method      string        `json:"method"`
	StatusCode  int           `json:"status_code,omitempty"`
	Message     string        `json:"message"`
	Code        string        `json:"code"`
	Timeout     time.Duration `json:"timeout"`
	RetryCount  int           `json:"retry_count"`
	Timestamp   time.Time     `json:"timestamp"`
	Err         error         `json:"-"`
}

func (e NetworkError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("network error in %s %s %s: %s (code: %s, timeout: %v) - %v", 
			e.Method, e.URL, e.Operation, e.Message, e.Code, e.Timeout, e.Err)
	}
	return fmt.Sprintf("network error in %s %s %s: %s (code: %s, timeout: %v)", 
		e.Method, e.URL, e.Operation, e.Message, e.Code, e.Timeout)
}

func (e NetworkError) Unwrap() error {
	return e.Err
}

func (e NetworkError) CanRecover() bool {
	return e.StatusCode >= 500 || e.StatusCode == 429
}

func (e NetworkError) RecoveryStrategy() string {
	if e.StatusCode >= 500 {
		return "retry with exponential backoff"
	}
	if e.StatusCode == 429 {
		return "retry after rate limit reset"
	}
	return "no recovery possible"
}

func (e NetworkError) RetryAfter() time.Duration {
	if e.StatusCode == 429 {
		return 60 * time.Second
	}
	return time.Duration(e.RetryCount) * time.Second
}

// ============================================================================
// ERROR FACTORIES
// ============================================================================

// ErrorFactory creates consistent errors
type ErrorFactory struct {
	ServiceName string
	Version     string
}

func NewErrorFactory(serviceName, version string) *ErrorFactory {
	return &ErrorFactory{
		ServiceName: serviceName,
		Version:     version,
	}
}

func (ef *ErrorFactory) NewValidationError(field, message, code string, value interface{}) ValidationError {
	return ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
		Code:    code,
	}
}

func (ef *ErrorFactory) NewBusinessError(operation, resource, message, code string) BusinessError {
	return BusinessError{
		Operation: operation,
		Resource:  resource,
		Message:   message,
		Code:      code,
		Timestamp: time.Now(),
	}
}

func (ef *ErrorFactory) NewSystemError(component, operation, message, code string, severity Severity) SystemError {
	return SystemError{
		Component: component,
		Operation: operation,
		Message:   message,
		Code:      code,
		Timestamp: time.Now(),
		Severity:  severity,
	}
}

func (ef *ErrorFactory) NewDatabaseError(operation, table, message, code string, err error) DatabaseError {
	return DatabaseError{
		Operation: operation,
		Table:     table,
		Message:   message,
		Code:      code,
		Timestamp: time.Now(),
		Err:       err,
	}
}

func (ef *ErrorFactory) NewAPIError(statusCode int, errorCode, message, requestID string) APIError {
	return APIError{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    message,
		RequestID:  requestID,
		Timestamp:  time.Now(),
	}
}

func (ef *ErrorFactory) NewNetworkError(operation, url, method, message, code string, statusCode int, timeout time.Duration, err error) NetworkError {
	return NetworkError{
		Operation:  operation,
		URL:        url,
		Method:     method,
		StatusCode: statusCode,
		Message:    message,
		Code:       code,
		Timeout:    timeout,
		Timestamp:  time.Now(),
		Err:        err,
	}
}

// ============================================================================
// ERROR UTILITIES
// ============================================================================

// ErrorSerializer serializes errors to different formats
type ErrorSerializer struct{}

func NewErrorSerializer() *ErrorSerializer {
	return &ErrorSerializer{}
}

func (es *ErrorSerializer) ToJSON(err error) ([]byte, error) {
	return json.Marshal(err)
}

func (es *ErrorSerializer) ToMap(err error) map[string]interface{} {
	switch e := err.(type) {
	case ValidationError:
		return map[string]interface{}{
			"type":    "ValidationError",
			"field":   e.Field,
			"value":   e.Value,
			"message": e.Message,
			"code":    e.Code,
		}
	case BusinessError:
		return map[string]interface{}{
			"type":      "BusinessError",
			"operation": e.Operation,
			"resource":  e.Resource,
			"message":   e.Message,
			"code":      e.Code,
			"timestamp": e.Timestamp,
		}
	case SystemError:
		return map[string]interface{}{
			"type":      "SystemError",
			"component": e.Component,
			"operation": e.Operation,
			"message":   e.Message,
			"code":      e.Code,
			"severity":  e.Severity,
			"timestamp": e.Timestamp,
		}
	case APIError:
		return map[string]interface{}{
			"type":           "APIError",
			"status_code":    e.StatusCode,
			"error_code":     e.ErrorCode,
			"message":        e.Message,
			"request_id":     e.RequestID,
			"correlation_id": e.CorrelationID,
			"timestamp":      e.Timestamp,
		}
	default:
		return map[string]interface{}{
			"type":    "UnknownError",
			"message": err.Error(),
		}
	}
}

// ErrorComparator compares errors for equality
type ErrorComparator struct{}

func NewErrorComparator() *ErrorComparator {
	return &ErrorComparator{}
}

func (ec *ErrorComparator) IsSameType(err1, err2 error) bool {
	return fmt.Sprintf("%T", err1) == fmt.Sprintf("%T", err2)
}

func (ec *ErrorComparator) IsSameCode(err1, err2 error) bool {
	coded1, ok1 := err1.(CodedError)
	coded2, ok2 := err2.(CodedError)
	
	if !ok1 || !ok2 {
		return false
	}
	
	return coded1.Code() == coded2.Code()
}

func (ec *ErrorComparator) IsSameSeverity(err1, err2 error) bool {
	severity1, ok1 := err1.(CodedError)
	severity2, ok2 := err2.(CodedError)
	
	if !ok1 || !ok2 {
		return false
	}
	
	return severity1.Severity() == severity2.Severity()
}

// ============================================================================
// ERROR METRICS
// ============================================================================

// ErrorMetrics tracks error patterns and frequencies
type ErrorMetrics struct {
	ErrorCounts map[string]int       `json:"error_counts"`
	ErrorRates  map[string]float64   `json:"error_rates"`
	LastErrors  map[string]time.Time `json:"last_errors"`
}

func NewErrorMetrics() *ErrorMetrics {
	return &ErrorMetrics{
		ErrorCounts: make(map[string]int),
		ErrorRates:  make(map[string]float64),
		LastErrors:  make(map[string]time.Time),
	}
}

func (em *ErrorMetrics) RecordError(err error) {
	errorType := fmt.Sprintf("%T", err)
	em.ErrorCounts[errorType]++
	em.LastErrors[errorType] = time.Now()
	
	// Calculate error rate (simplified)
	totalErrors := 0
	for _, count := range em.ErrorCounts {
		totalErrors += count
	}
	
	if totalErrors > 0 {
		em.ErrorRates[errorType] = float64(em.ErrorCounts[errorType]) / float64(totalErrors)
	}
}

func (em *ErrorMetrics) GetErrorStats() map[string]interface{} {
	return map[string]interface{}{
		"error_counts": em.ErrorCounts,
		"error_rates":  em.ErrorRates,
		"last_errors":  em.LastErrors,
	}
}

// ============================================================================
// DEMONSTRATION FUNCTIONS
// ============================================================================

func demonstrateCustomErrors() {
	fmt.Println("🚨 CUSTOM ERROR TYPES MASTERY")
	fmt.Println("=============================")
	fmt.Println()
	
	// Create error factory
	factory := NewErrorFactory("user-service", "1.0.0")
	
	// Demonstrate basic custom errors
	fmt.Println("1. Basic Custom Error Types:")
	fmt.Println("----------------------------")
	
	// Validation error
	validationErr := factory.NewValidationError("email", "invalid email format", "INVALID_EMAIL", "not-an-email")
	fmt.Printf("   📊 Validation Error: %v\n", validationErr)
	
	// Business error
	businessErr := factory.NewBusinessError("create", "user", "user already exists", "USER_EXISTS")
	fmt.Printf("   📊 Business Error: %v\n", businessErr)
	
	// System error
	systemErr := factory.NewSystemError("database", "connect", "connection pool exhausted", "POOL_EXHAUSTED", SeverityHigh)
	fmt.Printf("   📊 System Error: %v\n", systemErr)
	
	fmt.Println()
	
	// Demonstrate error hierarchies
	fmt.Println("2. Error Hierarchies:")
	fmt.Println("--------------------")
	
	// Database error
	dbErr := factory.NewDatabaseError("select", "users", "query failed", "QUERY_FAILED", fmt.Errorf("syntax error"))
	fmt.Printf("   📊 Database Error: %v\n", dbErr)
	
	// Connection error
	connErr := ConnectionError{
		DatabaseError: DatabaseError{
			Operation: "connect",
			Table:     "users",
			Message:   "connection refused",
			Code:      "CONNECTION_REFUSED",
			Timestamp: time.Now(),
		},
		Host:    "localhost",
		Port:    5432,
		Timeout: 30,
	}
	fmt.Printf("   📊 Connection Error: %v\n", connErr)
	
	// Query error
	queryErr := QueryError{
		DatabaseError: DatabaseError{
			Operation: "select",
			Table:     "users",
			Message:   "syntax error",
			Code:      "SYNTAX_ERROR",
			Timestamp: time.Now(),
		},
		SQL: "SELECT * FROM users WHERE id = ?",
	}
	fmt.Printf("   📊 Query Error: %v\n", queryErr)
	
	fmt.Println()
	
	// Demonstrate error interfaces
	fmt.Println("3. Error Interfaces:")
	fmt.Println("-------------------")
	
	// API error (implements CodedError)
	apiErr := factory.NewAPIError(400, "VALIDATION_ERROR", "Invalid request data", "req-123")
	fmt.Printf("   📊 API Error: %v\n", apiErr)
	fmt.Printf("   📊 Error Code: %s\n", apiErr.Code())
	fmt.Printf("   📊 Severity: %s\n", apiErr.Severity())
	
	// Network error (implements RecoverableError)
	networkErr := factory.NewNetworkError("GET", "https://api.example.com/users", "GET", "timeout", "TIMEOUT", 0, 30*time.Second, fmt.Errorf("context deadline exceeded"))
	fmt.Printf("   📊 Network Error: %v\n", networkErr)
	fmt.Printf("   📊 Can Recover: %t\n", networkErr.CanRecover())
	fmt.Printf("   📊 Recovery Strategy: %s\n", networkErr.RecoveryStrategy())
	fmt.Printf("   📊 Retry After: %v\n", networkErr.RetryAfter())
	
	fmt.Println()
	
	// Demonstrate error serialization
	fmt.Println("4. Error Serialization:")
	fmt.Println("----------------------")
	
	serializer := NewErrorSerializer()
	
	// Serialize to JSON
	jsonData, _ := serializer.ToJSON(validationErr)
	fmt.Printf("   📊 JSON: %s\n", string(jsonData))
	
	// Serialize to map
	errorMap := serializer.ToMap(businessErr)
	fmt.Printf("   📊 Map: %+v\n", errorMap)
	
	fmt.Println()
	
	// Demonstrate error comparison
	fmt.Println("5. Error Comparison:")
	fmt.Println("-------------------")
	
	comparator := NewErrorComparator()
	
	fmt.Printf("   📊 Same Type (validation vs business): %t\n", 
		comparator.IsSameType(validationErr, businessErr))
	fmt.Printf("   📊 Same Type (validation vs validation): %t\n", 
		comparator.IsSameType(validationErr, factory.NewValidationError("name", "required", "REQUIRED", nil)))
	
	fmt.Println()
	
	// Demonstrate error metrics
	fmt.Println("6. Error Metrics:")
	fmt.Println("----------------")
	
	metrics := NewErrorMetrics()
	metrics.RecordError(validationErr)
	metrics.RecordError(businessErr)
	metrics.RecordError(systemErr)
	metrics.RecordError(validationErr) // Record again
	
	stats := metrics.GetErrorStats()
	fmt.Printf("   📊 Error Stats: %+v\n", stats)
	
	fmt.Println()
}

func demonstrateErrorHandling() {
	fmt.Println("🔧 ERROR HANDLING PATTERNS:")
	fmt.Println("===========================")
	
	// Create error factory
	factory := NewErrorFactory("user-service", "1.0.0")
	
	// Error handling with type switches
	errors := []error{
		factory.NewValidationError("email", "invalid", "INVALID_EMAIL", "test"),
		factory.NewBusinessError("create", "user", "already exists", "USER_EXISTS"),
		factory.NewSystemError("database", "connect", "timeout", "TIMEOUT", SeverityHigh),
	}
	
	for i, err := range errors {
		fmt.Printf("   📊 Error %d: ", i+1)
		
		switch e := err.(type) {
		case ValidationError:
			fmt.Printf("Validation Error - Field: %s, Code: %s\n", e.Field, e.Code)
		case BusinessError:
			fmt.Printf("Business Error - Operation: %s, Resource: %s, Code: %s\n", e.Operation, e.Resource, e.Code)
		case SystemError:
			fmt.Printf("System Error - Component: %s, Severity: %s, Code: %s\n", e.Component, e.Severity, e.Code)
		default:
			fmt.Printf("Unknown Error: %v\n", err)
		}
	}
	
	fmt.Println()
}

func demonstrateErrorRecovery() {
	fmt.Println("🔄 ERROR RECOVERY STRATEGIES:")
	fmt.Println("=============================")
	
	// Create error factory
	factory := NewErrorFactory("user-service", "1.0.0")
	
	// Simulate recoverable errors
	recoverableErrors := []error{
		factory.NewNetworkError("GET", "https://api.example.com", "GET", "timeout", "TIMEOUT", 0, 30*time.Second, fmt.Errorf("timeout")),
		factory.NewNetworkError("POST", "https://api.example.com", "POST", "rate limited", "RATE_LIMITED", 429, 30*time.Second, fmt.Errorf("rate limited")),
		factory.NewNetworkError("GET", "https://api.example.com", "GET", "server error", "SERVER_ERROR", 500, 30*time.Second, fmt.Errorf("internal server error")),
	}
	
	for i, err := range recoverableErrors {
		fmt.Printf("   📊 Error %d: %v\n", i+1, err)
		
		if recoverable, ok := err.(RecoverableError); ok {
			fmt.Printf("      ✅ Can Recover: %t\n", recoverable.CanRecover())
			fmt.Printf("      📋 Strategy: %s\n", recoverable.RecoveryStrategy())
			fmt.Printf("      ⏰ Retry After: %v\n", recoverable.RetryAfter())
		} else {
			fmt.Printf("      ❌ Cannot Recover\n")
		}
		fmt.Println()
	}
}

// ============================================================================
// MAIN DEMONSTRATION
// ============================================================================

func main() {
	fmt.Println("🚨 CUSTOM ERROR TYPES MASTERY")
	fmt.Println("=============================")
	fmt.Println()
	
	// Demonstrate custom errors
	demonstrateCustomErrors()
	
	// Demonstrate error handling patterns
	demonstrateErrorHandling()
	
	// Demonstrate error recovery strategies
	demonstrateErrorRecovery()
	
	fmt.Println("🎉 CUSTOM ERROR TYPES MASTERY COMPLETE!")
	fmt.Println("You have successfully demonstrated:")
	fmt.Println("✅ Custom error type design")
	fmt.Println("✅ Error hierarchies and inheritance")
	fmt.Println("✅ Error interfaces and contracts")
	fmt.Println("✅ Error factories and creation patterns")
	fmt.Println("✅ Error serialization and comparison")
	fmt.Println("✅ Error metrics and tracking")
	fmt.Println("✅ Error recovery strategies")
	fmt.Println()
	fmt.Println("🚀 You are now ready for Error Wrapping Mastery!")
}
