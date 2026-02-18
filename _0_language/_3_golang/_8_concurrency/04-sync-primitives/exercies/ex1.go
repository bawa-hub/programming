package exercies

import (
	"fmt"
	"sync"
)

// Exercise 1: Basic Mutex
// Create a counter protected by a mutex.
func Exercise1() {
	fmt.Println("Exercise 1: Basic Mutex")
	fmt.Println("=======================")
	
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}