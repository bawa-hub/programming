package exercises

import (
	"context"
	"fmt"
	"time"
)

// Exercise 3: Context Cancellation Chain
func Exercise3() {
	fmt.Println("\nExercise 3: Context Cancellation Chain")
	fmt.Println("======================================")
	
	// TODO: Create a context hierarchy and test cancellation
	// 1. Create parent context with timeout
	// 2. Create child context with different timeout
	// 3. Create grandchild context
	// 4. Test cancellation propagation
	
	parentCtx, parentCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer parentCancel()
	
	childCtx, childCancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer childCancel()
	
	grandchildCtx := context.WithValue(childCtx, "level", "grandchild")
	
	// Start goroutines for each level
	go processLevel("Parent", parentCtx)
	go processLevel("Child", childCtx)
	go processLevel("Grandchild", grandchildCtx)
	
	time.Sleep(4 * time.Second)
	fmt.Println("Exercise 3 completed")
}

func processLevel(name string, ctx context.Context) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  Exercise 3: %s cancelled: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("  Exercise 3: %s working... %d\n", name, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Printf("  Exercise 3: %s completed\n", name)
}