package main

import (
	"fmt"
	"sync"
	"time"
)

// Fan-out: Distribute work to multiple workers
func fanOut(input <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)
	
	for i := 0; i < workers; i++ {
		output := make(chan int)
		outputs[i] = output
		
		go func(workerID int, out chan<- int) {
			defer close(out)
			for n := range input {
				// Simulate work
				result := n * n
				fmt.Printf("Worker %d: Processing %d -> %d\n", workerID, n, result)
				time.Sleep(200 * time.Millisecond)
				out <- result
			}
		}(i, output)
	}
	
	return outputs
}

// Fan-in: Collect results from multiple channels
func fanIn(inputs []<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup
	
	// Start a goroutine for each input channel
	for _, input := range inputs {
		wg.Add(1)
		go func(in <-chan int) {
			defer wg.Done()
			for n := range in {
				output <- n
			}
		}(input)
	}
	
	// Close output when all inputs are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	return output
}

func main() {
	fmt.Println("=== Fan-in/Fan-out Pattern ===")
	
	// Create input channel
	input := make(chan int)
	
	// Start input generator
	go func() {
		defer close(input)
		for i := 1; i <= 6; i++ {
			fmt.Printf("Input: Sending %d\n", i)
			input <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Fan-out to 3 workers
	workerOutputs := fanOut(input, 3)
	
	// Fan-in results
	results := fanIn(workerOutputs)
	
	// Collect and print results
	fmt.Println("\nCollecting results:")
	for result := range results {
		fmt.Printf("Final result: %d\n", result)
	}
	
	fmt.Println("Fan-in/Fan-out completed!")
}
