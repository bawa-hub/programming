package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Goroutine Pitfalls Example ===")
	
	// Pitfall 1: Loop variable capture
	fmt.Println("\n1. Loop variable capture (WRONG):")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("Wrong: i = %d\n", i) // This will print 3, 3, 3 , it is non-deterministic and depends on machine
		}()
	}
	time.Sleep(100 * time.Millisecond)
	
	// Pitfall 1: Solution - pass variable as parameter
	fmt.Println("\n2. Loop variable capture (CORRECT):")
	for i := 0; i < 3; i++ {
		go func(val int) {
			fmt.Printf("Correct: val = %d\n", val)
		}(i) // Pass i as parameter
	}
	time.Sleep(100 * time.Millisecond)
	
	// Pitfall 2: Main goroutine exits too early
	fmt.Println("\n3. Main goroutine exits too early (WRONG):")
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("This might not print!")
	}()
	// No wait here - main might exit before goroutine completes
	
	// Pitfall 2: Solution - wait for goroutines using time.Sleep
	fmt.Println("\n4. Waiting for goroutines (CORRECT):")
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("This will definitely print!")
	}()
	time.Sleep(300 * time.Millisecond) // Wait longer than the goroutine
	
	fmt.Println("\nAll examples completed!")
}
