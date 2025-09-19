package main

import (
	"fmt"
	"time"

	"lru_cache/models"
	"lru_cache/services"
)

func main() {
	fmt.Println("=== LRU CACHE SYSTEM DEMONSTRATION ===\n")

	// Create cache service
	cacheService := services.NewCacheService()

	// Create different caches
	userCache, _ := cacheService.CreateCache("users", 3)
	sessionCache, _ := cacheService.CreateCache("sessions", 5)
	configCache, _ := cacheService.CreateCache("config", 2)

	fmt.Println("1. CACHE CREATION:")
	fmt.Printf("Created user cache with capacity: %d\n", userCache.Capacity())
	fmt.Printf("Created session cache with capacity: %d\n", sessionCache.Capacity())
	fmt.Printf("Created config cache with capacity: %d\n", configCache.Capacity())

	// Basic operations
	fmt.Println()
	fmt.Println("2. BASIC CACHE OPERATIONS:")
	
	// Put operations
	userCache.Put("user1", "Alice Johnson")
	userCache.Put("user2", "Bob Smith")
	userCache.Put("user3", "Charlie Brown")
	
	fmt.Printf("Added 3 users to cache. Current size: %d\n", userCache.Size())
	
	// Get operations
	if value, exists := userCache.Get("user1"); exists {
		fmt.Printf("Retrieved user1: %s\n", value)
	}
	
	if value, exists := userCache.Get("user2"); exists {
		fmt.Printf("Retrieved user2: %s\n", value)
	}

	// Test LRU eviction
	fmt.Println()
	fmt.Println("3. LRU EVICTION TEST:")
	fmt.Printf("Cache before adding 4th user: %v\n", userCache.GetLRUOrder())
	
	userCache.Put("user4", "Diana Prince") // This should evict user1 (least recently used)
	fmt.Printf("Added user4. Current size: %d\n", userCache.Size())
	fmt.Printf("Cache after eviction: %v\n", userCache.GetLRUOrder())
	
	// Check if user1 was evicted
	if _, exists := userCache.Get("user1"); !exists {
		fmt.Println("✅ user1 was evicted (LRU policy working)")
	} else {
		fmt.Println("❌ user1 was not evicted")
	}

	// Test cache statistics
	fmt.Println()
	fmt.Println("4. CACHE STATISTICS:")
	stats := userCache.GetStats()
	fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate()*100)
	fmt.Printf("Miss Rate: %.2f%%\n", stats.MissRate()*100)
	fmt.Printf("Total Hits: %d\n", stats.GetHits())
	fmt.Printf("Total Misses: %d\n", stats.GetMisses())
	fmt.Printf("Evictions: %d\n", stats.GetEvictions())

	// Test cache warming
	fmt.Println()
	fmt.Println("5. CACHE WARMING:")
	warmData := map[string]interface{}{
		"session1": "active_session_123",
		"session2": "active_session_456",
		"session3": "active_session_789",
	}
	
	cacheService.WarmUpCache("sessions", warmData)
	fmt.Printf("Warmed up session cache. Size: %d\n", sessionCache.Size())
	fmt.Printf("Session cache keys: %v\n", sessionCache.Keys())

	// Test batch operations
	fmt.Println()
	fmt.Println("6. BATCH OPERATIONS:")
	batchData := map[string]interface{}{
		"config1": "database_url",
		"config2": "api_key",
		"config3": "debug_mode",
	}
	
	cacheService.BatchPut("config", batchData)
	fmt.Printf("Batch put completed. Config cache size: %d\n", configCache.Size())
	
	// Test batch get
	keys := []string{"config1", "config2", "config3", "config4"}
	batchResult := cacheService.BatchGet("config", keys)
	fmt.Printf("Batch get result: %v\n", batchResult)

	// Test cache capacity change
	fmt.Println()
	fmt.Println("7. CACHE CAPACITY MANAGEMENT:")
	fmt.Printf("Original config cache capacity: %d\n", configCache.Capacity())
	
	err := configCache.SetCapacity(5)
	if err != nil {
		fmt.Printf("Error changing capacity: %v\n", err)
	} else {
		fmt.Printf("New config cache capacity: %d\n", configCache.Capacity())
	}

	// Test cache monitoring
	fmt.Println()
	fmt.Println("8. CACHE MONITORING:")
	monitor := services.NewCacheMonitor(cacheService, 100*time.Millisecond)
	monitor.Start()
	
	// Simulate some activity
	for i := 0; i < 5; i++ {
		userCache.Get("user2") // Hit
		userCache.Get("user5") // Miss
		time.Sleep(50 * time.Millisecond)
	}
	
	// Collect some metrics
	time.Sleep(200 * time.Millisecond)
	monitor.Stop()
	
	fmt.Println("Cache monitoring completed")

	// Test load balancer
	fmt.Println()
	fmt.Println("9. CACHE LOAD BALANCER:")
	
	// Create multiple caches for load balancing
	cache1 := models.NewLRUCache(2)
	cache2 := models.NewLRUCache(2)
	cache3 := models.NewLRUCache(2)
	
	loadBalancer := services.NewCacheLoadBalancer([]*models.LRUCache{cache1, cache2, cache3})
	
	// Distribute data across caches
	loadBalancer.Put("key1", "value1")
	loadBalancer.Put("key2", "value2")
	loadBalancer.Put("key3", "value3")
	loadBalancer.Put("key4", "value4")
	
	fmt.Printf("Load balancer distributed 4 keys across 3 caches\n")
	fmt.Printf("Cache 1 size: %d\n", cache1.Size())
	fmt.Printf("Cache 2 size: %d\n", cache2.Size())
	fmt.Printf("Cache 3 size: %d\n", cache3.Size())

	// Test cache persistence simulation
	fmt.Println()
	fmt.Println("10. CACHE PERSISTENCE SIMULATION:")
	
	// Simulate cache serialization
	cacheInfo := userCache.GetCacheInfo()
	fmt.Printf("Cache info for serialization: %v\n", cacheInfo)
	
	// Simulate cache restoration
	newCache := models.NewLRUCache(3)
	restoreData := map[string]interface{}{
		"user2": "Bob Smith",
		"user3": "Charlie Brown", 
		"user4": "Diana Prince",
	}
	
	for key, value := range restoreData {
		newCache.Put(key, value)
	}
	
	fmt.Printf("Restored cache size: %d\n", newCache.Size())
	fmt.Printf("Restored cache order: %v\n", newCache.GetLRUOrder())

	// Test edge cases
	fmt.Println()
	fmt.Println("11. EDGE CASES:")
	
	// Test empty cache
	emptyCache := models.NewLRUCache(1)
	fmt.Printf("Empty cache size: %d\n", emptyCache.Size())
	fmt.Printf("Empty cache is empty: %t\n", emptyCache.IsEmpty())
	
	// Test single item cache
	emptyCache.Put("only", "item")
	fmt.Printf("Single item cache size: %d\n", emptyCache.Size())
	fmt.Printf("Single item cache is full: %t\n", emptyCache.IsFull())
	
	// Test capacity 1 cache
	emptyCache.Put("new", "item") // Should evict "only"
	if _, exists := emptyCache.Get("only"); !exists {
		fmt.Println("✅ Single capacity cache eviction working")
	}

	// Test concurrent access simulation
	fmt.Println()
	fmt.Println("12. CONCURRENT ACCESS SIMULATION:")
	
	// Simulate concurrent reads and writes
	go func() {
		for i := 0; i < 10; i++ {
			userCache.Put(fmt.Sprintf("concurrent_%d", i), fmt.Sprintf("value_%d", i))
		}
	}()
	
	go func() {
		for i := 0; i < 10; i++ {
			userCache.Get(fmt.Sprintf("concurrent_%d", i))
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Concurrent operations completed. Final cache size: %d\n", userCache.Size())

	// Final statistics
	fmt.Println()
	fmt.Println("13. FINAL STATISTICS:")
	
	allStats := cacheService.GetCacheStats()
	for cacheName, stats := range allStats {
		fmt.Printf("\n%s Cache Statistics:\n", cacheName)
		if statsMap, ok := stats.(map[string]interface{}); ok {
			fmt.Printf("  Size: %v\n", statsMap["size"])
			fmt.Printf("  Capacity: %v\n", statsMap["capacity"])
			fmt.Printf("  Hit Rate: %.2f%%\n", statsMap["hit_rate"].(float64)*100)
			fmt.Printf("  Evictions: %v\n", statsMap["evictions"])
		}
	}

	// Test cache cleanup
	fmt.Println()
	fmt.Println("14. CACHE CLEANUP:")
	
	cacheCount := len(cacheService.ListCaches())
	fmt.Printf("Total caches before cleanup: %d\n", cacheCount)
	
	cacheService.DeleteCache("config")
	cacheCount = len(cacheService.ListCaches())
	fmt.Printf("Total caches after cleanup: %d\n", cacheCount)

	fmt.Println()
	fmt.Println("=== END OF DEMONSTRATION ===")
}
