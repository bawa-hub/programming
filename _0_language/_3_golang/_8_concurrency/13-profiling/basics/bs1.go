package basics

import (
	"fmt"
	"runtime/pprof"
)

// Example 1: Basic CPU Profiling
func basicCPUProfiling() {
	fmt.Println("\n1. Basic CPU Profiling")
	fmt.Println("=====================")
	
	// Create CPU profile file
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("  Error creating CPU profile: %v\n", err)
		return
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Printf("  Error starting CPU profile: %v\n", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Run CPU-intensive task
	cpuIntensiveTask()
	
	fmt.Println("  CPU profiling completed")
}

func cpuIntensiveTask() {
	for i := 0; i < 1000000; i++ {
		// CPU-intensive operation
		_ = i * i
	}
}
