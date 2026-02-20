package exercises

import (
	"fmt"
	"time"
)

// Exercise 9: Pipeline with Circuit Breaker
// Create a pipeline that uses circuit breakers.
func Exercise9() {
	fmt.Println("\nExercise 9: Pipeline with Circuit Breaker")
	fmt.Println("=========================================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	errors := make(chan error, numItems)
	
	// Simple circuit breaker simulation
	breaker1 := &SimpleCircuitBreaker{threshold: 3, timeout: 500 * time.Millisecond}
	breaker2 := &SimpleCircuitBreaker{threshold: 3, timeout: 500 * time.Millisecond}
	
	// Stage 1 with circuit breaker
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			if breaker1.Allow() {
				// Simulate occasional failures
				if data.ID%4 == 0 {
					breaker1.RecordFailure()
					errors <- fmt.Errorf("stage1 failed for item %d", data.ID)
					continue
				}
				
				time.Sleep(50 * time.Millisecond)
				breaker1.RecordSuccess()
				stage1 <- ProcessedData{
					ID:    data.ID,
					Value: fmt.Sprintf("Circuit Breaker Stage1: %s", data.Value),
					Key:   data.Key,
					Stage: "stage1",
				}
			} else {
				errors <- fmt.Errorf("circuit breaker open at stage1 for item %d", data.ID)
			}
		}
	}()
	
	// Stage 2 with circuit breaker
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			if breaker2.Allow() {
				// Simulate occasional failures
				if data.ID%3 == 0 {
					breaker2.RecordFailure()
					errors <- fmt.Errorf("stage2 failed for item %d", data.ID)
					continue
				}
				
				time.Sleep(50 * time.Millisecond)
				breaker2.RecordSuccess()
				stage2 <- ProcessedData{
					ID:    data.ID,
					Value: fmt.Sprintf("Circuit Breaker Stage2: %s", data.Value),
					Key:   data.Key,
					Stage: "stage2",
				}
			} else {
				errors <- fmt.Errorf("circuit breaker open at stage2 for item %d", data.ID)
			}
		}
	}()
	
	// Final stage
	go func() {
		defer close(output)
		for data := range stage2 {
			time.Sleep(50 * time.Millisecond)
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Circuit Breaker Final: %s", data.Value),
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
				Value: fmt.Sprintf("Circuit Breaker Item %d", i),
				Key:   fmt.Sprintf("circuit_key%d", i),
			}
		}
	}()
	
	// Collect results and errors
	fmt.Println("Exercise 9 Results:")
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

type SimpleCircuitBreaker struct {
	failures    int
	threshold   int
	timeout     time.Duration
	lastFailure time.Time
	state       int // 0: closed, 1: open, 2: half-open
}

func (cb *SimpleCircuitBreaker) Allow() bool {
	if cb.state == 0 { // closed
		return true
	} else if cb.state == 1 { // open
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = 2 // half-open
			return true
		}
		return false
	} else { // half-open
		return true
	}
}

func (cb *SimpleCircuitBreaker) RecordSuccess() {
	if cb.state == 2 { // half-open
		cb.state = 0 // closed
		cb.failures = 0
	}
}

func (cb *SimpleCircuitBreaker) RecordFailure() {
	cb.failures++
	cb.lastFailure = time.Now()
	
	if cb.failures >= cb.threshold {
		cb.state = 1 // open
	}
}