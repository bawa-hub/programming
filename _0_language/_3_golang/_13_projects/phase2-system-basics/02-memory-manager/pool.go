package main

import (
	"fmt"
	"strings"
	"time"
)

// NewMemoryPool creates a new memory pool
func NewMemoryPool(name string, size int, count int) *MemoryPool {
	pool := &MemoryPool{
		Name:      name,
		Size:      size,
		Count:     count,
		Available: count,
		Created:   time.Now(),
		objects:   make([]interface{}, 0, count),
	}
	
	// Pre-allocate objects
	for i := 0; i < count; i++ {
		obj := make([]byte, size)
		pool.objects = append(pool.objects, obj)
	}
	
	return pool
}

// Get gets an object from the pool
func (mp *MemoryPool) Get() (interface{}, bool) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	
	if mp.Available > 0 {
		obj := mp.objects[mp.Used]
		mp.Used++
		mp.Available--
		mp.LastUsed = time.Now()
		mp.TotalHits++
		
		// Update hit rate
		total := mp.TotalHits + mp.TotalMisses
		if total > 0 {
			mp.HitRate = float64(mp.TotalHits) / float64(total)
			mp.MissRate = float64(mp.TotalMisses) / float64(total)
		}
		
		return obj, true
	}
	
	mp.TotalMisses++
	
	// Update hit rate
	total := mp.TotalHits + mp.TotalMisses
	if total > 0 {
		mp.HitRate = float64(mp.TotalHits) / float64(total)
		mp.MissRate = float64(mp.TotalMisses) / float64(total)
	}
	
	return nil, false
}

// Put returns an object to the pool
func (mp *MemoryPool) Put(obj interface{}) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	
	if mp.Used > 0 {
		mp.Used--
		mp.Available++
		mp.LastUsed = time.Now()
	}
}

// Reset resets the pool
func (mp *MemoryPool) Reset() {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	
	mp.Used = 0
	mp.Available = mp.Count
	mp.TotalHits = 0
	mp.TotalMisses = 0
	mp.HitRate = 0
	mp.MissRate = 0
	mp.LastUsed = time.Now()
}

// CreateMemoryPool creates a new memory pool
func (mm *MemoryManager) CreateMemoryPool(name string, size int, count int) *MemoryPool {
	pool := NewMemoryPool(name, size, count)
	mm.mutex.Lock()
	defer mm.mutex.Unlock()
	mm.pools[name] = pool
	return pool
}

// GetMemoryPool gets a memory pool by name
func (mm *MemoryManager) GetMemoryPool(name string) (*MemoryPool, bool) {
	mm.mutex.RLock()
	defer mm.mutex.RUnlock()
	pool, exists := mm.pools[name]
	return pool, exists
}

// ListMemoryPools lists all memory pools
func (mm *MemoryManager) ListMemoryPools() map[string]*MemoryPool {
	mm.mutex.RLock()
	defer mm.mutex.RUnlock()
	
	pools := make(map[string]*MemoryPool)
	for name, pool := range mm.pools {
		pools[name] = pool
	}
	return pools
}

// TestMemoryPools tests memory pool performance
func (mm *MemoryManager) TestMemoryPools(opts *PoolOptions) (*PoolTestResults, error) {
	results := &PoolTestResults{
		Timestamp: time.Now(),
		Tests:     make([]PoolTest, 0),
	}
	
	// Test different pool sizes
	sizes := []int{64, 256, 1024, 4096, 16384}
	counts := []int{10, 50, 100, 500, 1000}
	
	for _, size := range sizes {
		for _, count := range counts {
			test := mm.runPoolTest(fmt.Sprintf("pool_%d_%d", size, count), size, count)
			results.Tests = append(results.Tests, test)
		}
	}
	
	return results, nil
}

// PoolTestResults represents pool test results
type PoolTestResults struct {
	Timestamp time.Time  `json:"timestamp"`
	Tests     []PoolTest `json:"tests"`
}

// PoolTest represents a single pool test
type PoolTest struct {
	Name           string        `json:"name"`
	Size           int           `json:"size"`
	Count          int           `json:"count"`
	Duration       time.Duration `json:"duration"`
	Operations     int           `json:"operations"`
	Hits           uint64        `json:"hits"`
	Misses         uint64        `json:"misses"`
	HitRate        float64       `json:"hit_rate"`
	MissRate       float64       `json:"miss_rate"`
	OpsPerSecond   float64       `json:"ops_per_second"`
	MemoryUsed     uint64        `json:"memory_used"`
	MemoryEfficient bool         `json:"memory_efficient"`
}

// runPoolTest runs a single pool test
func (mm *MemoryManager) runPoolTest(name string, size int, count int) PoolTest {
	// Create pool
	pool := NewMemoryPool(name, size, count)
	
	// Test parameters
	operations := 10000
	start := time.Now()
	
	// Run test
	for i := 0; i < operations; i++ {
		obj, hit := pool.Get()
		if hit {
			// Simulate some work
			_ = obj
			pool.Put(obj)
		}
	}
	
	duration := time.Since(start)
	stats := pool.GetStats()
	
	return PoolTest{
		Name:           name,
		Size:           size,
		Count:          count,
		Duration:       duration,
		Operations:     operations,
		Hits:           stats.TotalHits,
		Misses:         stats.TotalMisses,
		HitRate:        stats.HitRate,
		MissRate:       stats.MissRate,
		OpsPerSecond:   float64(operations) / duration.Seconds(),
		MemoryUsed:     uint64(size * count),
		MemoryEfficient: stats.HitRate > 0.8,
	}
}

// PrintPoolResults prints pool test results
func (mm *MemoryManager) PrintPoolResults(results *PoolTestResults) {
	fmt.Println("Memory Pool Test Results")
	fmt.Println("========================")
	fmt.Printf("Timestamp: %s\n", results.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("Total Tests: %d\n", len(results.Tests))
	fmt.Println()
	
	// Print header
	fmt.Printf("%-20s %-8s %-8s %-12s %-8s %-8s %-8s %-12s %-8s\n",
		"Name", "Size", "Count", "Duration", "Ops", "Hits", "Misses", "Hit Rate", "Efficient")
	fmt.Println(strings.Repeat("-", 100))
	
	// Print results
	for _, test := range results.Tests {
		fmt.Printf("%-20s %-8d %-8d %-12v %-8d %-8d %-8d %-12.2f %-8t\n",
			test.Name,
			test.Size,
			test.Count,
			test.Duration.Round(time.Microsecond),
			test.Operations,
			test.Hits,
			test.Misses,
			test.HitRate*100,
			test.MemoryEfficient)
	}
	
	fmt.Println()
	
	// Print summary
	efficientCount := 0
	totalHits := uint64(0)
	totalMisses := uint64(0)
	
	for _, test := range results.Tests {
		if test.MemoryEfficient {
			efficientCount++
		}
		totalHits += test.Hits
		totalMisses += test.Misses
	}
	
	fmt.Printf("Summary:\n")
	fmt.Printf("  Efficient Pools: %d/%d (%.2f%%)\n", 
		efficientCount, len(results.Tests), 
		float64(efficientCount)/float64(len(results.Tests))*100)
	fmt.Printf("  Total Hits: %d\n", totalHits)
	fmt.Printf("  Total Misses: %d\n", totalMisses)
	if totalHits+totalMisses > 0 {
		fmt.Printf("  Overall Hit Rate: %.2f%%\n", 
			float64(totalHits)/float64(totalHits+totalMisses)*100)
	}
}

// PrintDetailedStats prints detailed memory statistics
func (mm *MemoryManager) PrintDetailedStats(stats *DetailedStats) {
	fmt.Println("Detailed Memory Statistics")
	fmt.Println("=========================")
	fmt.Printf("Timestamp: %s\n", stats.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Println()
	
	// Print memory stats
	mm.printMemoryStats(stats.Memory)
	
	// Print allocator stats
	if len(stats.Allocators) > 0 {
		fmt.Println("Allocators:")
		fmt.Printf("%-20s %-15s %-15s %-15s %-15s %-12s\n",
			"Name", "Total Alloc", "Total Freed", "Current Size", "Allocations", "Fragmentation")
		fmt.Println(strings.Repeat("-", 100))
		
		for _, alloc := range stats.Allocators {
			fmt.Printf("%-20s %-15s %-15s %-15s %-15d %-12.2f%%\n",
				alloc.Name,
				formatBytes(alloc.TotalAllocated),
				formatBytes(alloc.TotalFreed),
				formatBytes(alloc.CurrentSize),
				alloc.AllocationCount,
				alloc.Fragmentation)
		}
		fmt.Println()
	}
	
	// Print pool stats
	if len(stats.Pools) > 0 {
		fmt.Println("Memory Pools:")
		fmt.Printf("%-20s %-8s %-8s %-8s %-8s %-8s %-8s %-12s\n",
			"Name", "Size", "Count", "Used", "Available", "Hits", "Misses", "Hit Rate")
		fmt.Println(strings.Repeat("-", 100))
		
		for _, pool := range stats.Pools {
			fmt.Printf("%-20s %-8d %-8d %-8d %-8d %-8d %-8d %-12.2f%%\n",
				pool.Name,
				pool.Size,
				pool.Count,
				pool.Used,
				pool.Available,
				pool.TotalHits,
				pool.TotalMisses,
				pool.HitRate*100)
		}
		fmt.Println()
	}
}
