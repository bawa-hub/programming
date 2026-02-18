package exercises

import (
	"fmt"
	"sync"
)

// Exercise 7: Fan-Out/Fan-In
// Implement fan-out and fan-in patterns.
func Exercise7() {
	fmt.Println("\nExercise 7: Fan-Out/Fan-In")
	fmt.Println("==========================")
	
	// Create input channel
	input := make(chan int)
	
	// Create output channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	// Fan-out goroutine
	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()
		
		for n := range input {
			ch1 <- n
			ch2 <- n
			ch3 <- n
		}
	}()
	
	// Send input
	go func() {
		defer close(input)
		for i := 1; i <= 3; i++ {
			input <- i
		}
	}()
	
	// Fan-in goroutine
	merged := fanInExercise(ch1, ch2, ch3)
	
	// Collect results
	fmt.Println("Fan-out/Fan-in results:")
	for value := range merged {
		fmt.Printf("Received: %d\n", value)
	}
}

func fanInExercise(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		var wg sync.WaitGroup
		
		for _, input := range inputs {
			wg.Add(1)
			go func(ch <-chan int) {
				defer wg.Done()
				for value := range ch {
					output <- value
				}
			}(input)
		}
		
		wg.Wait()
	}()
	
	return output
}