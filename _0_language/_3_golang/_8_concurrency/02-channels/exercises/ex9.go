package exercises

import (
	"fmt"
	"time"
)

// Exercise 9: Channel with Quit Signal
func Exercise9() {
	fmt.Println("\nExercise 9: Channel with Quit Signal")
	fmt.Println("====================================")
	
	workCh := make(chan int)
	quitCh := make(chan bool)
	
	// Worker goroutine
	go func() {
		for {
			select {
			case work := <-workCh:
				fmt.Printf("Processing work: %d\n", work)
				time.Sleep(100 * time.Millisecond)
			case <-quitCh:
				fmt.Println("Worker: Received quit signal, exiting")
				return
			}
		}
	}()
	
	// Send some work
	for i := 1; i <= 3; i++ {
		workCh <- i
	}
	
	// Send quit signal
	quitCh <- true
	time.Sleep(100 * time.Millisecond)
}