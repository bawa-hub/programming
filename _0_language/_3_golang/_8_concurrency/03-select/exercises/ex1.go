package exercises

import (
	"fmt"
	"time"
)

// Exercise 1: Basic Select
// Create a select statement that handles two channels.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Select")
	fmt.Println("========================")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Send messages
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Hello from channel 1"
	}()
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "Hello from channel 2"
	}()
	
	// Select from both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		}
	}
}