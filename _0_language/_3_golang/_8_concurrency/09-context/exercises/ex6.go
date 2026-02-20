package exercises

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Exercise 6: Context with Multiple Goroutines
func Exercise6() {
	fmt.Println("\nExercise 6: Context with Multiple Goroutines")
	fmt.Println("===========================================")
	
	// TODO: Use context to coordinate multiple goroutines
	// 1. Create context with timeout
	// 2. Start multiple goroutines
	// 3. Use WaitGroup to wait for completion
	// 4. Handle cancellation in all goroutines
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	var wg sync.WaitGroup
	numGoroutines := 3
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < 10; j++ {
				select {
				case <-ctx.Done():
					fmt.Printf("  Exercise 6: Goroutine %d cancelled: %v\n", id, ctx.Err())
					return
				default:
					fmt.Printf("  Exercise 6: Goroutine %d working... %d\n", id, j)
					time.Sleep(200 * time.Millisecond)
				}
			}
			fmt.Printf("  Exercise 6: Goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("Exercise 6 completed")
}