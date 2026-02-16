package exercises

import (
	"fmt"
	"time"
)

// Exercise 1: Basic Goroutines
// Create a program that starts 5 goroutines, each printing a unique number.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Goroutines")
	fmt.Println("============================")
	
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: Hello!\n", id)
		}(i)
	}
	
	// Wait for goroutines to complete
	time.Sleep(100 * time.Millisecond)
}