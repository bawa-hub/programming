package main

import (
	"fmt"
	"time"
)

// Function that only sends data (send-only channel)
func sender(ch chan<- string) {
	for i := 1; i <= 3; i++ {
		message := fmt.Sprintf("Message %d", i)
		ch <- message
		fmt.Printf("Sent: %s\n", message)
		time.Sleep(200 * time.Millisecond)
	}
	close(ch) // Close the channel when done sending
}

// Function that only receives data (receive-only channel)
func receiver(ch <-chan string) {
	for message := range ch {
		fmt.Printf("Received: %s\n", message)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== Channel Directions Example ===")
	
	// Create a bidirectional channel
	ch := make(chan string)
	
	// Start sender and receiver goroutines
	go sender(ch)   // ch becomes send-only in sender function
	go receiver(ch) // ch becomes receive-only in receiver function
	
	// Wait for both to complete
	time.Sleep(2 * time.Second)
	
	fmt.Println("\nAll messages sent and received!")
}

// The send–receive ordering is guaranteed by Go’s channel semantics.
// The print ordering is not guaranteed — it depends on scheduling + stdout buffering.