package main

import (
	"fmt"
	"sync"
	"time"
)

// ThreadSafeCache represents a thread-safe cache
type ThreadSafeCache struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

// NewThreadSafeCache creates a new thread-safe cache
func NewThreadSafeCache() *ThreadSafeCache {
	return &ThreadSafeCache{
		data: make(map[string]interface{}),
	}
}

// Set stores a value in the cache
func (c *ThreadSafeCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
	fmt.Printf("Set: %s = %v\n", key, value)
}

// Get retrieves a value from the cache
func (c *ThreadSafeCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, exists := c.data[key]
	if exists {
		fmt.Printf("Get: %s = %v\n", key, value)
	} else {
		fmt.Printf("Get: %s not found\n", key)
	}
	return value, exists
}

// Delete removes a value from the cache
func (c *ThreadSafeCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
	fmt.Printf("Delete: %s\n", key)
}

// Size returns the number of items in the cache
func (c *ThreadSafeCache) Size() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.data)
}

// Clear removes all items from the cache
func (c *ThreadSafeCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = make(map[string]interface{})
	fmt.Println("Cache cleared")
}

// GetAll returns a copy of all cache data
func (c *ThreadSafeCache) GetAll() map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	result := make(map[string]interface{})
	for k, v := range c.data {
		result[k] = v
	}
	return result
}

// DemonstrateCache demonstrates the thread-safe cache
func DemonstrateCache() {
	fmt.Println("=== Thread-Safe Cache Demonstration ===")
	
	cache := NewThreadSafeCache()
	var wg sync.WaitGroup

	// Multiple writers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(writerID int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key := fmt.Sprintf("writer%d_key%d", writerID, j)
				value := fmt.Sprintf("value_%d_%d", writerID, j)
				cache.Set(key, value)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Multiple readers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("writer%d_key%d", j%3, j%5)
				cache.Get(key)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	
	fmt.Printf("Cache size: %d\n", cache.Size())
	fmt.Println("All cache data:")
	for key, value := range cache.GetAll() {
		fmt.Printf("  %s: %v\n", key, value)
	}
}

// CacheWithTTL represents a cache with time-to-live
type CacheWithTTL struct {
	data      map[string]cacheItem
	mutex     sync.RWMutex
	cleanupCh chan string
}

type cacheItem struct {
	value     interface{}
	expiresAt time.Time
}

// NewCacheWithTTL creates a new cache with TTL
func NewCacheWithTTL() *CacheWithTTL {
	cache := &CacheWithTTL{
		data:      make(map[string]cacheItem),
		cleanupCh: make(chan string, 100),
	}
	
	// Start cleanup goroutine
	go cache.cleanup()
	
	return cache
}

// SetWithTTL stores a value with TTL
func (c *CacheWithTTL) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.data[key] = cacheItem{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
	
	fmt.Printf("Set with TTL: %s = %v (expires in %v)\n", key, value, ttl)
}

// Get retrieves a value if not expired
func (c *CacheWithTTL) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	item, exists := c.data[key]
	c.mutex.RUnlock()
	
	if !exists {
		fmt.Printf("Get: %s not found\n", key)
		return nil, false
	}
	
	if time.Now().After(item.expiresAt) {
		fmt.Printf("Get: %s expired\n", key)
		c.cleanupCh <- key
		return nil, false
	}
	
	fmt.Printf("Get: %s = %v\n", key, item.value)
	return item.value, true
}

// cleanup removes expired items
func (c *CacheWithTTL) cleanup() {
	for key := range c.cleanupCh {
		c.mutex.Lock()
		if item, exists := c.data[key]; exists && time.Now().After(item.expiresAt) {
			delete(c.data, key)
			fmt.Printf("Cleaned up expired item: %s\n", key)
		}
		c.mutex.Unlock()
	}
}

// DemonstrateCacheWithTTL demonstrates the TTL cache
func DemonstrateCacheWithTTL() {
	fmt.Println("\n=== Cache with TTL Demonstration ===")
	
	cache := NewCacheWithTTL()
	
	// Set items with different TTLs
	cache.SetWithTTL("short", "expires quickly", 500*time.Millisecond)
	cache.SetWithTTL("long", "expires slowly", 2*time.Second)
	
	// Try to get items immediately
	cache.Get("short")
	cache.Get("long")
	
	// Wait for short item to expire
	time.Sleep(600 * time.Millisecond)
	cache.Get("short") // Should be expired
	cache.Get("long")  // Should still be valid
	
	// Wait for long item to expire
	time.Sleep(2 * time.Second)
	cache.Get("long") // Should be expired
}
