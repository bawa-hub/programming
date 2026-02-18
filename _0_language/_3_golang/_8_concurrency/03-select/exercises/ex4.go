package exercises

import (
	"fmt"
	"time"
)

// Exercise 4: Priority Select
// Implement priority handling for multiple channels.
func Exercise4() {
	fmt.Println("\nExercise 4: Priority Select")
	fmt.Println("===========================")
	
	highPriority := make(chan string, 5)
	normalPriority := make(chan string, 5)
	lowPriority := make(chan string, 5)
	
	// Send messages to different priority channels
	go func() {
		time.Sleep(50 * time.Millisecond)
		normalPriority <- "Normal message 1"
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		lowPriority <- "Low message 1"
	}()
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		highPriority <- "High message 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		normalPriority <- "Normal message 2"
	}()
	
	// Handle messages with priority
	for i := 0; i < 4; i++ {
		select {
		case msg := <-highPriority:
			fmt.Printf("🔥 HIGH PRIORITY: %s\n", msg)
		case msg := <-normalPriority:
			fmt.Printf("📝 Normal priority: %s\n", msg)
		case msg := <-lowPriority:
			fmt.Printf("📄 Low priority: %s\n", msg)
		}
	}
}