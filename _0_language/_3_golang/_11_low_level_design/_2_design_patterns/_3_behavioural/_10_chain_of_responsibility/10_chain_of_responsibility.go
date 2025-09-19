package main

import (
	"fmt"
	"strings"
	"time"
)

// =============================================================================
// BASIC CHAIN OF RESPONSIBILITY PATTERN
// =============================================================================

// Request interface
type Request interface {
	GetType() string
	GetData() interface{}
	GetPriority() int
	IsProcessed() bool
	SetProcessed(processed bool)
}

// Concrete Request
type ConcreteRequest struct {
	requestType string
	data        interface{}
	priority    int
	processed   bool
}

func NewConcreteRequest(requestType string, data interface{}, priority int) *ConcreteRequest {
	return &ConcreteRequest{
		requestType: requestType,
		data:        data,
		priority:    priority,
		processed:   false,
	}
}

func (cr *ConcreteRequest) GetType() string {
	return cr.requestType
}

func (cr *ConcreteRequest) GetData() interface{} {
	return cr.data
}

func (cr *ConcreteRequest) GetPriority() int {
	return cr.priority
}

func (cr *ConcreteRequest) IsProcessed() bool {
	return cr.processed
}

func (cr *ConcreteRequest) SetProcessed(processed bool) {
	cr.processed = processed
}

// Handler interface
type Handler interface {
	Handle(request Request) bool
	SetNext(handler Handler)
	CanHandle(request Request) bool
}

// Abstract Handler
type AbstractHandler struct {
	next Handler
}

func (ah *AbstractHandler) SetNext(handler Handler) {
	ah.next = handler
}

func (ah *AbstractHandler) Handle(request Request) bool {
	if ah.next != nil {
		return ah.next.Handle(request)
	}
	return false
}

// Concrete Handlers
type AuthenticationHandler struct {
	AbstractHandler
}

func NewAuthenticationHandler() *AuthenticationHandler {
	return &AuthenticationHandler{}
}

func (ah *AuthenticationHandler) CanHandle(request Request) bool {
	return request.GetType() == "AUTH" || request.GetType() == "API"
}

func (ah *AuthenticationHandler) Handle(request Request) bool {
	if ah.CanHandle(request) {
		fmt.Printf("AuthenticationHandler: Processing %s request\n", request.GetType())
		// Simulate authentication logic
		time.Sleep(100 * time.Millisecond)
		request.SetProcessed(true)
		return true
	}
	return ah.AbstractHandler.Handle(request)
}

type AuthorizationHandler struct {
	AbstractHandler
}

func NewAuthorizationHandler() *AuthorizationHandler {
	return &AuthorizationHandler{}
}

func (azh *AuthorizationHandler) CanHandle(request Request) bool {
	return request.GetType() == "AUTH" || request.GetType() == "API" || request.GetType() == "ADMIN"
}

func (azh *AuthorizationHandler) Handle(request Request) bool {
	if azh.CanHandle(request) {
		fmt.Printf("AuthorizationHandler: Processing %s request\n", request.GetType())
		// Simulate authorization logic
		time.Sleep(50 * time.Millisecond)
		request.SetProcessed(true)
		return true
	}
	return azh.AbstractHandler.Handle(request)
}

type ValidationHandler struct {
	AbstractHandler
}

func NewValidationHandler() *ValidationHandler {
	return &ValidationHandler{}
}

func (vh *ValidationHandler) CanHandle(request Request) bool {
	return request.GetType() == "API" || request.GetType() == "DATA"
}

func (vh *ValidationHandler) Handle(request Request) bool {
	if vh.CanHandle(request) {
		fmt.Printf("ValidationHandler: Processing %s request\n", request.GetType())
		// Simulate validation logic
		time.Sleep(75 * time.Millisecond)
		request.SetProcessed(true)
		return true
	}
	return vh.AbstractHandler.Handle(request)
}

type LoggingHandler struct {
	AbstractHandler
}

func NewLoggingHandler() *LoggingHandler {
	return &LoggingHandler{}
}

func (lh *LoggingHandler) CanHandle(request Request) bool {
	return true // Logging handler can handle all requests
}

func (lh *LoggingHandler) Handle(request Request) bool {
	fmt.Printf("LoggingHandler: Logging %s request with priority %d\n", 
		request.GetType(), request.GetPriority())
	// Always pass to next handler
	return lh.AbstractHandler.Handle(request)
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. WEB MIDDLEWARE CHAIN
type HTTPRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
	User    string
	Role    string
}

type HTTPResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

type Middleware interface {
	Process(request *HTTPRequest, response *HTTPResponse) bool
	SetNext(middleware Middleware)
}

type AbstractMiddleware struct {
	next Middleware
}

func (am *AbstractMiddleware) SetNext(middleware Middleware) {
	am.next = middleware
}

func (am *AbstractMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	if am.next != nil {
		return am.next.Process(request, response)
	}
	return true
}

// CORS Middleware
type CORSMiddleware struct {
	AbstractMiddleware
}

func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}

func (cm *CORSMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	fmt.Println("CORSMiddleware: Adding CORS headers")
	response.Headers["Access-Control-Allow-Origin"] = "*"
	response.Headers["Access-Control-Allow-Methods"] = "GET, POST, PUT, DELETE"
	response.Headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
	return cm.AbstractMiddleware.Process(request, response)
}

// Authentication Middleware
type AuthMiddleware struct {
	AbstractMiddleware
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	fmt.Println("AuthMiddleware: Checking authentication")
	if request.Headers["Authorization"] == "" {
		response.StatusCode = 401
		response.Body = "Unauthorized"
		return false
	}
	return am.AbstractMiddleware.Process(request, response)
}

// Rate Limiting Middleware
type RateLimitMiddleware struct {
	AbstractMiddleware
	requests map[string]int
	limit    int
}

func NewRateLimitMiddleware(limit int) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		requests: make(map[string]int),
		limit:    limit,
	}
}

func (rlm *RateLimitMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	fmt.Println("RateLimitMiddleware: Checking rate limit")
	clientIP := request.Headers["X-Forwarded-For"]
	if clientIP == "" {
		clientIP = "unknown"
	}
	
	rlm.requests[clientIP]++
	if rlm.requests[clientIP] > rlm.limit {
		response.StatusCode = 429
		response.Body = "Too Many Requests"
		return false
	}
	return rlm.AbstractMiddleware.Process(request, response)
}

// Logging Middleware
type LoggingMiddleware struct {
	AbstractMiddleware
}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (lm *LoggingMiddleware) Process(request *HTTPRequest, response *HTTPResponse) bool {
	fmt.Printf("LoggingMiddleware: %s %s - User: %s\n", 
		request.Method, request.URL, request.User)
	return lm.AbstractMiddleware.Process(request, response)
}

// 2. VALIDATION CHAIN
type ValidationRule interface {
	Validate(data interface{}) (bool, string)
	SetNext(rule ValidationRule)
}

type AbstractValidationRule struct {
	next ValidationRule
}

func (avr *AbstractValidationRule) SetNext(rule ValidationRule) {
	avr.next = rule
}

func (avr *AbstractValidationRule) Validate(data interface{}) (bool, string) {
	if avr.next != nil {
		return avr.next.Validate(data)
	}
	return true, ""
}

// Email Validation Rule
type EmailValidationRule struct {
	AbstractValidationRule
}

func NewEmailValidationRule() *EmailValidationRule {
	return &EmailValidationRule{}
}

func (evr *EmailValidationRule) Validate(data interface{}) (bool, string) {
	email, ok := data.(string)
	if !ok {
		return false, "Data is not a string"
	}
	
	if !strings.Contains(email, "@") {
		return false, "Invalid email format"
	}
	
	return evr.AbstractValidationRule.Validate(data)
}

// Length Validation Rule
type LengthValidationRule struct {
	AbstractValidationRule
	minLength int
	maxLength int
}

func NewLengthValidationRule(minLength, maxLength int) *LengthValidationRule {
	return &LengthValidationRule{
		minLength: minLength,
		maxLength: maxLength,
	}
}

func (lvr *LengthValidationRule) Validate(data interface{}) (bool, string) {
	text, ok := data.(string)
	if !ok {
		return false, "Data is not a string"
	}
	
	if len(text) < lvr.minLength {
		return false, fmt.Sprintf("Text too short (minimum %d characters)", lvr.minLength)
	}
	
	if len(text) > lvr.maxLength {
		return false, fmt.Sprintf("Text too long (maximum %d characters)", lvr.maxLength)
	}
	
	return lvr.AbstractValidationRule.Validate(data)
}

// Required Field Validation Rule
type RequiredFieldValidationRule struct {
	AbstractValidationRule
}

func NewRequiredFieldValidationRule() *RequiredFieldValidationRule {
	return &RequiredFieldValidationRule{}
}

func (rfvr *RequiredFieldValidationRule) Validate(data interface{}) (bool, string) {
	if data == nil {
		return false, "Field is required"
	}
	
	text, ok := data.(string)
	if ok && strings.TrimSpace(text) == "" {
		return false, "Field cannot be empty"
	}
	
	return rfvr.AbstractValidationRule.Validate(data)
}

// 3. LOGGING CHAIN
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type LogEntry struct {
	Level     LogLevel
	Message   string
	Timestamp time.Time
	Source    string
}

type LogHandler interface {
	Handle(entry LogEntry) bool
	SetNext(handler LogHandler)
	CanHandle(entry LogEntry) bool
}

type AbstractLogHandler struct {
	next LogHandler
}

func (alh *AbstractLogHandler) SetNext(handler LogHandler) {
	alh.next = handler
}

func (alh *AbstractLogHandler) Handle(entry LogEntry) bool {
	if alh.next != nil {
		return alh.next.Handle(entry)
	}
	return true
}

// Console Log Handler
type ConsoleLogHandler struct {
	AbstractLogHandler
}

func NewConsoleLogHandler() *ConsoleLogHandler {
	return &ConsoleLogHandler{}
}

func (clh *ConsoleLogHandler) CanHandle(entry LogEntry) bool {
	return entry.Level >= INFO
}

func (clh *ConsoleLogHandler) Handle(entry LogEntry) bool {
	if clh.CanHandle(entry) {
		fmt.Printf("[%s] %s: %s\n", 
			entry.Timestamp.Format("15:04:05"), 
			clh.getLevelString(entry.Level), 
			entry.Message)
	}
	return clh.AbstractLogHandler.Handle(entry)
}

func (clh *ConsoleLogHandler) getLevelString(level LogLevel) string {
	switch level {
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

// File Log Handler
type FileLogHandler struct {
	AbstractLogHandler
	filename string
}

func NewFileLogHandler(filename string) *FileLogHandler {
	return &FileLogHandler{filename: filename}
}

func (flh *FileLogHandler) CanHandle(entry LogEntry) bool {
	return entry.Level >= WARN
}

func (flh *FileLogHandler) Handle(entry LogEntry) bool {
	if flh.CanHandle(entry) {
		fmt.Printf("FileLogHandler: Writing to %s - [%s] %s\n", 
			flh.filename, 
			entry.Timestamp.Format("15:04:05"), 
			entry.Message)
	}
	return flh.AbstractLogHandler.Handle(entry)
}

// Email Log Handler
type EmailLogHandler struct {
	AbstractLogHandler
	email string
}

func NewEmailLogHandler(email string) *EmailLogHandler {
	return &EmailLogHandler{email: email}
}

func (elh *EmailLogHandler) CanHandle(entry LogEntry) bool {
	return entry.Level >= ERROR
}

func (elh *EmailLogHandler) Handle(entry LogEntry) bool {
	if elh.CanHandle(entry) {
		fmt.Printf("EmailLogHandler: Sending email to %s - [%s] %s\n", 
			elh.email, 
			entry.Timestamp.Format("15:04:05"), 
			entry.Message)
	}
	return elh.AbstractLogHandler.Handle(entry)
}

// =============================================================================
// CHAIN BUILDER
// =============================================================================

type ChainBuilder struct {
	handlers []Handler
}

func NewChainBuilder() *ChainBuilder {
	return &ChainBuilder{
		handlers: make([]Handler, 0),
	}
}

func (cb *ChainBuilder) AddHandler(handler Handler) *ChainBuilder {
	cb.handlers = append(cb.handlers, handler)
	return cb
}

func (cb *ChainBuilder) Build() Handler {
	if len(cb.handlers) == 0 {
		return nil
	}
	
	// Link handlers in sequence
	for i := 0; i < len(cb.handlers)-1; i++ {
		cb.handlers[i].SetNext(cb.handlers[i+1])
	}
	
	return cb.handlers[0]
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== CHAIN OF RESPONSIBILITY PATTERN DEMONSTRATION ===\n")

	// 1. BASIC CHAIN OF RESPONSIBILITY
	fmt.Println("1. BASIC CHAIN OF RESPONSIBILITY:")
	chainBuilder := NewChainBuilder()
	chainBuilder.AddHandler(NewLoggingHandler())
	chainBuilder.AddHandler(NewAuthenticationHandler())
	chainBuilder.AddHandler(NewAuthorizationHandler())
	chainBuilder.AddHandler(NewValidationHandler())
	
	chain := chainBuilder.Build()
	
	// Test different request types
	requests := []Request{
		NewConcreteRequest("AUTH", "user123", 1),
		NewConcreteRequest("API", "api_data", 2),
		NewConcreteRequest("DATA", "data_payload", 3),
		NewConcreteRequest("UNKNOWN", "unknown_data", 4),
	}
	
	for _, request := range requests {
		fmt.Printf("\nProcessing %s request:\n", request.GetType())
		success := chain.Handle(request)
		if success {
			fmt.Printf("Request processed successfully: %t\n", request.IsProcessed())
		} else {
			fmt.Println("Request could not be processed")
		}
	}
	fmt.Println()

	// 2. WEB MIDDLEWARE CHAIN
	fmt.Println("2. WEB MIDDLEWARE CHAIN:")
	middlewareBuilder := NewChainBuilder()
	middlewareBuilder.AddHandler(NewCORSMiddleware())
	middlewareBuilder.AddHandler(NewLoggingMiddleware())
	middlewareBuilder.AddHandler(NewRateLimitMiddleware(5))
	middlewareBuilder.AddHandler(NewAuthMiddleware())
	
	middlewareChain := middlewareBuilder.Build()
	
	// Test HTTP request
	httpRequest := &HTTPRequest{
		Method: "GET",
		URL:    "/api/users",
		Headers: map[string]string{
			"Authorization": "Bearer token123",
			"X-Forwarded-For": "192.168.1.1",
		},
		Body: "",
		User: "user123",
		Role: "admin",
	}
	
	httpResponse := &HTTPResponse{
		StatusCode: 200,
		Headers:    make(map[string]string),
		Body:       "",
	}
	
	fmt.Println("Processing HTTP request:")
	success := middlewareChain.Process(httpRequest, httpResponse)
	if success {
		fmt.Printf("Request processed successfully. Status: %d\n", httpResponse.StatusCode)
		fmt.Printf("Response headers: %v\n", httpResponse.Headers)
	} else {
		fmt.Printf("Request failed. Status: %d, Body: %s\n", httpResponse.StatusCode, httpResponse.Body)
	}
	fmt.Println()

	// 3. VALIDATION CHAIN
	fmt.Println("3. VALIDATION CHAIN:")
	validationBuilder := NewChainBuilder()
	validationBuilder.AddHandler(NewRequiredFieldValidationRule())
	validationBuilder.AddHandler(NewLengthValidationRule(5, 50))
	validationBuilder.AddHandler(NewEmailValidationRule())
	
	validationChain := validationBuilder.Build()
	
	// Test validation
	testData := []interface{}{
		"",                    // Empty string
		"a@b",                 // Too short
		"valid@email.com",     // Valid email
		"verylongemailaddressthatiswaytoomanycharacters@domain.com", // Too long
		"invalid-email",       // Invalid format
	}
	
	for i, data := range testData {
		fmt.Printf("Validating data %d: %v\n", i+1, data)
		success := validationChain.Validate(data)
		if success {
			fmt.Println("  Validation passed")
		} else {
			fmt.Println("  Validation failed")
		}
	}
	fmt.Println()

	// 4. LOGGING CHAIN
	fmt.Println("4. LOGGING CHAIN:")
	loggingBuilder := NewChainBuilder()
	loggingBuilder.AddHandler(NewConsoleLogHandler())
	loggingBuilder.AddHandler(NewFileLogHandler("app.log"))
	loggingBuilder.AddHandler(NewEmailLogHandler("admin@example.com"))
	
	loggingChain := loggingBuilder.Build()
	
	// Test different log levels
	logEntries := []LogEntry{
		{Level: DEBUG, Message: "Debug message", Timestamp: time.Now(), Source: "app"},
		{Level: INFO, Message: "Info message", Timestamp: time.Now(), Source: "app"},
		{Level: WARN, Message: "Warning message", Timestamp: time.Now(), Source: "app"},
		{Level: ERROR, Message: "Error message", Timestamp: time.Now(), Source: "app"},
		{Level: FATAL, Message: "Fatal message", Timestamp: time.Now(), Source: "app"},
	}
	
	for _, entry := range logEntries {
		fmt.Printf("Processing log entry: %s\n", entry.Message)
		loggingChain.Handle(entry)
	}
	fmt.Println()

	// 5. DYNAMIC CHAIN MODIFICATION
	fmt.Println("5. DYNAMIC CHAIN MODIFICATION:")
	dynamicChain := NewChainBuilder()
	dynamicChain.AddHandler(NewLoggingHandler())
	dynamicChain.AddHandler(NewAuthenticationHandler())
	
	// Build initial chain
	initialChain := dynamicChain.Build()
	
	// Test with initial chain
	request := NewConcreteRequest("AUTH", "user123", 1)
	fmt.Println("Testing with initial chain:")
	initialChain.Handle(request)
	
	// Add more handlers dynamically
	dynamicChain.AddHandler(NewAuthorizationHandler())
	dynamicChain.AddHandler(NewValidationHandler())
	
	// Build new chain
	newChain := dynamicChain.Build()
	
	// Test with new chain
	request2 := NewConcreteRequest("API", "api_data", 2)
	fmt.Println("\nTesting with extended chain:")
	newChain.Handle(request2)
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
