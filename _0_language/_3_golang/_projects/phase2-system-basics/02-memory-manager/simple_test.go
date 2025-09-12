package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	fmt.Println("🧠 Memory Manager - Simple Test")
	fmt.Println("===============================")
	
	// Get system memory info
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v\n", err)
		return
	}
	
	// Get Go runtime stats
	var goMemStats runtime.MemStats
	runtime.ReadMemStats(&goMemStats)
	
	// Print system memory
	fmt.Println("\n📊 System Memory:")
	fmt.Printf("  Total:     %s\n", formatBytes(memInfo.Total))
	fmt.Printf("  Used:      %s (%.2f%%)\n", formatBytes(memInfo.Used), memInfo.UsedPercent)
	fmt.Printf("  Available: %s\n", formatBytes(memInfo.Available))
	fmt.Printf("  Free:      %s\n", formatBytes(memInfo.Free))
	
	// Print Go runtime memory
	fmt.Println("\n🔧 Go Runtime Memory:")
	fmt.Printf("  Heap Alloc:    %s\n", formatBytes(goMemStats.HeapAlloc))
	fmt.Printf("  Heap Sys:      %s\n", formatBytes(goMemStats.HeapSys))
	fmt.Printf("  Heap Idle:     %s\n", formatBytes(goMemStats.HeapIdle))
	fmt.Printf("  Heap Inuse:    %s\n", formatBytes(goMemStats.HeapInuse))
	fmt.Printf("  Heap Released: %s\n", formatBytes(goMemStats.HeapReleased))
	fmt.Printf("  Heap Objects:  %d\n", goMemStats.HeapObjects)
	
	// Print stack memory
	fmt.Println("\n📚 Stack Memory:")
	fmt.Printf("  Stack Inuse: %s\n", formatBytes(goMemStats.StackInuse))
	fmt.Printf("  Stack Sys:   %s\n", formatBytes(goMemStats.StackSys))
	
	// Print GC info
	fmt.Println("\n🗑️  Garbage Collection:")
	fmt.Printf("  GC Cycles:     %d\n", goMemStats.NumGC)
	fmt.Printf("  GC CPU:        %.4f%%\n", goMemStats.GCCPUFraction*100)
	fmt.Printf("  GC Pause Total: %s\n", formatDuration(goMemStats.PauseTotalNs))
	
	// Print goroutines
	fmt.Println("\n🔄 Goroutines:")
	fmt.Printf("  Count: %d\n", runtime.NumGoroutine())
	
	// Test memory allocation
	fmt.Println("\n🧪 Testing Memory Allocation:")
	start := time.Now()
	
	// Allocate some memory
	data := make([][]byte, 1000)
	for i := range data {
		data[i] = make([]byte, 1024) // 1KB each
	}
	
	allocTime := time.Since(start)
	fmt.Printf("  Allocated 1000 x 1KB in %v\n", allocTime)
	
	// Get stats after allocation
	runtime.ReadMemStats(&goMemStats)
	fmt.Printf("  Heap Alloc after: %s\n", formatBytes(goMemStats.HeapAlloc))
	
	// Force GC
	fmt.Println("\n🗑️  Forcing Garbage Collection:")
	start = time.Now()
	runtime.GC()
	gcTime := time.Since(start)
	fmt.Printf("  GC completed in %v\n", gcTime)
	
	// Get stats after GC
	runtime.ReadMemStats(&goMemStats)
	fmt.Printf("  Heap Alloc after GC: %s\n", formatBytes(goMemStats.HeapAlloc))
	
	// Test memory pool
	fmt.Println("\n🏊 Testing Memory Pool:")
	pool := createMemoryPool(1024, 10)
	
	// Test pool operations
	obj, hit := pool.get()
	if hit {
		fmt.Printf("  Got object from pool: %t\n", obj != nil)
		pool.put(obj)
		fmt.Println("  Returned object to pool")
	}
	
	// Print pool stats
	fmt.Printf("  Pool Stats: %d used, %d available\n", pool.used, pool.available)
	
	fmt.Println("\n✅ Memory Manager test completed!")
}

// Simple memory pool implementation
type simplePool struct {
	objects   []interface{}
	used      int
	available int
}

func createMemoryPool(size int, count int) *simplePool {
	pool := &simplePool{
		objects:   make([]interface{}, count),
		available: count,
	}
	
	// Pre-allocate objects
	for i := 0; i < count; i++ {
		pool.objects[i] = make([]byte, size)
	}
	
	return pool
}

func (p *simplePool) get() (interface{}, bool) {
	if p.available > 0 {
		obj := p.objects[p.used]
		p.used++
		p.available--
		return obj, true
	}
	return nil, false
}

func (p *simplePool) put(obj interface{}) {
	if p.used > 0 {
		p.used--
		p.available++
	}
}

// Utility functions
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatDuration(ns uint64) string {
	if ns < 1000 {
		return fmt.Sprintf("%d ns", ns)
	}
	if ns < 1000000 {
		return fmt.Sprintf("%.2f µs", float64(ns)/1000)
	}
	if ns < 1000000000 {
		return fmt.Sprintf("%.2f ms", float64(ns)/1000000)
	}
	return fmt.Sprintf("%.2f s", float64(ns)/1000000000)
}
