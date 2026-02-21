package basics

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

// Example 5: Basic Mutex Profiling
func basicMutexProfiling() {
	fmt.Println("\n5. Basic Mutex Profiling")
	fmt.Println("========================")
	
	// Create mutex profile file
	f, err := os.Create("mutex.prof")
	if err != nil {
		fmt.Printf("  Error creating mutex profile: %v\n", err)
		return
	}
	defer f.Close()

	// Start mutex profiling
	runtime.SetMutexProfileFraction(1)
	defer runtime.SetMutexProfileFraction(0)

	var mu sync.Mutex
	var wg sync.WaitGroup
	
	// Start multiple goroutines that will contend for mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mutexTask(&mu, &wg)
	}
	
	wg.Wait()
	
	// Write mutex profile
	if err := pprof.Lookup("mutex").WriteTo(f, 0); err != nil {
		fmt.Printf("  Error writing mutex profile: %v\n", err)
		return
	}
	
	fmt.Println("  Mutex profiling completed")
}

func mutexTask(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for i := 0; i < 1000; i++ {
		mu.Lock()
		// Critical section
		time.Sleep(1 * time.Microsecond)
		mu.Unlock()
	}
}
