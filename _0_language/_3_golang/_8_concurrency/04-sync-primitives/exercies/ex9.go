package exercies

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 9: Deadlock Prevention
func Exercise9() {
	fmt.Println("\nExercise 9: Deadlock Prevention")
	fmt.Println("===============================")
	
	var mu1, mu2 sync.Mutex
	
	// Function that acquires locks in order
	acquireLocks := func(id int, mu1, mu2 *sync.Mutex) {
		mu1.Lock()
		fmt.Printf("Goroutine %d: acquired lock1\n", id)
		time.Sleep(10 * time.Millisecond)
		
		mu2.Lock()
		fmt.Printf("Goroutine %d: acquired lock2\n", id)
		time.Sleep(10 * time.Millisecond)
		
		mu2.Unlock()
		fmt.Printf("Goroutine %d: released lock2\n", id)
		mu1.Unlock()
		fmt.Printf("Goroutine %d: released lock1\n", id)
	}
	
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			acquireLocks(id, &mu1, &mu2)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All goroutines completed without deadlock")
}