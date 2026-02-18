package exercises

import (
	"fmt"
	"time"
)

// Exercise 3: Timeout Select
// Add timeout handling to channel operations.
func Exercise3() {
	fmt.Println("\nExercise 3: Timeout Select")
	fmt.Println("==========================")
	
	ch := make(chan string)
	
	// Test with timeout
	fmt.Println("Testing timeout (no message sent):")
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout! No message received")
	}
	
	// Test with message sent after timeout
	fmt.Println("\nTesting timeout (message sent after timeout):")
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch <- "Late message"
	}()
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout! No message received")
	}
}