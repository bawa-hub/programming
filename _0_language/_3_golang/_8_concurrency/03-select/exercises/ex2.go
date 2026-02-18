package exercises

import "fmt"

// Exercise 2: Non-blocking Select
// Implement non-blocking send and receive operations.
func Exercise2() {
	fmt.Println("\nExercise 2: Non-blocking Select")
	fmt.Println("===============================")
	
	ch := make(chan string, 2)
	
	// Non-blocking send
	fmt.Println("Testing non-blocking send:")
	select {
	case ch <- "Message 1":
		fmt.Println("✓ Message 1 sent")
	default:
		fmt.Println("✗ Channel is full")
	}
	
	select {
	case ch <- "Message 2":
		fmt.Println("✓ Message 2 sent")
	default:
		fmt.Println("✗ Channel is full")
	}
	
	select {
	case ch <- "Message 3":
		fmt.Println("✓ Message 3 sent")
	default:
		fmt.Println("✗ Channel is full")
	}
	
	// Non-blocking receive
	fmt.Println("\nTesting non-blocking receive:")
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch:
			fmt.Printf("✓ Received: %s\n", msg)
		default:
			fmt.Println("✗ No message available")
		}
	}
}