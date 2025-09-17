package main

import (
	"context"
	"fmt"
	"time"
)

// Simulate HTTP request processing
func handleRequest(ctx context.Context, requestID string) {
	fmt.Printf("Request %s: Starting processing\n", requestID)
	
	// Add request ID to context
	ctx = context.WithValue(ctx, "requestID", requestID)
	
	// Simulate different stages of request processing
	stages := []string{"authentication", "authorization", "business logic", "response formatting"}
	
	for i, stage := range stages {
		select {
		case <-ctx.Done():
			fmt.Printf("Request %s: Cancelled during %s\n", requestID, stage)
			return
		default:
			fmt.Printf("Request %s: Processing %s...\n", requestID, stage)
			time.Sleep(200 * time.Millisecond)
			
			// Simulate timeout for some requests
			if requestID == "req002" && i == 1 {
				fmt.Printf("Request %s: Simulating timeout\n", requestID)
				time.Sleep(1 * time.Second) // This will cause timeout
			}
		}
	}
	
	fmt.Printf("Request %s: Completed successfully\n", requestID)
}

func main() {
	fmt.Println("=== HTTP-like Request Processing with Context ===")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	// Process multiple requests
	requests := []string{"req001", "req002", "req003"}
	
	for _, reqID := range requests {
		go handleRequest(ctx, reqID)
	}
	
	// Wait for all requests to complete or timeout
	time.Sleep(2 * time.Second)
	
	fmt.Println("All requests processed!")
}
