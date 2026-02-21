package basics

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

// Example 4: Basic Block Profiling
func basicBlockProfiling() {
	fmt.Println("\n4. Basic Block Profiling")
	fmt.Println("========================")
	
	// Create block profile file
	f, err := os.Create("block.prof")
	if err != nil {
		fmt.Printf("  Error creating block profile: %v\n", err)
		return
	}
	defer f.Close()

	// Start block profiling
	runtime.SetBlockProfileRate(1)
	defer runtime.SetBlockProfileRate(0)

	ch := make(chan int)
	var wg sync.WaitGroup
	
	// Start goroutines that will block
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go blockingTask(ch, &wg)
	}
	
	// Send data to unblock
	for i := 0; i < 5; i++ {
		ch <- i
	}
	
	wg.Wait()
	close(ch)
	
	// Write block profile
	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		fmt.Printf("  Error writing block profile: %v\n", err)
		return
	}
	
	fmt.Println("  Block profiling completed")
}

func blockingTask(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Blocking operation
	data := <-ch
	_ = data
}