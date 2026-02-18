package exercises

import (
	"fmt"
	"time"
)

// Exercise 8: Dynamic Select
// Create a select statement that handles a variable number of channels.
func Exercise8() {
	fmt.Println("\nExercise 8: Dynamic Select")
	fmt.Println("==========================")
	
	// Create multiple channels
	channels := make([]chan string, 3)
	for i := range channels {
		channels[i] = make(chan string)
	}
	
	// Send messages to different channels
	for i, ch := range channels {
		go func(id int, ch chan<- string) {
			defer close(ch)
			for j := 1; j <= 2; j++ {
				ch <- fmt.Sprintf("Channel %d, Message %d", id, j)
				time.Sleep(time.Duration(100+id*50) * time.Millisecond)
			}
		}(i, ch)
	}
	
	// Dynamic select (simulated with fixed number of channels)
	fmt.Println("Dynamic select results:")
	for i := 0; i < 6; i++ {
		select {
		case msg := <-channels[0]:
			fmt.Printf("  From channel 0: %s\n", msg)
		case msg := <-channels[1]:
			fmt.Printf("  From channel 1: %s\n", msg)
		case msg := <-channels[2]:
			fmt.Printf("  From channel 2: %s\n", msg)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("  Timeout")
			return
		}
	}
}