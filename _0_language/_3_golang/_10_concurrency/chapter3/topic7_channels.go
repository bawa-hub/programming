// 🧠 What is a Channel?

//     A channel is a pipe between goroutines — used to send and receive values safely between them.

// ch := make(chan int) // creates a channel of int type

// You can then:

//     ch <- value → send

//     val := <-ch → receive

// Channels synchronize automatically.

// 🔄 Unbuffered Channels
//     Send blocks until the receive happens, and vice versa.
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)
}
// 🧠 Key points:

//     go routine sends

//     main goroutine receives

//     neither moves forward until the other is ready

// 🧲 Buffered Channels
//     Let you send without waiting, up to a certain capacity.

ch := make(chan int, 2) // buffer of 2

ch <- 1
ch <- 2
// ch <- 3 // would block! buffer full

fmt.Println(<-ch) // reads 1
fmt.Println(<-ch) // reads 2

// Useful when:
//     You want a goroutine to offload work
//     You don’t need to wait for receiver immediately


// 📦 When to Use What?
// Type | Use when…
// Unbuffered | You want strong sync between goroutines
// Buffered | You want asynchronous message passing