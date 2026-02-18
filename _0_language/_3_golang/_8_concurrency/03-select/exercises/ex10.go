package exercises

import (
	"context"
	"fmt"
	"time"
)

// Exercise 10: Select with Context
func Exercise10() {
	fmt.Println("\nExercise 10: Select with Context")
	fmt.Println("===============================")
	
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	
	// Send a message after delay
	go func() {
		time.Sleep(300 * time.Millisecond)
		select {
		case ch <- "Context message":
			fmt.Println("Message sent successfully")
		case <-ctx.Done():
			fmt.Println("Context cancelled before message could be sent")
		}
		close(ch)
	}()
	
	// Select with context
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-ctx.Done():
		fmt.Printf("Context cancelled: %v\n", ctx.Err())
	}
}