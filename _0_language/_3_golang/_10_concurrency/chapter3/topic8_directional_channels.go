// 🎯 What Are Directional Channels?

// In Go, you can restrict a channel to be send-only or receive-only, which makes your code:

//     Safer

//     Easier to reason about

    // Better for large concurrent systems

// Syntax
// Type | Declaration | Meaning
// Bidirectional | chan int | Can send and receive
// Send-only | chan<- int | Can only send to the channel
// Receive-only | <-chan int | Can only receive from it



// ✏️ Example: Splitting Sender & Receiver
package main

import "fmt"

func sender(ch chan<- int) {
	ch <- 42
}

func receiver(ch <-chan int) {
	val := <-ch
	fmt.Println("Received:", val)
}

func main() {
	ch := make(chan int)

	go sender(ch)
	go receiver(ch)

	// Wait for goroutines to finish (simple delay for demo)
	select {}
}

// 🔐 Why Use Directional Channels?

//     Prevents accidental misuse

//     Enforces single-responsibility in goroutines

//     Cleaner APIs for concurrency primitives

// ✅ Example: Worker Pattern
// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("Worker", id, "processing job", j)
// 		results <- j * 2
// 	}
// }

// jobs is read-only → can't write to it
// results is write-only → can't read from it

// 🧠 Interview Insight

// Q: Why are directional channels important?

//     They enforce intent and reduce bugs by ensuring a goroutine only sends or only receives.

// Q: Can you convert a bidirectional channel to directional?

//     Yes! You can pass chan int as chan<- int or <-chan int inside functions.