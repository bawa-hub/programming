package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *Singleton
)

// Singleton pattern using sync.Once
type Singleton struct {
	ID string
}

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating singleton instance...")
		instance = &Singleton{ID: "unique-instance"}
		time.Sleep(100 * time.Millisecond) // Simulate expensive initialization
		fmt.Println("Singleton instance created!")
	})
	return instance
}

func main() {
	fmt.Println("=== Sync.Once Example ===")
	
	var wg sync.WaitGroup
	
	// Start 10 goroutines that all try to get the singleton
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: Requesting singleton...\n", id)
			singleton := GetInstance()
			fmt.Printf("Goroutine %d: Got singleton with ID: %s\n", id, singleton.ID)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All goroutines completed!")
}
