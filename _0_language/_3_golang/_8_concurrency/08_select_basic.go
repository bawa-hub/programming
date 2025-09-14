package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Basic Select Statement ===")
	
	// Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Start goroutines that send to different channels
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	// Select will pick whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		}
	}
	
	fmt.Println("All messages received!")
}
