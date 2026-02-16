package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 8: Goroutine with Error Handling
func Exercise8() {
	fmt.Println("\nExercise 8: Goroutine with Error Handling")
	fmt.Println("=========================================")
	
	var wg sync.WaitGroup
	errors := make(chan error, 3)
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					errors <- fmt.Errorf("goroutine %d panicked: %v", id, r)
				}
			}()
			
			if id == 1 {
				panic("Simulated panic in goroutine 1")
			}
			
			fmt.Printf("Goroutine %d: Working normally\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	
	wg.Wait()
	close(errors)
	
	// Check for errors
	for err := range errors {
		fmt.Printf("Error: %v\n", err)
	}
}