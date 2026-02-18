package exercies

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Exercise 10: Performance Comparison
func Exercise10() {
	fmt.Println("\nExercise 10: Performance Comparison")
	fmt.Println("===================================")
	
	const iterations = 100000
	
	// Mutex performance
	start := time.Now()
	var mu sync.Mutex
	var counter1 int
	for i := 0; i < iterations; i++ {
		mu.Lock()
		counter1++
		mu.Unlock()
	}
	mutexTime := time.Since(start)
	
	// Atomic performance
	start = time.Now()
	var counter2 int64
	for i := 0; i < iterations; i++ {
		atomic.AddInt64(&counter2, 1)
	}
	atomicTime := time.Since(start)
	
	fmt.Printf("Mutex time: %v\n", mutexTime)
	fmt.Printf("Atomic time: %v\n", atomicTime)
	fmt.Printf("Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))
}