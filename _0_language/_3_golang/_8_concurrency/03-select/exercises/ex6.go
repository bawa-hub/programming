package exercises

import (
	"fmt"
	"time"
)

// Exercise 6: Loop Select
// Implement a select statement in a loop with proper exit conditions.
func Exercise6() {
	fmt.Println("\nExercise 6: Loop Select")
	fmt.Println("=======================")
	
	ch := make(chan string)
	quit := make(chan bool)
	
	// Send messages
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Send quit signal
	go func() {
		time.Sleep(600 * time.Millisecond)
		quit <- true
	}()
	
	// Loop with select
	messageCount := 0
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed, exiting loop")
				return
			}
			messageCount++
			fmt.Printf("Received: %s (count: %d)\n", msg, messageCount)
		case <-quit:
			fmt.Println("Quit signal received, exiting loop")
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Println("Timeout in loop")
		}
	}
}