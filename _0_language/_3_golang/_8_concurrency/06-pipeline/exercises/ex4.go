package exercises

import (
	"fmt"
	"time"
)

// Exercise 4: Pipeline with Error Handling
// Implement a pipeline that handles errors from any stage.
func Exercise4() {
	fmt.Println("\nExercise 4: Pipeline with Error Handling")
	fmt.Println("=======================================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	errors := make(chan error, numItems)
	
	// Stage 1 with error handling
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			if data.ID%3 == 0 {
				errors <- fmt.Errorf("stage1 failed for item %d", data.ID)
				continue
			}
			
			time.Sleep(50 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Error Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2 with error handling
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			if data.ID%2 == 0 {
				errors <- fmt.Errorf("stage2 failed for item %d", data.ID)
				continue
			}
			
			time.Sleep(50 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Error Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Final stage with error handling
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Error Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 150 * time.Millisecond,
			}
		}
	}()
	
	// Close errors channel when done
	go func() {
		time.Sleep(2 * time.Second)
		close(errors)
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Error Item %d", i),
				Key:   fmt.Sprintf("error_key%d", i),
			}
		}
	}()
	
	// Collect results and errors
	fmt.Println("Exercise 4 Results:")
	for {
		select {
		case result, ok := <-output:
			if !ok {
				output = nil
			} else {
				fmt.Printf("  SUCCESS: Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
			}
		case err, ok := <-errors:
			if !ok {
				errors = nil
			} else {
				fmt.Printf("  ERROR: %v\n", err)
			}
		}
		
		if output == nil && errors == nil {
			break
		}
	}
}