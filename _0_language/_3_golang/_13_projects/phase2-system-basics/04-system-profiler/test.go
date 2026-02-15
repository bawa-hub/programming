package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("📊 System Profiler - Test")
	fmt.Println("=========================")
	
	// Create system profiler
	sp := NewSystemProfiler()
	defer sp.Close()
	
	// Test basic functionality
	fmt.Println("\n📊 Getting system metrics...")
	metrics, err := sp.GetSystemMetrics()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// Print basic metrics
	fmt.Printf("CPU Usage: %.2f%%\n", metrics.CPU.Usage)
	fmt.Printf("Memory Used: %s / %s (%.2f%%)\n", 
		formatBytes(metrics.Memory.Used), 
		formatBytes(metrics.Memory.Total), 
		metrics.Memory.UsedPercent)
	fmt.Printf("Goroutines: %d\n", metrics.Goroutines.Count)
	fmt.Printf("GC Cycles: %d\n", metrics.GC.Cycles)
	fmt.Printf("Go Version: %s\n", metrics.Runtime.Version)
	
	// Test CPU profiling
	fmt.Println("\n🖥️  Testing CPU profiling...")
	err = sp.StartCPUProfile("cpu_test.prof")
	if err != nil {
		fmt.Printf("Error starting CPU profile: %v\n", err)
	} else {
		// Run some CPU-intensive work
		start := time.Now()
		for time.Since(start) < 2*time.Second {
			sp.cpuIntensiveTask()
		}
		
		err = sp.StopCPUProfile()
		if err != nil {
			fmt.Printf("Error stopping CPU profile: %v\n", err)
		} else {
			fmt.Println("CPU profile completed")
		}
	}
	
	// Test memory profiling
	fmt.Println("\n🧠 Testing memory profiling...")
	err = sp.StartMemoryProfile("mem_test.prof")
	if err != nil {
		fmt.Printf("Error starting memory profile: %v\n", err)
	} else {
		// Run some memory-intensive work
		start := time.Now()
		for time.Since(start) < 2*time.Second {
			sp.memoryIntensiveTask()
		}
		fmt.Println("Memory profile completed")
	}
	
	// Test goroutine profiling
	fmt.Println("\n🔄 Testing goroutine profiling...")
	err = sp.StartGoroutineProfile("goroutine_test.prof")
	if err != nil {
		fmt.Printf("Error starting goroutine profile: %v\n", err)
	} else {
		// Run some goroutine-intensive work
		start := time.Now()
		for time.Since(start) < 2*time.Second {
			sp.goroutineIntensiveTask()
		}
		fmt.Println("Goroutine profile completed")
	}
	
	// Test CPU analysis
	fmt.Println("\n🔍 Testing CPU analysis...")
	cpuAnalysis, err := sp.AnalyzeCPUProfile("cpu_test.prof")
	if err != nil {
		fmt.Printf("Error analyzing CPU profile: %v\n", err)
	} else {
		fmt.Printf("CPU Analysis: %d samples, %d functions\n", 
			cpuAnalysis.TotalSamples, len(cpuAnalysis.TopFunctions))
	}
	
	// Test memory analysis
	fmt.Println("\n🔍 Testing memory analysis...")
	memAnalysis, err := sp.AnalyzeMemoryProfile("mem_test.prof")
	if err != nil {
		fmt.Printf("Error analyzing memory profile: %v\n", err)
	} else {
		fmt.Printf("Memory Analysis: %d allocs, %s bytes\n", 
			memAnalysis.TotalAllocs, formatBytes(memAnalysis.TotalBytes))
	}
	
	// Test goroutine analysis
	fmt.Println("\n🔍 Testing goroutine analysis...")
	goroutineAnalysis, err := sp.AnalyzeGoroutineProfile("goroutine_test.prof")
	if err != nil {
		fmt.Printf("Error analyzing goroutine profile: %v\n", err)
	} else {
		fmt.Printf("Goroutine Analysis: %d goroutines, %d functions\n", 
			goroutineAnalysis.TotalGoroutines, len(goroutineAnalysis.TopFunctions))
	}
	
	// Test CPU benchmark
	fmt.Println("\n🏃 Testing CPU benchmark...")
	cpuBenchmark, err := sp.RunCPUBenchmark()
	if err != nil {
		fmt.Printf("Error running CPU benchmark: %v\n", err)
	} else {
		fmt.Printf("CPU Benchmark: %d ops, %.2f ops/sec\n", 
			cpuBenchmark.Operations, cpuBenchmark.OpsPerSec)
	}
	
	// Test memory benchmark
	fmt.Println("\n🏃 Testing memory benchmark...")
	memBenchmark, err := sp.RunMemoryBenchmark()
	if err != nil {
		fmt.Printf("Error running memory benchmark: %v\n", err)
	} else {
		fmt.Printf("Memory Benchmark: %d allocs, %.2f allocs/sec\n", 
			memBenchmark.Allocs, memBenchmark.AllocsPerSec)
	}
	
	// Test goroutine benchmark
	fmt.Println("\n🏃 Testing goroutine benchmark...")
	goroutineBenchmark, err := sp.RunGoroutineBenchmark()
	if err != nil {
		fmt.Printf("Error running goroutine benchmark: %v\n", err)
	} else {
		fmt.Printf("Goroutine Benchmark: %d ops, %.2f goroutines/sec\n", 
			goroutineBenchmark.Operations, goroutineBenchmark.GoroutinesPerSec)
	}
	
	fmt.Println("\n✅ System Profiler test completed!")
}
