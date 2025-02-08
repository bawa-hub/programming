package main

import "fmt"
import "time"

func printMessage() {
    fmt.Println("Thread is running!")
    time.Sleep(2 * time.Second)
    fmt.Println("Thread has finished!")
}

func main() {
    fmt.Println("Before goroutine")
    go printMessage()  // Start the goroutine
    time.Sleep(3 * time.Second)  // Wait for goroutine to finish
    fmt.Println("After goroutine")
}
