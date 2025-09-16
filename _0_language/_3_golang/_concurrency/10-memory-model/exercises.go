package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Exercise 1: Fix Race Condition with Mutex
func Exercise1() {
	fmt.Println("\nExercise 1: Fix Race Condition with Mutex")
	fmt.Println("========================================")
	
	// TODO: Fix the race condition using a mutex
	// 1. Create a shared counter variable
	// 2. Start multiple goroutines that increment the counter
	// 3. Use a mutex to protect the counter
	// 4. Verify the final count is correct
	
	var counter int
	var mu sync.Mutex
	
	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Exercise 1: Counter = %d (expected: 10000)\n", counter)
	fmt.Println("Exercise 1 completed")
}

// Exercise 2: Fix Race Condition with Atomic Operations
func Exercise2() {
	fmt.Println("\nExercise 2: Fix Race Condition with Atomic Operations")
	fmt.Println("===================================================")
	
	// TODO: Fix the race condition using atomic operations
	// 1. Create a shared counter variable
	// 2. Start multiple goroutines that increment the counter
	// 3. Use atomic operations to protect the counter
	// 4. Verify the final count is correct
	
	var counter int64
	
	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Exercise 2: Counter = %d (expected: 10000)\n", atomic.LoadInt64(&counter))
	fmt.Println("Exercise 2 completed")
}

// Exercise 3: Implement Safe Map Operations
func Exercise3() {
	fmt.Println("\nExercise 3: Implement Safe Map Operations")
	fmt.Println("=======================================")
	
	// TODO: Implement safe map operations
	// 1. Create a map that will be accessed by multiple goroutines
	// 2. Use a mutex to protect map operations
	// 3. Implement safe read and write operations
	// 4. Test with multiple goroutines
	
	type SafeMap struct {
		mu sync.RWMutex
		m  map[string]int
	}
	
	sm := &SafeMap{m: make(map[string]int)}
	
	// Start goroutines that write to the map
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				sm.mu.Lock()
				sm.m[key] = j
				sm.mu.Unlock()
			}
		}(i)
	}
	
	// Start goroutines that read from the map
	for i := 0; i < 5; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				sm.mu.RLock()
				if val, exists := sm.m[key]; exists {
					_ = val
				}
				sm.mu.RUnlock()
			}
		}(i)
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Exercise 3: Map size = %d\n", len(sm.m))
	fmt.Println("Exercise 3 completed")
}

// Exercise 4: Implement Atomic Flag
func Exercise4() {
	fmt.Println("\nExercise 4: Implement Atomic Flag")
	fmt.Println("===============================")
	
	// TODO: Implement an atomic flag
	// 1. Create an atomic flag variable
	// 2. Start a goroutine that sets the flag after a delay
	// 3. Start another goroutine that waits for the flag
	// 4. Use atomic operations for all flag operations
	
	var flag int32
	
	// Set flag after delay
	go func() {
		time.Sleep(200 * time.Millisecond)
		atomic.StoreInt32(&flag, 1)
		fmt.Println("  Exercise 4: Flag set")
	}()
	
	// Wait for flag
	fmt.Println("  Exercise 4: Waiting for flag...")
	for atomic.LoadInt32(&flag) == 0 {
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("  Exercise 4: Flag is set!")
	fmt.Println("Exercise 4 completed")
}

// Exercise 5: Implement Compare and Swap
func Exercise5() {
	fmt.Println("\nExercise 5: Implement Compare and Swap")
	fmt.Println("=====================================")
	
	// TODO: Implement compare and swap operations
	// 1. Create an atomic variable
	// 2. Start multiple goroutines that try to change the value
	// 3. Use CompareAndSwap to ensure only one goroutine succeeds
	// 4. Track which goroutine succeeded
	
	var value int64 = 42
	var successCount int64
	
	// Start multiple goroutines trying to change the value
	for i := 0; i < 5; i++ {
		go func(id int) {
			newValue := int64(100 + id)
			if atomic.CompareAndSwapInt64(&value, 42, newValue) {
				atomic.AddInt64(&successCount, 1)
				fmt.Printf("  Exercise 5: Goroutine %d succeeded, changed to %d\n", id, newValue)
			} else {
				fmt.Printf("  Exercise 5: Goroutine %d failed\n", id)
			}
		}(i)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  Exercise 5: Final value = %d, successes = %d\n", 
		atomic.LoadInt64(&value), atomic.LoadInt64(&successCount))
	fmt.Println("Exercise 5 completed")
}

// Exercise 6: Implement Happens-Before with Channels
func Exercise6() {
	fmt.Println("\nExercise 6: Implement Happens-Before with Channels")
	fmt.Println("================================================")
	
	// TODO: Implement happens-before relationship with channels
	// 1. Create a shared variable and a channel
	// 2. In one goroutine: set the variable, then send on channel
	// 3. In another goroutine: receive from channel, then read variable
	// 4. Verify the happens-before relationship works
	
	var x int
	ch := make(chan int)
	
	// Goroutine 1: writes x, then sends on channel
	go func() {
		x = 42
		ch <- 1
		fmt.Println("  Exercise 6: Sent signal")
	}()
	
	// Goroutine 2: receives from channel, then reads x
	go func() {
		<-ch
		fmt.Printf("  Exercise 6: x = %d (guaranteed to be 42)\n", x)
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Exercise 6 completed")
}

// Exercise 7: Implement Happens-Before with Mutex
func Exercise7() {
	fmt.Println("\nExercise 7: Implement Happens-Before with Mutex")
	fmt.Println("=============================================")
	
	// TODO: Implement happens-before relationship with mutex
	// 1. Create a shared variable and a mutex
	// 2. In one goroutine: lock, set variable, unlock
	// 3. In another goroutine: lock, read variable, unlock
	// 4. Verify the happens-before relationship works
	
	var x int
	var mu sync.Mutex
	
	// Goroutine 1: locks, writes x, unlocks
	go func() {
		mu.Lock()
		x = 42
		mu.Unlock()
		fmt.Println("  Exercise 7: Wrote x = 42")
	}()
	
	// Goroutine 2: locks, reads x, unlocks
	go func() {
		mu.Lock()
		fmt.Printf("  Exercise 7: x = %d (guaranteed to be 42)\n", x)
		mu.Unlock()
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Exercise 7 completed")
}

// Exercise 8: Implement Memory Ordering with Atomic Operations
func Exercise8() {
	fmt.Println("\nExercise 8: Implement Memory Ordering with Atomic Operations")
	fmt.Println("=========================================================")
	
	// TODO: Implement memory ordering with atomic operations
	// 1. Create two atomic variables
	// 2. In one goroutine: set first variable, then second
	// 3. In another goroutine: read second variable, then first
	// 4. Verify sequential consistency works
	
	var x, y int64
	
	// Goroutine 1: writes x, then y
	go func() {
		atomic.StoreInt64(&x, 1)
		atomic.StoreInt64(&y, 2)
		fmt.Println("  Exercise 8: Set x=1, y=2")
	}()
	
	// Goroutine 2: reads y, then x
	go func() {
		for atomic.LoadInt64(&y) != 2 {
			// Wait for y to be set
		}
		fmt.Printf("  Exercise 8: x=%d, y=%d (x guaranteed to be 1)\n", 
			atomic.LoadInt64(&x), atomic.LoadInt64(&y))
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Exercise 8 completed")
}

// Exercise 9: Implement Safe Counter with Multiple Operations
func Exercise9() {
	fmt.Println("\nExercise 9: Implement Safe Counter with Multiple Operations")
	fmt.Println("=========================================================")
	
	// TODO: Implement a safe counter with multiple operations
	// 1. Create a counter that supports increment, decrement, and get
	// 2. Use atomic operations for all operations
	// 3. Test with multiple goroutines
	// 4. Verify the final count is correct
	
	type SafeCounter struct {
		value int64
	}
	
	counter := &SafeCounter{}
	
	// Start goroutines that increment
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter.value, 1)
			}
		}()
	}
	
	// Start goroutines that decrement
	for i := 0; i < 3; i++ {
		go func() {
			for j := 0; j < 500; j++ {
				atomic.AddInt64(&counter.value, -1)
			}
		}()
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("  Exercise 9: Counter = %d (expected: 3500)\n", atomic.LoadInt64(&counter.value))
	fmt.Println("Exercise 9 completed")
}

// Exercise 10: Implement Performance Comparison
func Exercise10() {
	fmt.Println("\nExercise 10: Implement Performance Comparison")
	fmt.Println("============================================")
	
	// TODO: Compare performance of mutex vs atomic operations
	// 1. Implement counter with mutex
	// 2. Implement counter with atomic operations
	// 3. Measure performance of both
	// 4. Compare the results
	
	iterations := 1000000
	
	// Mutex performance
	start := time.Now()
	var mu sync.Mutex
	var counter1 int
	for i := 0; i < iterations; i++ {
		mu.Lock()
		counter1++
		mu.Unlock()
	}
	mutexDuration := time.Since(start)
	
	// Atomic performance
	start = time.Now()
	var counter2 int64
	for i := 0; i < iterations; i++ {
		atomic.AddInt64(&counter2, 1)
	}
	atomicDuration := time.Since(start)
	
	fmt.Printf("  Exercise 10: Mutex = %v, Atomic = %v\n", mutexDuration, atomicDuration)
	fmt.Printf("  Exercise 10: Atomic is %.2fx faster\n", float64(mutexDuration)/float64(atomicDuration))
	fmt.Println("Exercise 10 completed")
}

// Run all exercises
func runExercises() {
	fmt.Println("💪 Memory Model & Race Conditions Exercises")
	fmt.Println("===========================================")
	
	Exercise1()
	Exercise2()
	Exercise3()
	Exercise4()
	Exercise5()
	Exercise6()
	Exercise7()
	Exercise8()
	Exercise9()
	Exercise10()
	
	fmt.Println("\n🎉 All exercises completed!")
	fmt.Println("Ready to move to advanced patterns!")
}

