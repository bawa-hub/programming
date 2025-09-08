package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Basic Channels Example ===")
	
	// Example 1: Simple unbuffered channel
	fmt.Println("\n1. Simple unbuffered channel:")
	ch := make(chan string)
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "Hello from goroutine!"
	}()
	
	// This will block until the goroutine sends data
	message := <-ch
	fmt.Printf("Received: %s\n", message)
	
	// Example 2: Sending multiple values
	fmt.Println("\n2. Sending multiple values:")
	numbers := make(chan int)
	
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
			fmt.Printf("Sent: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(numbers) // Important: close channel when done sending
	}()

	// 👉 Why close is important?
    //    Because the receiver (range numbers) would block forever if we didn’t close the channel.
	
	// Receive all values
	for num := range numbers {
		fmt.Printf("Received: %d\n", num)
	}
	
	fmt.Println("\nAll values received!")
}

// 🖼 Timeline of Execution (simplified)
// Goroutine 1 (main)          Goroutine 2 (worker)

// <- ch  (blocks)             sleep 500ms
//                              ch <- "Hello..."
// receive message
// print "Received: Hello..."

// loop range numbers           for i=1..5:
// wait for send                 numbers <- i
// receive num                   print "Sent: i"
// print "Received: i"           sleep 100ms

// channel closes → loop ends
// print "All values received!"

// How it actually runs (unbuffered channel)
// First iteration (i = 1):
// Sender reaches (A) → tries numbers <- 1.
// Since channel is unbuffered, this blocks until someone is receiving.
// Main goroutine is sitting at (C): for num := range numbers.
// This executes <-numbers, which is waiting for a value.
// The runtime matches sender (A) with receiver (C).
// Transfer happens:
// Value 1 is copied from sender to receiver.
// Both goroutines unblock.
// Main goroutine continues at (D):
// Prints Received: 1.
// Sender goroutine continues at (B):
// Prints Sent: 1.


// Key point: Sending blocks until receiver is ready
// numbers <- i blocks until the main goroutine executes <-numbers.
// Once the receiver takes the value, the send unblocks and continues with fmt.Printf("Sent: %d\n", i).
// So the real timeline is:

// Ordering
// So for every value i:
// Sender: numbers <- i   (blocks)
// Main:   <-numbers      (receives i, unblocks both)
// Main:   print "Received: i"
// Sender: print "Sent: i"
