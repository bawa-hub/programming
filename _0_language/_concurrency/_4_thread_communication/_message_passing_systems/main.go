package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch chan int) {
	for item := range ch {
		fmt.Println("Consumed:", item)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 5)
	go producer(ch)
	go consumer(ch)

	time.Sleep(12 * time.Second) // Allow time for processing
}
