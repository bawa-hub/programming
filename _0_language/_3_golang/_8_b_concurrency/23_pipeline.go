package main

import (
	"fmt"
	"sync"
	"time"
)

// Stage 1: Generate numbers
func generateNumbers(count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= count; i++ {
			fmt.Printf("Stage 1: Generating %d\n", i)
			out <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

// Stage 2: Square the numbers
func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			squared := n * n
			fmt.Printf("Stage 2: Squaring %d = %d\n", n, squared)
			out <- squared
			time.Sleep(150 * time.Millisecond)
		}
	}()
	return out
}

// Stage 3: Add 10 to the numbers
func addTen(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			result := n + 10
			fmt.Printf("Stage 3: Adding 10 to %d = %d\n", n, result)
			out <- result
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

// Stage 4: Print results
func printResults(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range in {
		fmt.Printf("Final result: %d\n", result)
	}
}

func main() {
	fmt.Println("=== Pipeline Pattern ===")
	
	var wg sync.WaitGroup
	wg.Add(1)
	
	// Create the pipeline: generate -> square -> addTen -> print
	numbers := generateNumbers(5)
	squared := squareNumbers(numbers)
	added := addTen(squared)
	
	// Start the final stage
	go printResults(added, &wg)
	
	// Wait for pipeline to complete
	wg.Wait()
	
	fmt.Println("Pipeline completed!")
}
