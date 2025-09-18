package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache with read-write mutex
type SafeCache struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		data: make(map[string]int),
	}
}

func (c *SafeCache) Set(key string, value int) {
	c.mu.Lock()         // Write lock (exclusive)
	defer c.mu.Unlock() // Unlock when done
	c.data[key] = value
	fmt.Printf("Set %s = %d\n", key, value)
}

func (c *SafeCache) Get(key string) (int, bool) {
	c.mu.RLock()         // Read lock (shared)
	defer c.mu.RUnlock() // Unlock when done
	value, exists := c.data[key]
	return value, exists
}

func main() {
	fmt.Println("=== Sync.RWMutex Example ===")
	
	cache := NewSafeCache()
	var wg sync.WaitGroup
	
	// Start multiple readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				key := fmt.Sprintf("key%d", j)
				value, exists := cache.Get(key)
				if exists {
					fmt.Printf("Reader %d: Got %s = %d\n", readerID, key, value)
				} else {
					fmt.Printf("Reader %d: Key %s not found\n", readerID, key)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	// Start a writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			key := fmt.Sprintf("key%d", i)
			cache.Set(key, i*10)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	wg.Wait()
	fmt.Println("All operations completed!")
}
