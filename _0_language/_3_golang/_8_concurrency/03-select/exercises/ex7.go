package exercises

import (
	"fmt"
	"time"
)

// Exercise 7: Error Handling Select
// Add error handling to select operations.
func Exercise7() {
	fmt.Println("\nExercise 7: Error Handling Select")
	fmt.Println("=================================")
	
	ch := make(chan string)
	errCh := make(chan error)
	
	// Send messages and errors
	go func() {
		defer close(ch)
		defer close(errCh)
		
		for i := 1; i <= 5; i++ {
			if i == 3 {
				errCh <- fmt.Errorf("Error at message %d", i)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Select with error handling
	messageCount := 0
	errorCount := 0
	
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				fmt.Printf("Total messages: %d, Total errors: %d\n", messageCount, errorCount)
				return
			}
			messageCount++
			fmt.Printf("✓ Received: %s\n", msg)
		case err, ok := <-errCh:
			if !ok {
				fmt.Println("Error channel closed")
				return
			}
			errorCount++
			fmt.Printf("✗ Error: %v\n", err)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout")
			return
		}
	}
}