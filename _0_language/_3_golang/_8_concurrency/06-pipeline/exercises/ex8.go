package exercises

import (
	"fmt"
	"time"
)

// Exercise 8: Pipeline with Backpressure
// Implement a pipeline that handles backpressure.
func Exercise8() {
	fmt.Println("\nExercise 8: Pipeline with Backpressure")
	fmt.Println("======================================")
	
	const numItems = 6
	const backpressureTimeout = 50 * time.Millisecond
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Stage 1 with backpressure
	stage1 := make(chan ProcessedData, 1) // Small buffer to trigger backpressure
	go func() {
		defer close(stage1)
		for data := range input {
			select {
			case stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Backpressure Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}:
				// Successfully sent
			case <-time.After(backpressureTimeout):
				// Backpressure detected, skip
				fmt.Printf("  Backpressure detected at stage1 for item %d\n", data.ID)
			}
		}
	}()
	
	// Stage 2 with backpressure
	stage2 := make(chan ProcessedData, 1) // Small buffer to trigger backpressure
	go func() {
		defer close(stage2)
		for data := range stage1 {
			select {
			case stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Backpressure Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}:
				// Successfully sent
			case <-time.After(backpressureTimeout):
				// Backpressure detected, skip
				fmt.Printf("  Backpressure detected at stage2 for item %d\n", data.ID)
			}
		}
	}()
	
	// Final stage with backpressure
	go func() {
		defer close(output)
		for data := range stage2 {
			select {
			case output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Backpressure Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}:
				// Successfully sent
			case <-time.After(backpressureTimeout):
				// Backpressure detected, skip
				fmt.Printf("  Backpressure detected at stage3 for item %d\n", data.ID)
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Backpressure Item %d", i),
				Key:   fmt.Sprintf("backpressure_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 8 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}