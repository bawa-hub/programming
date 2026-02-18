package exercises

import (
	"fmt"
	"time"
)

// Exercise 5: Select Statement
// Use select to handle multiple channels.
func Exercise5() {
	fmt.Println("\nExercise 5: Select Statement")
	fmt.Println("============================")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	
	// Start goroutines
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch3 <- "Message from channel 3"
	}()
	
	// Use select to receive from any channel
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		case msg3 := <-ch3:
			fmt.Printf("Received from ch3: %s\n", msg3)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}