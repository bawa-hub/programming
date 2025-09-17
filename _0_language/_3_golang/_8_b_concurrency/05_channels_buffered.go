package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Buffered vs Unbuffered Channels ===")
	
	// Example 1: Unbuffered channel (synchronous)
	fmt.Println("\n1. Unbuffered channel (synchronous):")
	unbuffered := make(chan string)
	
	go func() {
		fmt.Println("Goroutine: About to send message...")
		unbuffered <- "Hello from unbuffered!" // block here untill any goroutine receive this msg
		fmt.Println("Goroutine: Message sent!")
	}()
	
	time.Sleep(1 * time.Second) // Wait a bit
	fmt.Println("Main: About to receive...")
	message := <-unbuffered
	fmt.Printf("Main: Received: %s\n", message)
	
	// Example 2: Buffered channel (asynchronous up to buffer size)
	fmt.Println("\n2. Buffered channel (asynchronous):")
	buffered := make(chan string, 3) // Buffer size of 3
	
	go func() {
		fmt.Println("Goroutine: Sending message 1...")
		buffered <- "Message 1"
		fmt.Println("Goroutine: Sending message 2...")
		buffered <- "Message 2"
		fmt.Println("Goroutine: Sending message 3...")
		buffered <- "Message 3"
		fmt.Println("Goroutine: All messages sent!")
	}()
	
	time.Sleep(1 * time.Second) // Wait for all sends to complete
	fmt.Println("Main: About to receive all messages...")
	
	for i := 1; i <= 3; i++ {
		message := <-buffered
		fmt.Printf("Main: Received: %s\n", message)
	}
	
	fmt.Println("\nAll examples completed!")
}

// 1. The unbuffered goroutine’s timeline
// go func() {
//     fmt.Println("Goroutine: About to send message...")
//     unbuffered <- "Hello from unbuffered!" // blocks
//     fmt.Println("Goroutine: Message sent!")
// }()

// This goroutine starts immediately after being launched.
// It executes "About to send message...".
// Then it blocks at the send (unbuffered <- ...) because main has not received yet.
// It cannot print "Message sent!" until main executes message := <-unbuffered.

// 2. What main is doing
// time.Sleep(1 * time.Second)
// fmt.Println("Main: About to receive...")
// message := <-unbuffered
// fmt.Printf("Main: Received: %s\n", message)

// main is sleeping during this 1 second.
// While sleeping, no other code in main runs — so the buffered goroutine has not been created yet.
// When sleep ends, main immediately does the unbuffered receive.
// That unblocks the first goroutine, which then prints "Message sent!".

// 3. Then comes the buffered goroutine
// Only after finishing the unbuffered example does main reach this code:
// buffered := make(chan string, 3)
// go func() {
//     buffered <- "Message 1"
//     buffered <- "Message 2"
//     buffered <- "Message 3"
//     fmt.Println("Goroutine: All messages sent!")
// }()

// At this point, the unbuffered goroutine is already done.
// Now the buffered goroutine runs and prints "All messages sent!".

// ✅ So the correct ordering is:
// "Goroutine: About to send message..."
// "Main: About to receive..." (after 1s sleep)
// "Main: Received: Hello from unbuffered!"
// "Goroutine: Message sent!"
// "Goroutine: Sending message 1..." ... "All messages sent!"

// 🔎 Why you might have seen "All messages sent!" before "Message sent!"
// That would happen only if the print statements from goroutines get flushed to the console in a slightly reordered way (because stdout is shared and goroutines run concurrently).
// The Go runtime guarantees the program order inside a single goroutine.
// But across goroutines, the exact moment when logs hit the terminal can interleave differently.
// That’s why "All messages sent!" may appear visually before "Message sent!", even though logically "Message sent!" happened first.