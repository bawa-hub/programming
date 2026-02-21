package basics

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

// Example 2: Basic Memory Profiling
func basicMemoryProfiling() {
	fmt.Println("\n2. Basic Memory Profiling")
	fmt.Println("=========================")
	
	// Create memory profile file
	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Printf("  Error creating memory profile: %v\n", err)
		return
	}
	defer f.Close()

	// Run memory-intensive task
	memoryIntensiveTask()

	// Write memory profile
	runtime.GC() // Force garbage collection
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Printf("  Error writing memory profile: %v\n", err)
		return
	}
	
	fmt.Println("  Memory profiling completed")
}

func memoryIntensiveTask() {
	var data [][]int
	
	for i := 0; i < 1000; i++ {
		// Allocate memory
		slice := make([]int, 1000)
		data = append(data, slice)
	}
}