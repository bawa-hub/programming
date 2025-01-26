package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage("Hello from Goroutine") // Run as a goroutine
	printMessage("Hello from Main")         // Run in the main thread
}
