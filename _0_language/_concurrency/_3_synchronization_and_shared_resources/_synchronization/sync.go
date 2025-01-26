package main

// In Go, you can use the sync.Mutex type to ensure that only one goroutine can access critical sections of code at a time.

import (
    "fmt"
    "sync"
)

type Counter struct {
    count int
    mu    sync.Mutex
}

func (c *Counter) increment() {
    c.mu.Lock()         // Lock the critical section
    c.count++
    c.mu.Unlock()       // Unlock the critical section
}

func main() {
    var counter Counter

    // Create two goroutines to increment the counter
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        for i := 0; i < 1000; i++ {
            counter.increment()
        }
        wg.Done()
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            counter.increment()
        }
        wg.Done()
    }()

    wg.Wait()  // Wait for both goroutines to finish
    fmt.Println("Final count:", counter.count)  // Expected output: 2000
}


// We use sync.Mutex to lock and unlock the critical section inside the increment() method, ensuring that only one goroutine can modify the counter at a time.
