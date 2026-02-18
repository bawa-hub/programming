package exercises

import (
	"fmt"
	"time"
)

// Exercise 9: Select with Ticker
func Exercise9() {
	fmt.Println("\nExercise 9: Select with Ticker")
	fmt.Println("=============================")
	
	ch := make(chan string)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	
	// Send messages
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()
	
	// Select with ticker
	tickCount := 0
	messageCount := 0
	
	for i := 0; i < 10; i++ {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				fmt.Printf("Total ticks: %d, Total messages: %d\n", tickCount, messageCount)
				return
			}
			messageCount++
			fmt.Printf("📨 Received: %s\n", msg)
		case <-ticker.C:
			tickCount++
			fmt.Printf("⏰ Tick %d\n", tickCount)
		}
	}
}