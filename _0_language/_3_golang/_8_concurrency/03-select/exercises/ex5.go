package exercises

import (
	"fmt"
	"time"
)

// Exercise 5: Multiplexing Select
// Create a fan-in pattern using select.
func Exercise5() {
	fmt.Println("\nExercise 5: Multiplexing Select")
	fmt.Println("===============================")
	
	input1 := make(chan string)
	input2 := make(chan string)
	output := make(chan string)
	
	// Fan-in goroutine
	go func() {
		defer close(output)
		for {
			select {
			case msg, ok := <-input1:
				if !ok {
					input1 = nil // Disable this case
				} else {
					output <- "Input1: " + msg
				}
			case msg, ok := <-input2:
				if !ok {
					input2 = nil // Disable this case
				} else {
					output <- "Input2: " + msg
				}
			}
			
			// Exit when both channels are closed
			if input1 == nil && input2 == nil {
				return
			}
		}
	}()
	
	// Send messages to inputs
	go func() {
		defer close(input1)
		for i := 1; i <= 3; i++ {
			input1 <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(input2)
		for i := 1; i <= 3; i++ {
			input2 <- fmt.Sprintf("Data %d", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()
	
	// Collect multiplexed output
	fmt.Println("Multiplexed output:")
	for msg := range output {
		fmt.Printf("  %s\n", msg)
	}
}