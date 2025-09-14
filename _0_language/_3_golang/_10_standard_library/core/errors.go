package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// Custom error types
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}

type NetworkError struct {
	Operation string
	Host      string
	Err       error
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error during %s to %s: %v", e.Operation, e.Host, e.Err)
}

func (e NetworkError) Unwrap() error {
	return e.Err
}

type DatabaseError struct {
	Query string
	Err   error
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error executing query '%s': %v", e.Query, e.Err)
}

func (e DatabaseError) Unwrap() error {
	return e.Err
}

func main() {
	fmt.Println("🚀 Go errors Package Mastery Examples")
	fmt.Println("=====================================")

	// 1. Basic Error Creation
	fmt.Println("\n1. Basic Error Creation:")
	
	// Create simple error
	err := errors.New("something went wrong")
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Error type: %T\n", err)
	
	// Create error with context
	err = fmt.Errorf("failed to process user %d: %w", 123, err)
	fmt.Printf("Wrapped error: %v\n", err)

	// 2. Error Wrapping
	fmt.Println("\n2. Error Wrapping:")
	
	// Simulate file operation error
	fileErr := simulateFileOperation("nonexistent.txt")
	if fileErr != nil {
		// Wrap with context
		wrappedErr := fmt.Errorf("failed to process file: %w", fileErr)
		fmt.Printf("Wrapped error: %v\n", wrappedErr)
		
		// Unwrap error
		originalErr := errors.Unwrap(wrappedErr)
		fmt.Printf("Original error: %v\n", originalErr)
	}

	// 3. Error Checking with Is()
	fmt.Println("\n3. Error Checking with Is():")
	
	// Check for specific error types
	if errors.Is(fileErr, os.ErrNotExist) {
		fmt.Println("File does not exist")
	}
	
	// Check wrapped error
	wrappedFileErr := fmt.Errorf("operation failed: %w", os.ErrNotExist)
	if errors.Is(wrappedFileErr, os.ErrNotExist) {
		fmt.Println("Wrapped error contains os.ErrNotExist")
	}

	// 4. Error Type Assertion with As()
	fmt.Println("\n4. Error Type Assertion with As():")
	
	// Create path error
	pathErr := &os.PathError{
		Op:   "open",
		Path: "test.txt",
		Err:  os.ErrNotExist,
	}
	
	// Check if error can be converted to PathError
	var pathError *os.PathError
	if errors.As(pathErr, &pathError) {
		fmt.Printf("Path error: %s %s: %v\n", pathError.Op, pathError.Path, pathError.Err)
	}

	// 5. Custom Error Types
	fmt.Println("\n5. Custom Error Types:")
	
	// Create validation error
	validationErr := ValidationError{
		Field:   "email",
		Message: "invalid email format",
	}
	fmt.Printf("Validation error: %v\n", validationErr)
	
	// Create network error
	networkErr := NetworkError{
		Operation: "GET",
		Host:      "api.example.com",
		Err:       io.EOF,
	}
	fmt.Printf("Network error: %v\n", networkErr)
	
	// Unwrap network error
	originalErr := errors.Unwrap(networkErr)
	fmt.Printf("Unwrapped error: %v\n", originalErr)

	// 6. Error Chains
	fmt.Println("\n6. Error Chains:")
	
	// Create error chain
	chainErr := createErrorChain()
	fmt.Printf("Error chain: %v\n", chainErr)
	
	// Check if chain contains specific error
	if errors.Is(chainErr, os.ErrNotExist) {
		fmt.Println("Error chain contains os.ErrNotExist")
	}
	
	// Unwrap entire chain
	fmt.Println("Unwrapping error chain:")
	currentErr := chainErr
	for i := 0; currentErr != nil && i < 5; i++ {
		fmt.Printf("  Level %d: %v\n", i, currentErr)
		currentErr = errors.Unwrap(currentErr)
	}

	// 7. Error Joining
	fmt.Println("\n7. Error Joining:")
	
	// Create multiple errors
	err1 := errors.New("first error")
	err2 := errors.New("second error")
	err3 := errors.New("third error")
	
	// Join errors
	joinedErr := errors.Join(err1, err2, err3)
	fmt.Printf("Joined errors: %v\n", joinedErr)
	
	// Check if joined errors contain specific error
	if errors.Is(joinedErr, err2) {
		fmt.Println("Joined errors contain err2")
	}

	// 8. Error Handling Patterns
	fmt.Println("\n8. Error Handling Patterns:")
	
	// Pattern 1: Early return
	if err := validateInput(""); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	}
	
	// Pattern 2: Error wrapping
	if err := processData("test"); err != nil {
		wrappedErr := fmt.Errorf("data processing failed: %w", err)
		fmt.Printf("Wrapped error: %v\n", wrappedErr)
	}
	
	// Pattern 3: Error logging
	if err := riskyOperation(); err != nil {
		fmt.Printf("Risky operation failed: %v\n", err)
		// In real application, you would log this
	}

	// 9. Error Recovery
	fmt.Println("\n9. Error Recovery:")
	
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	
	// Simulate panic
	// panic("something went wrong")
	fmt.Println("Panic simulation skipped for demo")

	// 10. Error Context
	fmt.Println("\n10. Error Context:")
	
	// Add context to errors
	ctxErr := addContextToError("database operation", "SELECT * FROM users")
	fmt.Printf("Context error: %v\n", ctxErr)
	
	// Extract context
	if dbErr, ok := ctxErr.(*DatabaseError); ok {
		fmt.Printf("Database query: %s\n", dbErr.Query)
		fmt.Printf("Original error: %v\n", dbErr.Err)
	}

	// 11. Error Metrics
	fmt.Println("\n11. Error Metrics:")
	
	// Count different error types
	errorCounts := make(map[string]int)
	errors := []error{
		os.ErrNotExist,
		io.EOF,
		ValidationError{Field: "name", Message: "required"},
		NetworkError{Operation: "POST", Host: "api.com", Err: io.EOF},
		os.ErrNotExist,
	}
	
	for _, err := range errors {
		errorType := fmt.Sprintf("%T", err)
		errorCounts[errorType]++
	}
	
	fmt.Println("Error counts:")
	for errorType, count := range errorCounts {
		fmt.Printf("  %s: %d\n", errorType, count)
	}

	// 12. Error Retry Logic
	fmt.Println("\n12. Error Retry Logic:")
	
	// Retry operation with exponential backoff
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err := simulateNetworkCall()
		if err == nil {
			fmt.Println("Network call succeeded")
			break
		}
		
		fmt.Printf("Attempt %d failed: %v\n", i+1, err)
		if i < maxRetries-1 {
			backoff := time.Duration(1<<uint(i)) * time.Second
			fmt.Printf("Retrying in %v...\n", backoff)
			time.Sleep(backoff)
		} else {
			fmt.Println("All retry attempts failed")
		}
	}

	// 13. Error Validation
	fmt.Println("\n13. Error Validation:")
	
	// Validate user input
	userInputs := []string{"", "invalid-email", "valid@email.com", "123"}
	for _, input := range userInputs {
		if err := validateUserInput(input); err != nil {
			fmt.Printf("Input '%s' validation failed: %v\n", input, err)
		} else {
			fmt.Printf("Input '%s' is valid\n", input)
		}
	}

	// 14. Error Aggregation
	fmt.Println("\n14. Error Aggregation:")
	
	// Collect multiple errors
	var errorList []error
	errorList = append(errorList, validateUserInput(""))
	errorList = append(errorList, validateUserInput("invalid"))
	errorList = append(errorList, validateUserInput("test@example.com"))
	
	// Filter out nil errors
	var actualErrors []error
	for _, err := range errorList {
		if err != nil {
			actualErrors = append(actualErrors, err)
		}
	}
	
	if len(actualErrors) > 0 {
		aggregatedErr := errors.Join(actualErrors...)
		fmt.Printf("Aggregated errors: %v\n", aggregatedErr)
	} else {
		fmt.Println("No validation errors")
	}

	// 15. Error Serialization
	fmt.Println("\n15. Error Serialization:")
	
	// Create error with structured data
	structuredErr := createStructuredError()
	fmt.Printf("Structured error: %v\n", structuredErr)
	
	// Extract structured information
	if validationErr, ok := structuredErr.(ValidationError); ok {
		fmt.Printf("Field: %s\n", validationErr.Field)
		fmt.Printf("Message: %s\n", validationErr.Message)
	}

	fmt.Println("\n🎉 errors Package Mastery Complete!")
}

// Helper functions

func simulateFileOperation(filename string) error {
	// Simulate file operation that fails
	return os.ErrNotExist
}

func createErrorChain() error {
	// Create a chain of wrapped errors
	err1 := os.ErrNotExist
	err2 := fmt.Errorf("file operation failed: %w", err1)
	err3 := fmt.Errorf("processing failed: %w", err2)
	return fmt.Errorf("application error: %w", err3)
}

func validateInput(input string) error {
	if input == "" {
		return ValidationError{Field: "input", Message: "cannot be empty"}
	}
	return nil
}

func processData(data string) error {
	if data == "error" {
		return errors.New("data processing error")
	}
	return nil
}

func riskyOperation() error {
	// Simulate risky operation
	return errors.New("risky operation failed")
}

func addContextToError(operation, query string) error {
	return &DatabaseError{
		Query: query,
		Err:   errors.New("connection timeout"),
	}
}

func simulateNetworkCall() error {
	// Simulate network call that sometimes fails
	if time.Now().UnixNano()%2 == 0 {
		return NetworkError{
			Operation: "GET",
			Host:      "api.example.com",
			Err:       errors.New("timeout"),
		}
	}
	return nil
}

func validateUserInput(input string) error {
	if input == "" {
		return ValidationError{Field: "input", Message: "cannot be empty"}
	}
	if input == "invalid" {
		return ValidationError{Field: "input", Message: "invalid format"}
	}
	if input == "123" {
		return ValidationError{Field: "input", Message: "must be a string"}
	}
	return nil
}

func createStructuredError() error {
	return ValidationError{
		Field:   "email",
		Message: "invalid email format",
	}
}

// Advanced error handling utilities

// ErrorCollector collects multiple errors
type ErrorCollector struct {
	errors []error
}

func (ec *ErrorCollector) Add(err error) {
	if err != nil {
		ec.errors = append(ec.errors, err)
	}
}

func (ec *ErrorCollector) HasErrors() bool {
	return len(ec.errors) > 0
}

func (ec *ErrorCollector) Error() error {
	if len(ec.errors) == 0 {
		return nil
	}
	return errors.Join(ec.errors...)
}

// RetryableError represents an error that can be retried
type RetryableError struct {
	Err      error
	RetryAfter time.Duration
}

func (e RetryableError) Error() string {
	return fmt.Sprintf("retryable error: %v (retry after %v)", e.Err, e.RetryAfter)
}

func (e RetryableError) Unwrap() error {
	return e.Err
}

// ErrorWithCode represents an error with a specific code
type ErrorWithCode struct {
	Code    int
	Message string
	Err     error
}

func (e ErrorWithCode) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error %d: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func (e ErrorWithCode) Unwrap() error {
	return e.Err
}

// ErrorWithTimestamp represents an error with timestamp
type ErrorWithTimestamp struct {
	Timestamp time.Time
	Err       error
}

func (e ErrorWithTimestamp) Error() string {
	return fmt.Sprintf("[%s] %v", e.Timestamp.Format(time.RFC3339), e.Err)
}

func (e ErrorWithTimestamp) Unwrap() error {
	return e.Err
}
