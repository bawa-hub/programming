package main

import (
    "chain/handler"
    "chain/middleware"
    "chain/validation"
    "chain/logging"
    "fmt"
    "time"
)

func main() {
	fmt.Println("=== CHAIN OF RESPONSIBILITY PATTERN DEMONSTRATION ===\n")

	// // 1. BASIC CHAIN OF RESPONSIBILITY
	fmt.Println("1. BASIC CHAIN OF RESPONSIBILITY:")
	chainBuilder := handler.NewChainBuilder()
	chainBuilder.AddHandler(handler.NewLoggingHandler())
	chainBuilder.AddHandler(handler.NewAuthenticationHandler())
	chainBuilder.AddHandler(handler.NewAuthorizationHandler())
	chainBuilder.AddHandler(handler.NewValidationHandler())
	
	chain := chainBuilder.Build()
	
	// Test different request types
	requests := []handler.Request{
		handler.NewConcreteRequest("AUTH", "user123", 1),
		handler.NewConcreteRequest("API", "api_data", 2),
		handler.NewConcreteRequest("DATA", "data_payload", 3),
		handler.NewConcreteRequest("UNKNOWN", "unknown_data", 4),
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

    // // 2. WEB MIDDLEWARE CHAIN
    fmt.Println("2. WEB MIDDLEWARE CHAIN:")
    // Build middleware chain using a builder
    mwBuilder := middleware.NewMiddlewareChainBuilder()
    mwBuilder.
        Add(middleware.NewCORSMiddleware()).
        Add(middleware.NewLoggingMiddleware()).
        Add(middleware.NewRateLimitMiddleware(5)).
        Add(middleware.NewAuthMiddleware())
    middlewareChain := mwBuilder.Build()

    // Test HTTP request
    httpRequest := &middleware.HTTPRequest{
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
    
    httpResponse := &middleware.HTTPResponse{
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

	// // 3. VALIDATION CHAIN
    fmt.Println("3. VALIDATION CHAIN:")
    vBuilder := validation.NewValidationChainBuilder()
    vBuilder.
        Add(validation.NewRequiredFieldValidationRule()).
        Add(validation.NewLengthValidationRule(5, 50)).
        Add(validation.NewEmailValidationRule())
    vChain := vBuilder.Build()
	
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
        ok, msg := vChain.Validate(data)
        if ok {
            fmt.Println("  Validation passed")
        } else {
            fmt.Printf("  Validation failed: %s\n", msg)
        }
    }
	fmt.Println()

    // 4. LOGGING CHAIN
    fmt.Println("4. LOGGING CHAIN:")
    lb := logging.NewLoggingChainBuilder()
    lb.
        Add(logging.NewConsoleLogHandler()).
        Add(logging.NewFileLogHandler("app.log")).
        Add(logging.NewEmailLogHandler("admin@example.com"))
    logChain := lb.Build()

    // Test different log levels
    logEntries := []logging.LogEntry{
        {Level: logging.DEBUG, Message: "Debug message", Timestamp: time.Now(), Source: "app"},
        {Level: logging.INFO, Message: "Info message", Timestamp: time.Now(), Source: "app"},
        {Level: logging.WARN, Message: "Warning message", Timestamp: time.Now(), Source: "app"},
        {Level: logging.ERROR, Message: "Error message", Timestamp: time.Now(), Source: "app"},
        {Level: logging.FATAL, Message: "Fatal message", Timestamp: time.Now(), Source: "app"},
    }

    for _, entry := range logEntries {
        fmt.Printf("Processing log entry: %s\n", entry.Message)
        logChain.Handle(entry)
    }
    fmt.Println()

	// // 5. DYNAMIC CHAIN MODIFICATION
	fmt.Println("5. DYNAMIC CHAIN MODIFICATION:")
	dynamicChain := handler.NewChainBuilder()
	dynamicChain.AddHandler(handler.NewLoggingHandler())
	dynamicChain.AddHandler(handler.NewAuthenticationHandler())
	
	// Build initial chain
	initialChain := dynamicChain.Build()
	
	// Test with initial chain
	request := handler.NewConcreteRequest("AUTH", "user123", 1)
	fmt.Println("Testing with initial chain:")
	initialChain.Handle(request)
	
	// Add more handlers dynamically
	dynamicChain.AddHandler(handler.NewAuthorizationHandler())
	dynamicChain.AddHandler(handler.NewValidationHandler())
	
	// Build new chain
	newChain := dynamicChain.Build()
	
	// Test with new chain
	request2 := handler.NewConcreteRequest("API", "api_data", 2)
	fmt.Println("\nTesting with extended chain:")
	newChain.Handle(request2)
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
