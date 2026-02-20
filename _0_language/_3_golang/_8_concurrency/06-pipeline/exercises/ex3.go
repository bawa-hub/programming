package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 3: Fan-Out/Fan-In Pipeline
// Create a pipeline that distributes work to multiple workers.
func Exercise3() {
	fmt.Println("\nExercise 3: Fan-Out/Fan-In Pipeline")
	fmt.Println("===================================")
	
	const numItems = 10
	const numWorkers = 3
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Fan-out: Distribute work to multiple workers
	workers := make([]chan Data, numWorkers)
	for i := range workers {
		workers[i] = make(chan Data)
	}
	
	// Start workers
	var wg sync.WaitGroup
	for i, worker := range workers {
		wg.Add(1)
		go func(workerID int, jobs <-chan Data) {
			defer wg.Done()
			for data := range jobs {
				time.Sleep(time.Duration(50+workerID*10) * time.Millisecond)
				result := Result{
					ID:      data.ID,
					Value:   fmt.Sprintf("FanOut Worker %d: %s", workerID, data.Value),
					Key:     data.Key,
					Stages:  []string{"fanout", "worker", "fanin"},
					Duration: time.Duration(50+workerID*10) * time.Millisecond,
				}
				output <- result
			}
		}(i, worker)
	}
	
	// Distribute work
	go func() {
		defer func() {
			for _, worker := range workers {
				close(worker)
			}
		}()
		
		for data := range input {
			workers[data.ID%numWorkers] <- data
		}
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("FanOut Item %d", i),
				Key:   fmt.Sprintf("fanout_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 3 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}