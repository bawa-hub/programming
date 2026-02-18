package exercies

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 5: Cond
// Use condition variables to coordinate goroutines.
func Exercise5() {
	fmt.Println("\nExercise 5: Cond")
	fmt.Println("================")
	
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var ready bool
	var wg sync.WaitGroup
	
	// Waiters
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			for !ready {
				fmt.Printf("Waiter %d: waiting for condition\n", id)
				cond.Wait()
			}
			fmt.Printf("Waiter %d: condition met!\n", id)
			mu.Unlock()
		}(i)
	}
	
	// Signaler
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		mu.Lock()
		ready = true
		fmt.Println("Signaler: condition is ready, broadcasting...")
		cond.Broadcast()
		mu.Unlock()
	}()
	
	wg.Wait()
}