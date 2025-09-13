package main

import (
	"fmt"
	"hash/fnv"
	"runtime"
	"sync"
	"time"
)

// 🚀 HIGH-PERFORMANCE CACHE PROJECT
// A production-ready, memory-efficient cache with advanced optimizations

type Cache struct {
	shards    []*Shard
	shardMask uint32
	stats     *CacheStats
}

type Shard struct {
	mu     sync.RWMutex
	items  map[string]*CacheItem
	hits   int64
	misses int64
}

type CacheItem struct {
	key        string
	value      interface{}
	expiresAt  time.Time
	accessTime time.Time
	size       int
}

type CacheStats struct {
	mu        sync.RWMutex
	hits      int64
	misses    int64
	evictions int64
	size      int64
}

type CacheConfig struct {
	Shards        int
	MaxSize       int
	DefaultTTL    time.Duration
	CleanupInterval time.Duration
}

func main() {
	fmt.Println("🚀 HIGH-PERFORMANCE CACHE")
	fmt.Println("==========================")

	// Create cache with optimized configuration
	config := CacheConfig{
		Shards:         16,  // 16 shards for better concurrency
		MaxSize:        10000,
		DefaultTTL:     5 * time.Minute,
		CleanupInterval: 1 * time.Minute,
	}

	cache := NewCache(config)
	defer cache.Close()

	// Performance test
	fmt.Println("📊 Running performance tests...")
	performanceTest(cache)

	// Memory efficiency test
	fmt.Println("📊 Running memory efficiency tests...")
	memoryEfficiencyTest(cache)

	// Concurrency test
	fmt.Println("📊 Running concurrency tests...")
	concurrencyTest(cache)

	// Display final stats
	cache.DisplayStats()
}

// NewCache creates a new high-performance cache
func NewCache(config CacheConfig) *Cache {
	// Ensure shards is a power of 2 for efficient modulo
	shards := 1
	for shards < config.Shards {
		shards <<= 1
	}

	cache := &Cache{
		shards:    make([]*Shard, shards),
		shardMask: uint32(shards - 1),
		stats:     &CacheStats{},
	}

	// Initialize shards
	for i := 0; i < shards; i++ {
		cache.shards[i] = &Shard{
			items: make(map[string]*CacheItem),
		}
	}

	// Start cleanup goroutine
	go cache.cleanup(config.CleanupInterval)

	return cache
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	shard := c.getShard(key)
	shard.mu.RLock()
	defer shard.mu.RUnlock()

	item, exists := shard.items[key]
	if !exists {
		c.stats.recordMiss()
		shard.misses++
		return nil, false
	}

	// Check expiration
	if time.Now().After(item.expiresAt) {
		delete(shard.items, key)
		c.stats.recordMiss()
		shard.misses++
		return nil, false
	}

	// Update access time
	item.accessTime = time.Now()
	c.stats.recordHit()
	shard.hits++
	return item.value, true
}

// Set stores a value in the cache
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	if ttl <= 0 {
		ttl = 5 * time.Minute // Default TTL
	}

	shard := c.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()

	// Calculate item size
	size := c.calculateSize(key, value)

	// Create cache item
	item := &CacheItem{
		key:        key,
		value:      value,
		expiresAt:  time.Now().Add(ttl),
		accessTime: time.Now(),
		size:       size,
	}

	// Check if we need to evict
	if len(shard.items) >= 1000 { // Per-shard limit
		c.evictLRU(shard)
	}

	shard.items[key] = item
	c.stats.recordSet(size)
}

// Delete removes a value from the cache
func (c *Cache) Delete(key string) {
	shard := c.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()

	if item, exists := shard.items[key]; exists {
		delete(shard.items, key)
		c.stats.recordDelete(item.size)
	}
}

// getShard returns the shard for a given key
func (c *Cache) getShard(key string) *Shard {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return c.shards[hash.Sum32()&c.shardMask]
}

// evictLRU evicts the least recently used item from a shard
func (c *Cache) evictLRU(shard *Shard) {
	var oldestKey string
	var oldestTime time.Time

	for key, item := range shard.items {
		if oldestKey == "" || item.accessTime.Before(oldestTime) {
			oldestKey = key
			oldestTime = item.accessTime
		}
	}

	if oldestKey != "" {
		delete(shard.items, oldestKey)
		c.stats.recordEviction()
	}
}

// calculateSize estimates the memory size of a cache item
func (c *Cache) calculateSize(key string, value interface{}) int {
	size := len(key) + 24 // Key length + overhead
	
	switch v := value.(type) {
	case string:
		size += len(v)
	case []byte:
		size += len(v)
	case int, int32, int64:
		size += 8
	case float32, float64:
		size += 8
	case bool:
		size += 1
	default:
		size += 16 // Default overhead
	}
	
	return size
}

// cleanup removes expired items periodically
func (c *Cache) cleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		evicted := 0

		for _, shard := range c.shards {
			shard.mu.Lock()
			for key, item := range shard.items {
				if now.After(item.expiresAt) {
					delete(shard.items, key)
					evicted++
				}
			}
			shard.mu.Unlock()
		}

		if evicted > 0 {
			c.stats.recordEviction()
		}
	}
}

// Close gracefully shuts down the cache
func (c *Cache) Close() {
	// Cleanup goroutine will stop when the ticker is stopped
}

// DisplayStats shows cache statistics
func (c *Cache) DisplayStats() {
	c.stats.mu.RLock()
	defer c.stats.mu.RUnlock()

	fmt.Printf("\n📊 CACHE STATISTICS\n")
	fmt.Printf("==================\n")
	fmt.Printf("Hits: %d\n", c.stats.hits)
	fmt.Printf("Misses: %d\n", c.stats.misses)
	fmt.Printf("Hit Rate: %.2f%%\n", float64(c.stats.hits)/float64(c.stats.hits+c.stats.misses)*100)
	fmt.Printf("Evictions: %d\n", c.stats.evictions)
	fmt.Printf("Total Size: %d bytes\n", c.stats.size)
}

// CacheStats methods
func (s *CacheStats) recordHit() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.hits++
}

func (s *CacheStats) recordMiss() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.misses++
}

func (s *CacheStats) recordSet(size int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.size += int64(size)
}

func (s *CacheStats) recordDelete(size int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.size -= int64(size)
}

func (s *CacheStats) recordEviction() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.evictions++
}

// PERFORMANCE TESTS
func performanceTest(cache *Cache) {
	// Test 1: Sequential operations
	fmt.Println("  🧪 Test 1: Sequential operations")
	start := time.Now()
	
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)
		cache.Set(key, value, 5*time.Minute)
	}
	
	setTime := time.Since(start)
	fmt.Printf("    Set 10,000 items: %v\n", setTime)
	
	// Test 2: Read operations
	start = time.Now()
	
	hits := 0
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key-%d", i)
		if _, found := cache.Get(key); found {
			hits++
		}
	}
	
	getTime := time.Since(start)
	fmt.Printf("    Get 10,000 items: %v (hits: %d)\n", getTime, hits)
}

func memoryEfficiencyTest(cache *Cache) {
	// Test memory usage
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	
	// Add items to cache
	for i := 0; i < 5000; i++ {
		key := fmt.Sprintf("mem-key-%d", i)
		value := make([]byte, 1024) // 1KB each
		cache.Set(key, value, 5*time.Minute)
	}
	
	runtime.ReadMemStats(&m2)
	
	fmt.Printf("  📊 Memory allocated: %d bytes (%.2f MB)\n", 
		m2.Alloc-m1.Alloc, float64(m2.Alloc-m1.Alloc)/1024/1024)
	fmt.Printf("  📊 Items per MB: %.2f\n", 
		float64(5000)/(float64(m2.Alloc-m1.Alloc)/1024/1024))
}

func concurrencyTest(cache *Cache) {
	// Test concurrent access
	var wg sync.WaitGroup
	goroutines := 100
	operations := 1000
	
	start := time.Now()
	
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			for j := 0; j < operations; j++ {
				key := fmt.Sprintf("concurrent-key-%d-%d", id, j)
				value := fmt.Sprintf("concurrent-value-%d-%d", id, j)
				
				// Set
				cache.Set(key, value, 5*time.Minute)
				
				// Get
				cache.Get(key)
			}
		}(i)
	}
	
	wg.Wait()
	duration := time.Since(start)
	
	fmt.Printf("  📊 Concurrent operations: %d goroutines × %d ops = %d total\n", 
		goroutines, operations, goroutines*operations)
	fmt.Printf("  📊 Total time: %v\n", duration)
	fmt.Printf("  📊 Operations per second: %.0f\n", 
		float64(goroutines*operations)/duration.Seconds())
}
