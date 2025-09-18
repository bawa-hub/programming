package main

import (
	"context"
	"fmt"
	"time"
)

// Custom type for context keys to avoid collisions
type contextKey string

const (
	userIDKey    contextKey = "userID"
	requestIDKey contextKey = "requestID"
)

func processRequest(ctx context.Context, requestID string) {
	// Add request ID to context
	ctx = context.WithValue(ctx, requestIDKey, requestID)
	
	fmt.Printf("Processing request %s\n", requestID)
	
	// Simulate some work
	time.Sleep(500 * time.Millisecond)
	
	// Call another function that needs the context
	processUserData(ctx)
}

func processUserData(ctx context.Context) {
	// Get request ID from context
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		fmt.Println("No request ID found in context")
		return
	}
	
	fmt.Printf("Processing user data for request %s\n", requestID)
	
	// Simulate database call
	time.Sleep(300 * time.Millisecond)
	
	// Call another function
	logActivity(ctx)
}

func logActivity(ctx context.Context) {
	// Get request ID from context
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		fmt.Println("No request ID found in context")
		return
	}
	
	fmt.Printf("Logging activity for request %s\n", requestID)
}

func main() {
	fmt.Println("=== Context Values Example ===")
	
	// Create base context
	ctx := context.Background()
	
	// Add user ID to context
	ctx = context.WithValue(ctx, userIDKey, "user123")
	
	// Process multiple requests
	requests := []string{"req001", "req002", "req003"}
	
	for _, reqID := range requests {
		go processRequest(ctx, reqID)
	}
	
	// Wait for all requests to complete
	time.Sleep(2 * time.Second)
	
	fmt.Println("All requests processed!")
}
