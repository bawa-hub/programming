package exercises

import (
	"fmt"
	"runtime"
	"time"
)

// Exercise 7: Goroutine Monitoring
func Exercise7() {
	fmt.Println("\nExercise 7: Goroutine Monitoring")
	fmt.Println("===============================")
	
	// Start some goroutines
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 3; j++ {
				fmt.Printf("Goroutine %d: iteration %d\n", id, j)
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}
	
	// Monitor goroutines
	for i := 0; i < 3; i++ {
		fmt.Printf("Active goroutines: %d\n", runtime.NumGoroutine())
		time.Sleep(300 * time.Millisecond)
	}
}