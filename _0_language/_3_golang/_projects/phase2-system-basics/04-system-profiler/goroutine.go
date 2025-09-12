package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

// GoroutineAnalysis represents goroutine profile analysis results
type GoroutineAnalysis struct {
	Filename      string           `json:"filename"`
	Duration      time.Duration    `json:"duration"`
	TotalGoroutines int           `json:"total_goroutines"`
	TopFunctions  []GoroutineInfo  `json:"top_functions"`
	Deadlocks     []DeadlockInfo   `json:"deadlocks"`
	Recommendations []string       `json:"recommendations"`
}

// GoroutineInfo represents goroutine information
type GoroutineInfo struct {
	Function    string `json:"function"`
	Count       int    `json:"count"`
	Percent     float64 `json:"percent"`
	State       string `json:"state"`
	Description string `json:"description"`
}

// DeadlockInfo represents deadlock information
type DeadlockInfo struct {
	Location    string `json:"location"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

// GoroutineBenchmark represents goroutine benchmark results
type GoroutineBenchmark struct {
	Timestamp      time.Time     `json:"timestamp"`
	Duration       time.Duration `json:"duration"`
	Operations     int64         `json:"operations"`
	Goroutines     int           `json:"goroutines"`
	MaxGoroutines  int           `json:"max_goroutines"`
	AvgGoroutines  float64       `json:"avg_goroutines"`
	GoroutinesPerSec float64     `json:"goroutines_per_sec"`
	MemoryUsage    uint64        `json:"memory_usage_bytes"`
	ContextSwitches int64        `json:"context_switches"`
}

// AnalyzeGoroutineProfile analyzes a goroutine profile
func (sp *SystemProfiler) AnalyzeGoroutineProfile(filename string) (*GoroutineAnalysis, error) {
	analysis := &GoroutineAnalysis{
		Filename: filename,
		Duration: 30 * time.Second, // Default duration
	}
	
	// In a real implementation, you would parse the profile file
	// For now, we'll simulate the analysis
	
	// Simulate profile analysis
	analysis.TotalGoroutines = runtime.NumGoroutine()
	
	analysis.TopFunctions = []GoroutineInfo{
		{
			Function:    "main.goroutineIntensiveTask",
			Count:       50,
			Percent:     45.5,
			State:       "running",
			Description: "Main worker goroutines",
		},
		{
			Function:    "runtime.gcBgMarkWorker",
			Count:       20,
			Percent:     18.2,
			State:       "running",
			Description: "Garbage collection workers",
		},
		{
			Function:    "main.networkHandler",
			Count:       15,
			Percent:     13.6,
			State:       "waiting",
			Description: "Network I/O goroutines",
		},
		{
			Function:    "main.databaseWorker",
			Count:       10,
			Percent:     9.1,
			State:       "waiting",
			Description: "Database worker goroutines",
		},
		{
			Function:    "runtime.sysmon",
			Count:       5,
			Percent:     4.5,
			State:       "running",
			Description: "System monitor goroutine",
		},
	}
	
	analysis.Deadlocks = []DeadlockInfo{
		{
			Location:    "main.goroutineIntensiveTask:30",
			Type:        "channel",
			Description: "Potential deadlock on channel receive",
			Severity:    "medium",
		},
		{
			Location:    "main.databaseWorker:15",
			Type:        "mutex",
			Description: "Potential deadlock on mutex lock",
			Severity:    "high",
		},
	}
	
	analysis.Recommendations = []string{
		"Review goroutine lifecycle management",
		"Check for potential deadlocks in channel operations",
		"Consider using context for goroutine cancellation",
		"Implement proper goroutine pooling",
		"Monitor goroutine count to prevent goroutine leaks",
	}
	
	return analysis, nil
}

// RunGoroutineBenchmark runs a goroutine benchmark
func (sp *SystemProfiler) RunGoroutineBenchmark() (*GoroutineBenchmark, error) {
	benchmark := &GoroutineBenchmark{
		Timestamp: time.Now(),
		Duration:  5 * time.Second,
	}
	
	// Get initial goroutine count
	initialGoroutines := runtime.NumGoroutine()
	maxGoroutines := initialGoroutines
	
	// Start benchmark
	start := time.Now()
	
	// Run goroutine-intensive operations
	operations := int64(0)
	var wg sync.WaitGroup
	
	for time.Since(start) < benchmark.Duration {
		// Create goroutines
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				sp.goroutineIntensiveTask()
			}()
		}
		
		// Update max goroutines
		currentGoroutines := runtime.NumGoroutine()
		if currentGoroutines > maxGoroutines {
			maxGoroutines = currentGoroutines
		}
		
		operations++
	}
	
	// Wait for all goroutines to complete
	wg.Wait()
	
	benchmark.Operations = operations
	benchmark.Goroutines = runtime.NumGoroutine()
	benchmark.MaxGoroutines = maxGoroutines
	benchmark.AvgGoroutines = float64(initialGoroutines+maxGoroutines) / 2
	benchmark.GoroutinesPerSec = float64(operations*10) / benchmark.Duration.Seconds()
	
	// Get memory usage
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	benchmark.MemoryUsage = memStats.HeapAlloc
	
	// Simulate context switches
	benchmark.ContextSwitches = operations * 100
	
	return benchmark, nil
}

// goroutineIntensiveTask performs goroutine-intensive work
func (sp *SystemProfiler) goroutineIntensiveTask() {
	// Simulate some work
	time.Sleep(10 * time.Millisecond)
	
	// Simulate channel operations
	ch := make(chan int, 1)
	ch <- 42
	<-ch
	close(ch)
}

// PrintGoroutineAnalysis prints goroutine analysis results
func (sp *SystemProfiler) PrintGoroutineAnalysis(analysis *GoroutineAnalysis) {
	fmt.Println("Goroutine Profile Analysis")
	fmt.Println("=========================")
	fmt.Printf("File: %s\n", analysis.Filename)
	fmt.Printf("Duration: %v\n", analysis.Duration)
	fmt.Printf("Total Goroutines: %d\n", analysis.TotalGoroutines)
	fmt.Println()
	
	// Print top functions
	if len(analysis.TopFunctions) > 0 {
		fmt.Println("Top Functions:")
		fmt.Printf("%-40s %-8s %-10s %-15s %-30s\n",
			"Function", "Count", "Percent", "State", "Description")
		fmt.Println(strings.Repeat("-", 110))
		
		for _, fn := range analysis.TopFunctions {
			fmt.Printf("%-40s %-8d %-10.2f %-15s %-30s\n",
				fn.Function,
				fn.Count,
				fn.Percent,
				fn.State,
				fn.Description)
		}
		fmt.Println()
	}
	
	// Print deadlocks
	if len(analysis.Deadlocks) > 0 {
		fmt.Println("Potential Deadlocks:")
		fmt.Printf("%-30s %-15s %-30s %-10s\n",
			"Location", "Type", "Description", "Severity")
		fmt.Println(strings.Repeat("-", 90))
		
		for _, deadlock := range analysis.Deadlocks {
			fmt.Printf("%-30s %-15s %-30s %-10s\n",
				deadlock.Location,
				deadlock.Type,
				deadlock.Description,
				deadlock.Severity)
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

// PrintGoroutineBenchmark prints goroutine benchmark results
func (sp *SystemProfiler) PrintGoroutineBenchmark(benchmark *GoroutineBenchmark) {
	fmt.Println("Goroutine Benchmark Results")
	fmt.Println("==========================")
	fmt.Printf("Timestamp: %s\n", benchmark.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("Duration: %v\n", benchmark.Duration)
	fmt.Printf("Operations: %d\n", benchmark.Operations)
	fmt.Printf("Goroutines: %d\n", benchmark.Goroutines)
	fmt.Printf("Max Goroutines: %d\n", benchmark.MaxGoroutines)
	fmt.Printf("Average Goroutines: %.2f\n", benchmark.AvgGoroutines)
	fmt.Printf("Goroutines/sec: %.2f\n", benchmark.GoroutinesPerSec)
	fmt.Printf("Memory Usage: %s\n", formatBytes(benchmark.MemoryUsage))
	fmt.Printf("Context Switches: %d\n", benchmark.ContextSwitches)
	fmt.Println()
}
