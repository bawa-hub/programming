package main

import (
	"fmt"
	"time"
)

// TODO: Create a function that sends messages to a channel every 500ms
// TODO: Create a function that sends messages to another channel every 300ms
// TODO: Use select to receive from both channels
// TODO: Add a timeout case that triggers after 3 seconds
// TODO: Print which channel sent the message or if timeout occurred

func main() {
	fmt.Println("=== Select Exercise ===")
	
	// Your code goes here:
	// 1. Create two channels
	// 2. Start goroutines that send messages at different intervals
	// 3. Use select to receive from both channels
	// 4. Add timeout case
	// 5. Count and print messages received
	
	fmt.Println("Exercise completed!")
}
