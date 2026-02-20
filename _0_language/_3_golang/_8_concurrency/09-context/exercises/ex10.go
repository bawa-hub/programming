package exercises

import (
	"context"
	"fmt"
	"time"
)

// Exercise 10: Context Performance Testing
func Exercise10() {
	fmt.Println("\nExercise 10: Context Performance Testing")
	fmt.Println("=======================================")
	
	// TODO: Test context performance
	// 1. Measure context creation time
	// 2. Measure context value lookup time
	// 3. Measure context cancellation time
	// 4. Compare different context types
	
	// Test context creation performance
	start := time.Now()
	for i := 0; i < 100000; i++ {
		ctx := context.WithValue(context.Background(), "key", i)
		_ = ctx.Value("key")
	}
	duration := time.Since(start)
	fmt.Printf("  Exercise 10: Context creation time: %v\n", duration)
	
	// Test context cancellation performance
	ctx, cancel := context.WithCancel(context.Background())
	
	// Create many goroutines
	numGoroutines := 1000
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					// Do work
					time.Sleep(1 * time.Millisecond)
				}
			}
		}(i)
	}
	
	// Measure cancellation time
	start = time.Now()
	cancel()
	duration = time.Since(start)
	fmt.Printf("  Exercise 10: Context cancellation time: %v\n", duration)
	
	// Wait for cancellation to propagate
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Exercise 10 completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Context Package Exercises")
	fmt.Println("============================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}