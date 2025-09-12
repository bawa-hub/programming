package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// MemoryAnalysis represents memory profile analysis results
type MemoryAnalysis struct {
	Filename      string        `json:"filename"`
	Duration      time.Duration `json:"duration"`
	TotalAllocs   uint64        `json:"total_allocs"`
	TotalBytes    uint64        `json:"total_bytes"`
	TopAllocs     []AllocInfo   `json:"top_allocs"`
	Leaks         []LeakInfo    `json:"leaks"`
	Recommendations []string    `json:"recommendations"`
}

// AllocInfo represents memory allocation information
type AllocInfo struct {
	Function    string  `json:"function"`
	Allocs      uint64  `json:"allocs"`
	Bytes       uint64  `json:"bytes"`
	Percent     float64 `json:"percent"`
	AvgSize     uint64  `json:"avg_size"`
	MaxSize     uint64  `json:"max_size"`
}

// LeakInfo represents memory leak information
type LeakInfo struct {
	Location    string  `json:"location"`
	Size        uint64  `json:"size_bytes"`
	Count       int     `json:"count"`
	Percent     float64 `json:"percent"`
	Description string  `json:"description"`
}

// MemoryBenchmark represents memory benchmark results
type MemoryBenchmark struct {
	Timestamp     time.Time     `json:"timestamp"`
	Duration      time.Duration `json:"duration"`
	Operations    int64         `json:"operations"`
	Allocs        uint64        `json:"allocs"`
	Bytes         uint64        `json:"bytes"`
	AllocsPerSec  float64       `json:"allocs_per_sec"`
	BytesPerSec   float64       `json:"bytes_per_sec"`
	AvgAllocSize  uint64        `json:"avg_alloc_size"`
	GCs           uint32        `json:"gc_cycles"`
	MemoryUsage   uint64        `json:"memory_usage_bytes"`
}

// AnalyzeMemoryProfile analyzes a memory profile
func (sp *SystemProfiler) AnalyzeMemoryProfile(filename string) (*MemoryAnalysis, error) {
	analysis := &MemoryAnalysis{
		Filename: filename,
		Duration: 30 * time.Second, // Default duration
	}
	
	// In a real implementation, you would parse the profile file
	// For now, we'll simulate the analysis
	
	// Simulate profile analysis
	analysis.TotalAllocs = 10000
	analysis.TotalBytes = 1024 * 1024 * 10 // 10MB
	
	analysis.TopAllocs = []AllocInfo{
		{
			Function: "main.memoryIntensiveTask",
			Allocs:   5000,
			Bytes:    5120 * 1024, // 5MB
			Percent:  50.0,
			AvgSize:  1024,
			MaxSize:  4096,
		},
		{
			Function: "runtime.mallocgc",
			Allocs:   3000,
			Bytes:    3072 * 1024, // 3MB
			Percent:  30.0,
			AvgSize:  1024,
			MaxSize:  2048,
		},
		{
			Function: "main.createLargeObject",
			Allocs:   2000,
			Bytes:    2048 * 1024, // 2MB
			Percent:  20.0,
			AvgSize:  1024,
			MaxSize:  8192,
		},
	}
	
	analysis.Leaks = []LeakInfo{
		{
			Location:    "main.memoryIntensiveTask:25",
			Size:        1024 * 1024, // 1MB
			Count:       100,
			Percent:     10.0,
			Description: "Potential memory leak in loop",
		},
		{
			Location:    "main.createLargeObject:10",
			Size:        512 * 1024, // 512KB
			Count:       50,
			Percent:     5.0,
			Description: "Large object not being freed",
		},
	}
	
	analysis.Recommendations = []string{
		"Optimize memoryIntensiveTask to reduce allocations",
		"Use object pools for frequently allocated objects",
		"Check for memory leaks in createLargeObject",
		"Consider using sync.Pool for temporary objects",
		"Profile memory allocation patterns more frequently",
	}
	
	return analysis, nil
}

// RunMemoryBenchmark runs a memory benchmark
func (sp *SystemProfiler) RunMemoryBenchmark() (*MemoryBenchmark, error) {
	benchmark := &MemoryBenchmark{
		Timestamp: time.Now(),
		Duration:  5 * time.Second,
	}
	
	// Get initial memory stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	initialAllocs := memStats.Mallocs
	initialBytes := memStats.TotalAlloc
	
	// Start benchmark
	start := time.Now()
	
	// Run memory-intensive operations
	operations := int64(0)
	for time.Since(start) < benchmark.Duration {
		// Simulate memory-intensive work
		sp.memoryIntensiveTask()
		operations++
	}
	
	// Get final memory stats
	runtime.ReadMemStats(&memStats)
	finalAllocs := memStats.Mallocs
	finalBytes := memStats.TotalAlloc
	
	benchmark.Operations = operations
	benchmark.Allocs = finalAllocs - initialAllocs
	benchmark.Bytes = finalBytes - initialBytes
	benchmark.AllocsPerSec = float64(benchmark.Allocs) / benchmark.Duration.Seconds()
	benchmark.BytesPerSec = float64(benchmark.Bytes) / benchmark.Duration.Seconds()
	if benchmark.Allocs > 0 {
		benchmark.AvgAllocSize = benchmark.Bytes / benchmark.Allocs
	} else {
		benchmark.AvgAllocSize = 0
	}
	benchmark.GCs = memStats.NumGC
	benchmark.MemoryUsage = memStats.HeapAlloc
	
	return benchmark, nil
}

// memoryIntensiveTask performs memory-intensive work
func (sp *SystemProfiler) memoryIntensiveTask() {
	// Simulate memory-intensive computation
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	
	// Simulate some processing
	sum := 0
	for _, b := range data {
		sum += int(b)
	}
	_ = sum // Prevent optimization
}

// createLargeObject creates a large object for testing
func (sp *SystemProfiler) createLargeObject() []byte {
	return make([]byte, 4096)
}

// PrintMemoryAnalysis prints memory analysis results
func (sp *SystemProfiler) PrintMemoryAnalysis(analysis *MemoryAnalysis) {
	fmt.Println("Memory Profile Analysis")
	fmt.Println("======================")
	fmt.Printf("File: %s\n", analysis.Filename)
	fmt.Printf("Duration: %v\n", analysis.Duration)
	fmt.Printf("Total Allocations: %d\n", analysis.TotalAllocs)
	fmt.Printf("Total Bytes: %s\n", formatBytes(analysis.TotalBytes))
	fmt.Println()
	
	// Print top allocations
	if len(analysis.TopAllocs) > 0 {
		fmt.Println("Top Allocations:")
		fmt.Printf("%-40s %-12s %-15s %-10s %-12s %-12s\n",
			"Function", "Allocs", "Bytes", "Percent", "Avg Size", "Max Size")
		fmt.Println(strings.Repeat("-", 110))
		
		for _, alloc := range analysis.TopAllocs {
			fmt.Printf("%-40s %-12d %-15s %-10.2f %-12d %-12d\n",
				alloc.Function,
				alloc.Allocs,
				formatBytes(alloc.Bytes),
				alloc.Percent,
				alloc.AvgSize,
				alloc.MaxSize)
		}
		fmt.Println()
	}
	
	// Print leaks
	if len(analysis.Leaks) > 0 {
		fmt.Println("Memory Leaks:")
		fmt.Printf("%-30s %-12s %-8s %-10s %-30s\n",
			"Location", "Size", "Count", "Percent", "Description")
		fmt.Println(strings.Repeat("-", 90))
		
		for _, leak := range analysis.Leaks {
			fmt.Printf("%-30s %-12s %-8d %-10.2f %-30s\n",
				leak.Location,
				formatBytes(leak.Size),
				leak.Count,
				leak.Percent,
				leak.Description)
		}
		fmt.Println()
	}
	
	// Print recommendations
	if len(analysis.Recommendations) > 0 {
		fmt.Println("Optimization Recommendations:")
		for i, rec := range analysis.Recommendations {
			fmt.Printf("%d. %s\n", i+1, rec)
		}
		fmt.Println()
	}
}

// PrintMemoryBenchmark prints memory benchmark results
func (sp *SystemProfiler) PrintMemoryBenchmark(benchmark *MemoryBenchmark) {
	fmt.Println("Memory Benchmark Results")
	fmt.Println("=======================")
	fmt.Printf("Timestamp: %s\n", benchmark.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("Duration: %v\n", benchmark.Duration)
	fmt.Printf("Operations: %d\n", benchmark.Operations)
	fmt.Printf("Allocations: %d\n", benchmark.Allocs)
	fmt.Printf("Bytes: %s\n", formatBytes(benchmark.Bytes))
	fmt.Printf("Allocations/sec: %.2f\n", benchmark.AllocsPerSec)
	fmt.Printf("Bytes/sec: %s\n", formatBytes(uint64(benchmark.BytesPerSec)))
	fmt.Printf("Average Alloc Size: %d bytes\n", benchmark.AvgAllocSize)
	fmt.Printf("GC Cycles: %d\n", benchmark.GCs)
	fmt.Printf("Memory Usage: %s\n", formatBytes(benchmark.MemoryUsage))
	fmt.Println()
}
