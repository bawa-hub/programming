package exercises

import (
	"fmt"
	"time"
)

// Exercise 2: Buffered Pipeline
// Implement a pipeline with buffered channels.
func Exercise2() {
	fmt.Println("\nExercise 2: Buffered Pipeline")
	fmt.Println("=============================")
	
	const numItems = 8
	const bufferSize = 3
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Stage 1 with buffering
	stage1 := make(chan ProcessedData, bufferSize)
	go func() {
		defer close(stage1)
		for data := range input {
			time.Sleep(30 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Buffered Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2 with buffering
	stage2 := make(chan ProcessedData, bufferSize)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			time.Sleep(30 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Buffered Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Final stage with buffering
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(30 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Buffered Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 90 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Buffered Item %d", i),
				Key:   fmt.Sprintf("buffered_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 2 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}