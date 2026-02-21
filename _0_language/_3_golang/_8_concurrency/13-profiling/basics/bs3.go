package basics

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

// Example 3: Basic Goroutine Profiling
func basicGoroutineProfiling() {
	fmt.Println("\n3. Basic Goroutine Profiling")
	fmt.Println("============================")
	
	// Create goroutine profile file
	f, err := os.Create("goroutine.prof")
	if err != nil {
		fmt.Printf("  Error creating goroutine profile: %v\n", err)
		return
	}
	defer f.Close()

	var wg sync.WaitGroup
	
	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go goroutineTask(i, &wg)
	}
	
	// Wait for completion
	wg.Wait()
	
	// Write goroutine profile
	if err := pprof.Lookup("goroutine").WriteTo(f, 0); err != nil {
		fmt.Printf("  Error writing goroutine profile: %v\n", err)
		return
	}
	
	fmt.Println("  Goroutine profiling completed")
}

func goroutineTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Millisecond)
	}
}