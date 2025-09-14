package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Non-blocking Select with Default ===")
	
	ch := make(chan string)
	
	// Try to receive from channel (will be empty initially)
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message available, doing other work...")
	}
	
	// Start a goroutine that will send a message later
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello from goroutine!"
	}()
	
	// Check for messages multiple times
	for i := 0; i < 5; i++ {
		select {
		case msg := <-ch:
			fmt.Printf("Received: %s\n", msg)
		default:
			fmt.Printf("Check %d: No message yet, continuing...\n", i+1)
		}
		time.Sleep(300 * time.Millisecond)
	}
	
	fmt.Println("Done checking for messages!")
}
