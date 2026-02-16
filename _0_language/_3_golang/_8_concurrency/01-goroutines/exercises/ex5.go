package exercises

import (
	"fmt"
	"time"
)

// Exercise 5: Goroutine Lifecycle
// Implement a goroutine that can be started, paused, and stopped.
func Exercise5() {
	fmt.Println("\nExercise 5: Goroutine Lifecycle")
	fmt.Println("===============================")
	
	start := make(chan bool)
	pause := make(chan bool)
	stop := make(chan bool)
	
	// Controllable goroutine
	go func() {
		running := false
		for {
			select {
			case <-start:
				running = true
				fmt.Println("Goroutine: Started")
			case <-pause:
				if running {
					running = false
					fmt.Println("Goroutine: Paused")
				}
			case <-stop:
				fmt.Println("Goroutine: Stopped")
				return
			default:
				if running {
					fmt.Println("Goroutine: Working...")
					time.Sleep(200 * time.Millisecond)
				}
			}
		}
	}()
	
	// Control the goroutine
	start <- true
	time.Sleep(500 * time.Millisecond)
	
	pause <- true
	time.Sleep(500 * time.Millisecond)
	
	start <- true
	time.Sleep(500 * time.Millisecond)
	
	stop <- true
	time.Sleep(100 * time.Millisecond)
}