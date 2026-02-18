package exercies

import (
	"fmt"
	"sync"
	"time"
)

// Cache type for Exercise 2
type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// Exercise 2: RWMutex
// Implement a thread-safe cache using RWMutex.
func Exercise2() {
	fmt.Println("\nExercise 2: RWMutex")
	fmt.Println("===================")
	
	cache := &Cache{
		data: make(map[string]int),
	}
	
	// Set values
	cache.Set("key1", 1)
	cache.Set("key2", 2)
	cache.Set("key3", 3)
	
	// Multiple readers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				value := cache.Get("key1")
				fmt.Printf("Reader %d: key1 = %d\n", id, value)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	
	// One writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			cache.Set("key1", cache.Get("key1")+1)
			fmt.Printf("Writer: updated key1 to %d\n", cache.Get("key1"))
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	wg.Wait()
	fmt.Printf("Final value: %d\n", cache.Get("key1"))
}
