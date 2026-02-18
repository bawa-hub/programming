package exercies

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 4: Once
// Implement a singleton pattern using Once.
func Exercise4() {
	fmt.Println("\nExercise 4: Once")
	fmt.Println("================")
	
	var once sync.Once
	var instance *g.Singleton
	var wg sync.WaitGroup
	
	// Multiple goroutines trying to create singleton
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				instance = &Singleton{ID: id, Created: time.Now()}
				fmt.Printf("Goroutine %d created singleton\n", id)
			})
			fmt.Printf("Goroutine %d got instance: %+v\n", id, instance)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final instance: %+v\n", instance)
}

type Singleton struct {
	ID      int
	Created time.Time
}