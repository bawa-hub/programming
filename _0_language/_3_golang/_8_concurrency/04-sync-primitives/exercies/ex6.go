package exercies

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Exercise 6: Atomic Operations
// Implement a counter using atomic operations.
func Exercise6() {
	fmt.Println("\nExercise 6: Atomic Operations")
	fmt.Println("============================")
	
	var counter int64
	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))
	
	// Compare and swap
	oldValue := atomic.LoadInt64(&counter)
	newValue := oldValue + 100
	if atomic.CompareAndSwapInt64(&counter, oldValue, newValue) {
		fmt.Printf("CAS successful: %d -> %d\n", oldValue, newValue)
	} else {
		fmt.Println("CAS failed")
	}
}