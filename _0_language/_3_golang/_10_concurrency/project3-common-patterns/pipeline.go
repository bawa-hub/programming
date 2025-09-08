package main

import (
	"fmt"
	"sync"
	"time"
)

// Pipeline demonstrates a simple pipeline pattern
func Pipeline() {
	fmt.Println("=== Simple Pipeline Pattern ===")
	
	// Stage 1: Generate numbers
	stage1 := make(chan int)
	go func() {
		defer close(stage1)
		for i := 1; i <= 5; i++ {
			fmt.Printf("Stage 1: generating %d\n", i)
			stage1 <- i
		}
	}()
	
	// Stage 2: Square numbers
	stage2 := make(chan int)
	go func() {
		defer close(stage2)
		for n := range stage1 {
			result := n * n
			fmt.Printf("Stage 2: squaring %d -> %d\n", n, result)
			stage2 <- result
		}
	}()
	
	// Stage 3: Add 10
	stage3 := make(chan int)
	go func() {
		defer close(stage3)
		for n := range stage2 {
			result := n + 10
			fmt.Printf("Stage 3: adding 10 to %d -> %d\n", n, result)
			stage3 <- result
		}
	}()
	
	// Collect final results
	fmt.Println("Pipeline results:")
	for result := range stage3 {
		fmt.Printf("  Final result: %d\n", result)
	}
}

// AdvancedPipeline demonstrates an advanced pipeline with error handling
func AdvancedPipeline() {
	fmt.Println("\n=== Advanced Pipeline with Error Handling ===")
	
	// Stage 1: Generate numbers
	stage1 := make(chan int)
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			fmt.Printf("Stage 1: generating %d\n", i)
			stage1 <- i
		}
	}()
	
	// Stage 2: Square numbers with error handling
	stage2 := make(chan PipelineResult)
	go func() {
		defer close(stage2)
		for n := range stage1 {
			result := PipelineResult{Input: n}
			
			// Simulate error for certain inputs
			if n%3 == 0 {
				result.Error = fmt.Errorf("error processing %d", n)
			} else {
				result.Output = n * n
			}
			
			fmt.Printf("Stage 2: processing %d -> %v\n", n, result)
			stage2 <- result
		}
	}()
	
	// Stage 3: Add 10 with error handling
	stage3 := make(chan PipelineResult)
	go func() {
		defer close(stage3)
		for result := range stage2 {
			if result.Error != nil {
				stage3 <- result // Pass error through
			} else {
				result.Output = result.Output + 10
				fmt.Printf("Stage 3: adding 10 to %d -> %d\n", result.Input, result.Output)
			}
			stage3 <- result
		}
	}()
	
	// Collect final results
	fmt.Println("Pipeline results:")
	for result := range stage3 {
		if result.Error != nil {
			fmt.Printf("  ERROR: %v\n", result.Error)
		} else {
			fmt.Printf("  Final result: %d\n", result.Output)
		}
	}
}

// PipelineResult represents the result of a pipeline stage
type PipelineResult struct {
	Input  int
	Output int
	Error  error
}

// ParallelPipeline demonstrates a pipeline with parallel stages
func ParallelPipeline() {
	fmt.Println("\n=== Parallel Pipeline ===")
	
	// Stage 1: Generate numbers
	stage1 := make(chan int)
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			fmt.Printf("Stage 1: generating %d\n", i)
			stage1 <- i
		}
	}()
	
	// Stage 2: Parallel processing (3 workers)
	stage2 := make(chan int)
	var wg sync.WaitGroup
	
	// Start 3 parallel workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for n := range stage1 {
				result := n * n
				fmt.Printf("Worker %d: squaring %d -> %d\n", workerID, n, result)
				stage2 <- result
			}
		}(i)
	}
	
	// Close stage2 when all workers are done
	go func() {
		wg.Wait()
		close(stage2)
	}()
	
	// Stage 3: Add 10
	stage3 := make(chan int)
	go func() {
		defer close(stage3)
		for n := range stage2 {
			result := n + 10
			fmt.Printf("Stage 3: adding 10 to %d -> %d\n", n, result)
			stage3 <- result
		}
	}()
	
	// Collect final results
	fmt.Println("Pipeline results:")
	for result := range stage3 {
		fmt.Printf("  Final result: %d\n", result)
	}
}

// BufferedPipeline demonstrates a pipeline with buffered channels
func BufferedPipeline() {
	fmt.Println("\n=== Buffered Pipeline ===")
	
	// Stage 1: Generate numbers (buffered)
	stage1 := make(chan int, 5)
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			fmt.Printf("Stage 1: generating %d\n", i)
			stage1 <- i
		}
	}()
	
	// Stage 2: Square numbers (buffered)
	stage2 := make(chan int, 5)
	go func() {
		defer close(stage2)
		for n := range stage1 {
			result := n * n
			fmt.Printf("Stage 2: squaring %d -> %d\n", n, result)
			stage2 <- result
		}
	}()
	
	// Stage 3: Add 10 (buffered)
	stage3 := make(chan int, 5)
	go func() {
		defer close(stage3)
		for n := range stage2 {
			result := n + 10
			fmt.Printf("Stage 3: adding 10 to %d -> %d\n", n, result)
			stage3 <- result
		}
	}()
	
	// Collect final results
	fmt.Println("Pipeline results:")
	for result := range stage3 {
		fmt.Printf("  Final result: %d\n", result)
	}
}

// PipelineWithTimeout demonstrates a pipeline with timeout
func PipelineWithTimeout() {
	fmt.Println("\n=== Pipeline with Timeout ===")
	
	// Stage 1: Generate numbers with timeout
	stage1 := make(chan int)
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			select {
			case stage1 <- i:
				fmt.Printf("Stage 1: generating %d\n", i)
			case <-time.After(1 * time.Second):
				fmt.Println("Stage 1: timeout, stopping")
				return
			}
		}
	}()
	
	// Stage 2: Square numbers with timeout
	stage2 := make(chan int)
	go func() {
		defer close(stage2)
		for n := range stage1 {
			select {
			case stage2 <- n * n:
				fmt.Printf("Stage 2: squaring %d -> %d\n", n, n*n)
			case <-time.After(1 * time.Second):
				fmt.Println("Stage 2: timeout, stopping")
				return
			}
		}
	}()
	
	// Stage 3: Add 10 with timeout
	stage3 := make(chan int)
	go func() {
		defer close(stage3)
		for n := range stage2 {
			select {
			case stage3 <- n + 10:
				fmt.Printf("Stage 3: adding 10 to %d -> %d\n", n, n+10)
			case <-time.After(1 * time.Second):
				fmt.Println("Stage 3: timeout, stopping")
				return
			}
		}
	}()
	
	// Collect final results with timeout
	timeout := time.After(5 * time.Second)
	
	fmt.Println("Pipeline results:")
	for {
		select {
		case result, ok := <-stage3:
			if !ok {
				fmt.Println("Pipeline completed")
				return
			}
			fmt.Printf("  Final result: %d\n", result)
		case <-timeout:
			fmt.Println("Pipeline timeout reached")
			return
		}
	}
}

// PipelineWithBackpressure demonstrates a pipeline with backpressure
func PipelineWithBackpressure() {
	fmt.Println("\n=== Pipeline with Backpressure ===")
	
	// Stage 1: Generate numbers
	stage1 := make(chan int, 2) // Small buffer
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			fmt.Printf("Stage 1: generating %d\n", i)
			stage1 <- i
		}
	}()
	
	// Stage 2: Slow processing (causes backpressure)
	stage2 := make(chan int, 2) // Small buffer
	go func() {
		defer close(stage2)
		for n := range stage1 {
			// Slow processing
			time.Sleep(500 * time.Millisecond)
			result := n * n
			fmt.Printf("Stage 2: squaring %d -> %d\n", n, result)
			stage2 <- result
		}
	}()
	
	// Stage 3: Fast processing
	stage3 := make(chan int, 2) // Small buffer
	go func() {
		defer close(stage3)
		for n := range stage2 {
			result := n + 10
			fmt.Printf("Stage 3: adding 10 to %d -> %d\n", n, result)
			stage3 <- result
		}
	}()
	
	// Collect final results
	fmt.Println("Pipeline results:")
	for result := range stage3 {
		fmt.Printf("  Final result: %d\n", result)
	}
}

// PipelineWithMetrics demonstrates a pipeline with metrics
func PipelineWithMetrics() {
	fmt.Println("\n=== Pipeline with Metrics ===")
	
	// Metrics
	metrics := &PipelineMetrics{}
	
	// Stage 1: Generate numbers
	stage1 := make(chan int)
	go func() {
		defer close(stage1)
		for i := 1; i <= 10; i++ {
			metrics.IncrementGenerated()
			fmt.Printf("Stage 1: generating %d\n", i)
			stage1 <- i
		}
	}()
	
	// Stage 2: Square numbers
	stage2 := make(chan int)
	go func() {
		defer close(stage2)
		for n := range stage1 {
			result := n * n
			metrics.IncrementProcessed()
			fmt.Printf("Stage 2: squaring %d -> %d\n", n, result)
			stage2 <- result
		}
	}()
	
	// Stage 3: Add 10
	stage3 := make(chan int)
	go func() {
		defer close(stage3)
		for n := range stage2 {
			result := n + 10
			metrics.IncrementCompleted()
			fmt.Printf("Stage 3: adding 10 to %d -> %d\n", n, result)
			stage3 <- result
		}
	}()
	
	// Collect final results
	fmt.Println("Pipeline results:")
	for result := range stage3 {
		fmt.Printf("  Final result: %d\n", result)
	}
	
	// Print metrics
	fmt.Printf("Metrics: Generated=%d, Processed=%d, Completed=%d\n", 
		metrics.GetGenerated(), metrics.GetProcessed(), metrics.GetCompleted())
}

// PipelineMetrics tracks pipeline metrics
type PipelineMetrics struct {
	generated int64
	processed int64
	completed int64
	mutex     sync.RWMutex
}

func (pm *PipelineMetrics) IncrementGenerated() {
	pm.mutex.Lock()
	pm.generated++
	pm.mutex.Unlock()
}

func (pm *PipelineMetrics) IncrementProcessed() {
	pm.mutex.Lock()
	pm.processed++
	pm.mutex.Unlock()
}

func (pm *PipelineMetrics) IncrementCompleted() {
	pm.mutex.Lock()
	pm.completed++
	pm.mutex.Unlock()
}

func (pm *PipelineMetrics) GetGenerated() int64 {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	return pm.generated
}

func (pm *PipelineMetrics) GetProcessed() int64 {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	return pm.processed
}

func (pm *PipelineMetrics) GetCompleted() int64 {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	return pm.completed
}
