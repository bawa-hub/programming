package exercises

import "fmt"
import "sync"
import "time"

// Exercise 2: Goroutine Synchronization
// Use WaitGroup to wait for 3 goroutines to complete.
func Exercise2() {
	fmt.Println("\nExercise 2: Goroutine Synchronization")
	fmt.Println("=====================================")
	
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d: Starting work\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Worker %d: Work completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All workers completed!")
}
