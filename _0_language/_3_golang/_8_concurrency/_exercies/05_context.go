package main

import (
	"context"
	"fmt"
	"time"
)

// TODO: Create a function that simulates a database query
// TODO: The function should respect context cancellation and timeout
// TODO: Use context.WithTimeout to limit the query to 2 seconds
// TODO: Add request ID to context and use it in logging
// TODO: Start multiple goroutines with different request IDs
// TODO: Some queries should complete, others should timeout

func main() {
	fmt.Println("=== Context Exercise ===")
	
	// Your code goes here:
	// 1. Create a databaseQuery function that takes context and requestID
	// 2. Use context.WithTimeout for 2-second limit
	// 3. Add request ID to context using WithValue
	// 4. Start multiple goroutines with different request IDs
	// 5. Some should complete, others should timeout
	
	fmt.Println("Exercise completed!")
}
