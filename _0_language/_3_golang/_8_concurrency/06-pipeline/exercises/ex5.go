package exercises

import (
	"context"
	"fmt"
	"time"
)

// Exercise 5: Pipeline with Timeout
// Create a pipeline that handles timeouts for individual stages.
func Exercise5() {
	fmt.Println("\nExercise 5: Pipeline with Timeout")
	fmt.Println("=================================")
	
	const numItems = 6
	const timeout = 1 * time.Second
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	// Stage 1 with timeout
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for {
			select {
			case data, ok := <-input:
				if !ok {
					return
				}
				
				select {
				case <-time.After(time.Duration(data.ID*100) * time.Millisecond):
					stage1 <- ProcessedData{
						ID:    data.ID,
						Value: fmt.Sprintf("Timeout Stage1: %s", data.Value),
						Key:   data.Key,
						Stage: "stage1",
					}
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Stage 2 with timeout
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for {
			select {
			case data, ok := <-stage1:
				if !ok {
					return
				}
				
				select {
				case <-time.After(50 * time.Millisecond):
					stage2 <- ProcessedData{
						ID:    data.ID,
						Value: fmt.Sprintf("Timeout Stage2: %s", data.Value),
						Key:   data.Key,
						Stage: "stage2",
					}
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Final stage with timeout
	go func() {
		defer close(output)
		for {
			select {
			case data, ok := <-stage2:
				if !ok {
					return
				}
				
				select {
				case <-time.After(50 * time.Millisecond):
					output <- Result{
						ID:      data.ID,
						Value:   fmt.Sprintf("Timeout Final: %s", data.Value),
						Key:     data.Key,
						Stages:  []string{"stage1", "stage2", "stage3"},
						Duration: time.Duration(data.ID*100+100) * time.Millisecond,
					}
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			select {
			case input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Timeout Item %d", i),
				Key:   fmt.Sprintf("timeout_key%d", i),
			}:
			case <-ctx.Done():
				return
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 5 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}