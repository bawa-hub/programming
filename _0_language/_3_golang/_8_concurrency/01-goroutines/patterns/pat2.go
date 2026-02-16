package patterns

import (
	"context"
	"fmt"
	"time"
)

// Advanced Pattern 2: Goroutine with Context Cancellation
func ContextGoroutine() {
	fmt.Println("Advanced Pattern 2: Context Cancellation")
	fmt.Println("========================================")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Start goroutine with context
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine: Context cancelled, exiting")
				return
			default:
				fmt.Println("Goroutine: Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	
	// Let it run for a bit
	time.Sleep(3 * time.Second)
}