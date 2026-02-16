package exercises

import (
	"fmt"
	"time"
)

// Exercise 4: Goroutine Communication
// Create 2 goroutines that communicate through a channel.
func Exercise4() {
	fmt.Println("\nExercise 4: Goroutine Communication")
	fmt.Println("===================================")
	
	ch := make(chan string)
	
	// Producer goroutine
	go func() {
		messages := []string{"Hello", "from", "producer", "goroutine"}
		for _, msg := range messages {
			ch <- msg
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()
	
	// Consumer goroutine
	go func() {
		for msg := range ch {
			fmt.Printf("Consumer received: %s\n", msg)
		}
		fmt.Println("Consumer: Channel closed, exiting")
	}()
	
	// Wait for communication to complete
	time.Sleep(1 * time.Second)
}