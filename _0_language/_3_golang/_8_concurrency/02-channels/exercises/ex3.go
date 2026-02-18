package exercises

import (
	"fmt"
	"time"
)

// Exercise 3: Channel Direction
// Create functions that use send-only and receive-only channels.
func Exercise3() {
	fmt.Println("\nExercise 3: Channel Direction")
	fmt.Println("=============================")
	
	ch := make(chan int)
	
	// Send-only function
	go sendOnly(ch)
	
	// Receive-only function
	receiveOnly(ch)
}

func sendOnly(ch chan<- int) {
	fmt.Println("Send-only function: Sending values...")
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func receiveOnly(ch <-chan int) {
	fmt.Println("Receive-only function: Receiving values...")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println("Channel closed")
}
