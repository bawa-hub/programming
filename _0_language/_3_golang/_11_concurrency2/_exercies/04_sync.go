package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Create a thread-safe counter that can be incremented and decremented
// TODO: Use sync.RWMutex for efficient read operations
// TODO: Create multiple goroutines that increment and decrement the counter
// TODO: Use sync.WaitGroup to wait for all goroutines to complete
// TODO: Print the final counter value

func main() {
	fmt.Println("=== Sync Package Exercise ===")
	
	// Your code goes here:
	// 1. Create a SafeCounter struct with RWMutex
	// 2. Implement Increment(), Decrement(), and GetValue() methods
	// 3. Start multiple goroutines that increment/decrement
	// 4. Use WaitGroup to wait for completion
	// 5. Print final value
	
	fmt.Println("Exercise completed!")
}
