package main

import (
	"fmt"
	"time"
)

// Function that will run in a goroutine
func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello %s! (iteration %d)\n", name, i+1)
		time.Sleep(100 * time.Millisecond) // Simulate some work
	}
}

func main() {
	fmt.Println("=== Goroutines Basic Example ===")
	
	// Without goroutines (sequential execution)
	fmt.Println("\n1. Sequential execution:")
	sayHello("Alice")
	sayHello("Bob")
	
	// With goroutines (concurrent execution)
	fmt.Println("\n2. Concurrent execution with goroutines:")
	go sayHello("Charlie")  // This runs concurrently
	go sayHello("David")    // This also runs concurrently
	
	// Wait a bit to see the output
	time.Sleep(500 * time.Millisecond)
	
	fmt.Println("\n3. Anonymous function goroutines:")
	go func(name string) {
		for i := 0; i < 2; i++ {
			fmt.Printf("Anonymous: Hello %s! (iteration %d)\n", name, i+1)
			time.Sleep(150 * time.Millisecond)
		}
	}("Eve")
	
	// Wait for goroutines to complete
	time.Sleep(1 * time.Second)
	fmt.Println("\nMain function ending...")
}
