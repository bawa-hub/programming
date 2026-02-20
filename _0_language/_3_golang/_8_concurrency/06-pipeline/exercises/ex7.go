package exercises

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 7: Pipeline with Metrics
// Create a pipeline that collects and reports performance metrics.
func Exercise7() {
	fmt.Println("\nExercise 7: Pipeline with Metrics")
	fmt.Println("=================================")
	
	const numItems = 8
	input := make(chan Data, numItems)
	output := make(chan Result, numItems)
	
	metrics := &ExercisePipelineMetrics{}
	
	// Stage 1 with metrics
	stage1 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage1)
		for data := range input {
			start := time.Now()
			time.Sleep(50 * time.Millisecond)
			metrics.recordStage("stage1", time.Since(start))
			stage1 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Metrics Stage1: %s", data.Value),
				Key:   data.Key,
				Stage: "stage1",
			}
		}
	}()
	
	// Stage 2 with metrics
	stage2 := make(chan ProcessedData, numItems)
	go func() {
		defer close(stage2)
		for data := range stage1 {
			start := time.Now()
			time.Sleep(50 * time.Millisecond)
			metrics.recordStage("stage2", time.Since(start))
			stage2 <- ProcessedData{
				ID:    data.ID,
				Value: fmt.Sprintf("Metrics Stage2: %s", data.Value),
				Key:   data.Key,
				Stage: "stage2",
			}
		}
	}()
	
	// Final stage with metrics
	go func() {
		defer close(output)
		for data := range stage2 {
			start := time.Now()
			time.Sleep(50 * time.Millisecond)
			metrics.recordStage("stage3", time.Since(start))
			output <- Result{
				ID:      data.ID,
				Value:   fmt.Sprintf("Metrics Final: %s", data.Value),
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
				Value: fmt.Sprintf("Metrics Item %d", i),
				Key:   fmt.Sprintf("metrics_key%d", i),
			}
		}
	}()
	
	// Collect results
	fmt.Println("Exercise 7 Results:")
	for result := range output {
		fmt.Printf("  Item %d: %s (stages: %v)\n", result.ID, result.Value, result.Stages)
	}
	
	// Print metrics
	fmt.Printf("\nExercise 7 Metrics:\n")
	fmt.Printf("  Stage1 Total Time: %v\n", metrics.getStageTime("stage1"))
	fmt.Printf("  Stage2 Total Time: %v\n", metrics.getStageTime("stage2"))
	fmt.Printf("  Stage3 Total Time: %v\n", metrics.getStageTime("stage3"))
	fmt.Printf("  Total Items Processed: %d\n", metrics.getTotalItems())
}

type ExercisePipelineMetrics struct {
	stage1Time    time.Duration
	stage2Time    time.Duration
	stage3Time    time.Duration
	totalItems    int64
	mu            sync.RWMutex
}

func (pm *ExercisePipelineMetrics) recordStage(stage string, duration time.Duration) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	pm.totalItems++
	switch stage {
	case "stage1":
		pm.stage1Time += duration
	case "stage2":
		pm.stage2Time += duration
	case "stage3":
		pm.stage3Time += duration
	}
}

func (pm *ExercisePipelineMetrics) getStageTime(stage string) time.Duration {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	
	switch stage {
	case "stage1":
		return pm.stage1Time
	case "stage2":
		return pm.stage2Time
	case "stage3":
		return pm.stage3Time
	default:
		return 0
	}
}

func (pm *ExercisePipelineMetrics) getTotalItems() int64 {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.totalItems
}