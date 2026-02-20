package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 10: Pipeline with Caching
// Implement a pipeline that caches results.
func Exercise10() {
	fmt.Println("\nExercise 10: Pipeline with Caching")
	fmt.Println("==================================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	cache := NewSimpleCache()
	
	// Stage 1 with caching
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			// Check cache first
			if cached, found := cache.Get(data.Key); found {
				stage1 <- cached
			} else {
				time.Sleep(50 * time.Millisecond)
				processed := ProcessedData{
					ID:    data.ID,
					Value: fmt.Sprintf("Cached Stage1: %s", data.Value),
					Key:   data.Key,
					Stage: "stage1",
				}
				cache.Set(data.Key, processed)
				stage1 <- processed
			}
		}
	}()
	
	// Stage 2 with caching
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			// Check cache first
			cacheKey := fmt.Sprintf("stage2_%s", data.Key)
			if cached, found := cache.Get(cacheKey); found {
				stage2 <- cached
			} else {
				time.Sleep(50 * time.Millisecond)
				processed := ProcessedData{
					ID:    data.ID,
					Value: fmt.Sprintf("Cached Stage2: %s", data.Value),
					Key:   data.Key,
					Stage: "stage2",
				}
				cache.Set(cacheKey, processed)
				stage2 <- processed
			}
		}
	}()
	
	// Final stage
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Cached Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Cached Item %d", i),
				Key:   fmt.Sprintf("cached_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 10 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
	
	// Print cache stats
	fmt.Printf("\nExercise 10 Cache Stats:\n")
	fmt.Printf("  Cache Hits: %d\n", cache.GetHits())
	fmt.Printf("  Cache Misses: %d\n", cache.GetMisses())
}

type SimpleCache struct {
	data   map[string]ProcessedData
	hits   int64
	misses int64
	mu     sync.RWMutex
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		data: make(map[string]ProcessedData),
	}
}

func (c *SimpleCache) Get(key string) (ProcessedData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if value, found := c.data[key]; found {
		c.hits++
		return value, true
	}
	c.misses++
	return ProcessedData{}, false
}

func (c *SimpleCache) Set(key string, value ProcessedData) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *SimpleCache) GetHits() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.hits
}

func (c *SimpleCache) GetMisses() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.misses
}

// Run all exercises
func RunAllExercises() {
	fmt.Println("🧪 Running All Pipeline Exercises")
	fmt.Println("=================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n✅ All exercises completed!")
}