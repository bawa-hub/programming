package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🧠 Memory Manager - Test")
	fmt.Println("========================")
	
	// Create memory manager
	mm := NewMemoryManager()
	defer mm.Close()
	
	// Test basic functionality
	fmt.Println("\n📊 Getting memory statistics...")
	stats, err := mm.GetMemoryStats()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// Print basic stats
	fmt.Printf("Total Memory: %s\n", formatBytes(stats.TotalMemory))
	fmt.Printf("Used Memory: %s (%.2f%%)\n", formatBytes(stats.UsedMemory), stats.MemoryPercent)
	fmt.Printf("Available Memory: %s\n", formatBytes(stats.AvailableMem))
	fmt.Printf("Heap Alloc: %s\n", formatBytes(stats.HeapAlloc))
	fmt.Printf("Goroutines: %d\n", stats.NumGoroutines)
	fmt.Printf("GC Cycles: %d\n", stats.NumGC)
	
	// Test memory pool
	fmt.Println("\n🏊 Testing memory pool...")
	pool := mm.CreateMemoryPool("test_pool", 1024, 10)
	
	// Test pool operations
	obj, hit := pool.Get()
	if hit {
		fmt.Printf("Got object from pool: %t\n", obj != nil)
		pool.Put(obj)
		fmt.Println("Returned object to pool")
	}
	
	// Test pool stats
	poolStats := pool.GetStats()
	fmt.Printf("Pool Stats: %d used, %d available, %.2f%% hit rate\n",
		poolStats.Used, poolStats.Available, poolStats.HitRate*100)
	
	// Test leak detection
	fmt.Println("\n🔍 Testing leak detection...")
	leakReport, err := mm.DetectLeaks(&LeakDetectOptions{
		Threshold: 10.0,
		Duration:  5 * time.Second,
	})
	if err != nil {
		fmt.Printf("Error detecting leaks: %v\n", err)
	} else {
		fmt.Printf("Leak detection completed: %d leaks found\n", len(leakReport.Leaks))
	}
	
	// Test optimization
	fmt.Println("\n⚡ Testing memory optimization...")
	optResults, err := mm.OptimizeMemory(&OptimizeOptions{})
	if err != nil {
		fmt.Printf("Error optimizing memory: %v\n", err)
	} else {
		fmt.Printf("Optimization completed: %d actions performed\n", len(optResults.Actions))
		fmt.Printf("Memory saved: %s (%.2f%%)\n",
			formatBytes(uint64(optResults.MemorySaved)),
			optResults.PercentageSaved)
	}
	
	fmt.Println("\n✅ Memory Manager test completed!")
}
