package exercies

import (
	"fmt"
	"sync"
)

// Exercise 7: Concurrent Map
// Use sync.Map for thread-safe map operations.
func Exercise7() {
	fmt.Println("\nExercise 7: Concurrent Map")
	fmt.Println("=========================")
	
	var m sync.Map
	var wg sync.WaitGroup
	
	// Store values
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			m.Store(key, value)
			fmt.Printf("Stored: %s = %s\n", key, value)
		}(i)
	}
	
	// Load values
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			if value, ok := m.Load(key); ok {
				fmt.Printf("Loaded: %s = %s\n", key, value)
			} else {
				fmt.Printf("Key %s not found\n", key)
			}
		}(i)
	}
	
	wg.Wait()
	
	// Range over map
	fmt.Println("All key-value pairs:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %s = %s\n", key, value)
		return true
	})
}