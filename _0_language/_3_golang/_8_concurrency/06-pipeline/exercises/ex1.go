package exercises

import (
	"fmt"
	"time"
)

// Exercise 1: Basic Pipeline
// Create a basic pipeline with three stages.

// Data represents input data for the pipeline
type Data struct {
	ID    int
	Value string
	Key   string
}

// ProcessedData represents data after processing
type ProcessedData struct {
	ID    int
	Value string
	Key   string
	Stage string
}

// Result represents the final result
type Result struct {
	ID      int
	Value   string
	Key     string
	Stages  []string
	Duration time.Duration
}

func Exercise1() {
	fmt.Println("Exercise 1: Basic Pipeline")
	fmt.Println("==========================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Stage 1: Process input
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			time.Sleep(50 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Exercise1 Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2: Process stage1 output
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			time.Sleep(50 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Exercise1 Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Stage 3: Final processing
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Exercise1 Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Exercise Item %d", i),
				Key:   fmt.Sprintf("exercise_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 1 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}