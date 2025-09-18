package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Select with Timeout ===")
	
	ch := make(chan string)
	
	// Example 1: Timeout before message arrives
	fmt.Println("\n1. Timeout before message:")
	go func() {
		time.Sleep(2 * time.Second) // Message arrives after 2 seconds
		ch <- "This message arrives too late!"
	}()
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second): // Timeout after 1 second
		fmt.Println("Timeout! No message received in time.")
	}
	
	// Example 2: Message arrives before timeout
	fmt.Println("\n2. Message arrives before timeout:")
	go func() {
		time.Sleep(500 * time.Millisecond) // Message arrives after 0.5 seconds
		ch <- "This message arrives on time!"
	}()
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second): // Timeout after 1 second
		fmt.Println("Timeout! No message received in time.")
	}
	
	// Example 3: Multiple timeouts
	fmt.Println("\n3. Multiple timeouts:")
	timeout1 := time.After(1 * time.Second)
	timeout2 := time.After(2 * time.Second)
	
	select {
	case <-timeout1:
		fmt.Println("First timeout reached!")
	case <-timeout2:
		fmt.Println("Second timeout reached!")
	}
	
	fmt.Println("\nAll timeout examples completed!")
}
