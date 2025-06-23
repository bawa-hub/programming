// Problem
// Spawn 3 goroutines that send their ID (1, 2, 3) into a channel, and the main goroutine should receive and print all 3 values.


package main

import (
	"fmt"
)

func sendID(id int, ch chan int) {
	ch <- id
}

func main() {
	ch := make(chan int)

	// Launch 3 goroutines
	go sendID(1, ch)
	go sendID(2, ch)
	go sendID(3, ch)

	// Receive 3 messages
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Println("Received:", val)
	}
}


// 🧠 Notes:

//     We use an unbuffered channel, so:

//         ch <- id will block until val := <-ch is ready to receive.

//     for i := 0; i < 3; i++ ensures we collect exactly 3 values.

// ⚠️ The order may vary (e.g., you might see 2, 1, 3) since goroutines run concurrently.


// 🚀 Buffered Channel Version of the Challenge
// package main

// import (
// 	"fmt"
// )

// func sendID(id int, ch chan int) {
// 	ch <- id
// }

// func main() {
// 	ch := make(chan int, 3) // Buffered channel with capacity 3

// 	// Launch 3 goroutines to send data
// 	go sendID(1, ch)
// 	go sendID(2, ch)
// 	go sendID(3, ch)

// 	// Receive 3 values from the buffer
// 	for i := 0; i < 3; i++ {
// 		val := <-ch
// 		fmt.Println("Received:", val)
// 	}
// }

// 🔍 Key Differences from Unbuffered Version:
// Concept | Unbuffered | Buffered
// Blocking Send | Sender blocks until receiver is ready | Sender does not block until buffer is full
// Parallelism | More tightly synchronized | Loosely synchronized
// Channel Capacity | make(chan int) | make(chan int, 3)

// In this buffered version:

//     All 3 goroutines can send immediately into the channel because there’s enough buffer space.

//     main() doesn’t need to be ready immediately to receive.

// 🧠 Interview Insight:

// Q: Why would you use a buffered channel?

//     To avoid blocking the sender goroutine and allow for async communication up to a limit.