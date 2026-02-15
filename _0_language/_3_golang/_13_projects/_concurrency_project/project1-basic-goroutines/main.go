package main

import (
	"fmt"
	"time"
)

// Exercise 1: Basic Goroutines
// Create a simple program that prints numbers 1-10 using goroutines
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(100 * time.Millisecond) // Simulate some work
	}
}

// Exercise 2: Channel Communication
// Producer-consumer pattern with channels
func producer(numbers chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		numbers <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(numbers) // Close channel when done
}

func consumer(numbers <-chan int) {
	for num := range numbers {
		fmt.Printf("Consuming: %d\n", num)
		time.Sleep(300 * time.Millisecond)
	}
}

// Exercise 3: Select Statement
// Handle multiple channel operations with timeouts
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start goroutines that will send to channels
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// Use select to handle multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received: %s\n", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout! No message received")
			return
		}
	}
}

func main() {
	fmt.Println("=== Go Concurrency Learning - Project 1 ===")

	// Exercise 1: Basic Goroutines
	fmt.Println("Exercise 1: Basic Goroutines")
	fmt.Println("Running printNumbers in a goroutine...")
	go printNumbers()
	time.Sleep(2 * time.Second) // Wait for goroutine to complete
	fmt.Println()

	// Exercise 2: Channel Communication
	fmt.Println("Exercise 2: Channel Communication")
	numbers := make(chan int)
	go producer(numbers)
	consumer(numbers)
	fmt.Println()

	// Exercise 3: Select Statement
	fmt.Println("Exercise 3: Select Statement")
	selectExample()
	fmt.Println()

	fmt.Println("=== Basic exercises completed! ===")
	fmt.Println()

	// Exercise 4: Calculator Implementation
	fmt.Println("Exercise 4: Calculator Implementation")
	fmt.Println("Run: go run main.go calculator.go")
}
