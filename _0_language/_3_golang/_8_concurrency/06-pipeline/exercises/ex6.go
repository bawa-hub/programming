package exercises

import (
	"fmt"
	"time"
)

// Exercise 6: Pipeline with Rate Limiting
// Implement a pipeline that limits the rate of data processing.
func Exercise6() {
	fmt.Println("\nExercise 6: Pipeline with Rate Limiting")
	fmt.Println("=======================================")
	
	const numItems = 6
	const rateLimit = 3 // items per second
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	// Rate limiter
	rateLimiter := time.NewTicker(time.Second / rateLimit)
	defer rateLimiter.Stop()
	
	// Stage 1 with rate limiting
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			<-rateLimiter.C
			time.Sleep(100 * time.Millisecond)
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Rate Limited Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2 with rate limiting
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			<-rateLimiter.C
			time.Sleep(100 * time.Millisecond)
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Rate Limited Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Final stage with rate limiting
	go func() {
		defer close(output)
		for data := range stage2 {
			<-rateLimiter.C
			time.Sleep(100 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Rate Limited Final: %s", data.Value),
				Key:     data.Key,
				Stages:  []string{"stage1", "stage2", "stage3"},
				Duration: 300 * time.Millisecond,
			}
		}
	}()
	
	// Submit data
	go func() {
		defer close(input)
		for i := 0; i < numItems; i++ {
			input <- Data{
				ID:    i,
				Value: fmt.Sprintf("Rate Limited Item %d", i),
				Key:   fmt.Sprintf("rate_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 6 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
}