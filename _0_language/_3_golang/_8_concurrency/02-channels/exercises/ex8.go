package exercises

import (
	"fmt"
	"time"
)

// Exercise 8: Channel Timeout
// Add timeout handling to channel operations.
func Exercise8() {
	fmt.Println("\nExercise 8: Channel Timeout")
	fmt.Println("===========================")
	
	ch := make(chan string)
	
	// Send a message after delay
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "Delayed message"
	}()
	
	// Wait for message with timeout
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout! No message received")
	}
}