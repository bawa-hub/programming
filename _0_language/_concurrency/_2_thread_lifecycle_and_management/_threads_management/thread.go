package main

import (
    "fmt"
    "time"
    "sync"
)

func printMessage(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Println("Goroutine running:", i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)  // Add 1 goroutine to the WaitGroup

    // Start the goroutine
    go printMessage(&wg)

    // Wait for the goroutine to finish
    wg.Wait()

    fmt.Println("Main goroutine ends.")
}

// In Go, we use a WaitGroup to wait for a goroutine to finish before the main function exits. 
// Goroutines are a more lightweight form of thread, and the sync package helps with managing their execution.