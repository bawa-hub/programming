package main

import (
	"fmt"
	"sync"
)

// Counter with mutex protection
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()         // Lock before accessing shared data
	defer c.mu.Unlock() // Unlock when function exits
	c.value++
}

func (c *SafeCounter) GetValue() int {
	c.mu.Lock()         // Lock before reading
	defer c.mu.Unlock() // Unlock when function exits
	return c.value
}

func main() {
	fmt.Println("=== Sync.Mutex Example ===")
	
	counter := &SafeCounter{}
	var wg sync.WaitGroup
	
	// Start 10 goroutines that increment the counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
			fmt.Printf("Goroutine %d: Finished incrementing\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.GetValue())
	fmt.Println("Expected: 10000 (10 goroutines × 1000 increments each)")
}
