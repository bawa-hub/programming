package main

import "fmt"

func producer(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Println("Produced:", i)
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for item := range ch {
        fmt.Println("Consumed:", item)
    }
}

func main() {
    ch := make(chan int)

    go producer(ch)
    go consumer(ch)

    // Wait for the goroutines to finish
    var input string
    fmt.Scanln(&input)
}

// n Go, the producer sends integers to a channel, and the consumer reads from the same channel. The communication is done through the channel, which provides synchronization