// Go uses goroutines for concurrent execution.

package main

import "fmt"

func printMessage() {
    fmt.Println("Thread is running!")
}

func main() {
    go printMessage()  // Start the goroutine
    // Give time for the goroutine to run
    fmt.Scanln()
}
