// 🚀 2. Goroutines — Go’s Lightweight Threads
// 🔹 What is a Goroutine?

// A goroutine is a function or method that runs concurrently with other functions.
// It’s like a super lightweight thread, managed by the Go runtime.

//     They're cheap (can spawn thousands).

//     They’re non-blocking (don’t freeze your main thread).

//     Created with the go keyword.

// 🧪 Basic Syntax

// go someFunction()

// That’s it! Go handles the scheduling.

package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage("Goroutine") // runs concurrently
	printMessage("Main")         // runs in main thread
}

// ⚠️ Output:
// You’ll likely see interleaved prints from both "Main" and "Goroutine".

// 💡 Behind the Scenes

// Goroutines are multiplexed onto fewer OS threads by the Go runtime:

//     Go 1.5+ uses an M:N scheduler (M goroutines on N threads).

//     Go runtime decides when and where to run each goroutine.


// ⚠️ Common Mistake: Main Exits Early

// func sayHi() {
// 	println("Hi!")
// }

// func main() {
// 	go sayHi()
// }

// ❌ Problem: The program may exit before sayHi() executes.
// ✅ Fix: Use sync.WaitGroup (next topic) or time.Sleep (bad practice, but okay for demo).



// 🧠 Interview Insights

// Q: What’s the difference between a goroutine and a thread?

//     Goroutines are managed by the Go runtime, use less memory (2KB stack vs 1MB+), and are cheaper to create than OS threads.

// Q: How many goroutines can I spawn?

//     Thousands, even millions — Go will manage them across available cores.

// Q: Do goroutines run in parallel?

//     They can, if you allow it (GOMAXPROCS > 1) and you have multiple cores.

// 🔎 Recap
// Feature | Goroutines
// Created with | go keyword
// Managed by | Go runtime
// Stack size | Starts small (2KB)
// Cost | Very low
// Parallel? | Yes, if system allows