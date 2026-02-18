package exercises

import "fmt"

// Exercise 1: Basic Channel Operations
// Create a program that sends and receives values through a channel.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Channel Operations")
	fmt.Println("====================================")
	
	ch := make(chan string)
	
	// Send values
	go func() {
		ch <- "Hello"
		ch <- "World"
		ch <- "from"
		ch <- "Go"
		close(ch)
	}()
	
	// Receive values
	for msg := range ch {
		fmt.Printf("Received: %s\n", msg)
	}
}