package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrent calculation patterns
func concurrentCalculationPatterns() {
	fmt.Println("=== Concurrent Calculation Patterns ===")

	// Pattern 1: Fan-out/Fan-in
	fmt.Println("1. Fan-out/Fan-in Pattern:")
	fanOutFanInExample()

	// Pattern 2: Pipeline
	fmt.Println("\n2. Pipeline Pattern:")
	pipelineExample()

	// Pattern 3: Worker Pool
	fmt.Println("\n3. Worker Pool Pattern:")
	workerPoolExample()

	// Pattern 4: Generator
	fmt.Println("\n4. Generator Pattern:")
	generatorExample()
}

// Fan-out/Fan-in: Distribute work to multiple workers, then collect results
func fanOutFanInExample() {
	input := make(chan int)
	output := make(chan int)

	// Fan-out: Multiple workers
	numWorkers := 3
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for num := range input {
				result := num * num // Square the number
				fmt.Printf("Worker %d: %d^2 = %d\n", workerID, num, result)
				output <- result
			}
		}(i)
	}

	// Close output when all workers are done
	go func() {
		wg.Wait()
		close(output)
	}()

	// Send input data
	go func() {
		for i := 1; i <= 6; i++ {
			input <- i
		}
		close(input)
	}()

	// Fan-in: Collect results
	fmt.Println("Results:")
	for result := range output {
		fmt.Printf("Received result: %d\n", result)
	}
}

// Pipeline: Chain of processing stages
func pipelineExample() {
	// Stage 1: Generate numbers
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Stage 2: Square numbers
	squares := make(chan int)
	go func() {
		for num := range numbers {
			square := num * num
			fmt.Printf("Stage 2: %d^2 = %d\n", num, square)
			squares <- square
		}
		close(squares)
	}()

	// Stage 3: Add 10
	results := make(chan int)
	go func() {
		for square := range squares {
			result := square + 10
			fmt.Printf("Stage 3: %d + 10 = %d\n", square, result)
			results <- result
		}
		close(results)
	}()

	// Collect final results
	fmt.Println("Pipeline results:")
	for result := range results {
		fmt.Printf("Final result: %d\n", result)
	}
}

// Worker Pool: Fixed number of workers processing jobs
func workerPoolExample() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	numWorkers := 3
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				// Simulate work
				time.Sleep(100 * time.Millisecond)
				result := job * 2
				fmt.Printf("Worker %d processed job %d -> %d\n", workerID, job, result)
				results <- result
			}
		}(i)
	}

	// Send jobs
	go func() {
		for i := 1; i <= 9; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("Worker pool results:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// Generator: Function that returns a channel
func generatorExample() {
	// Generator function
	fibonacci := func() <-chan int {
		ch := make(chan int)
		go func() {
			a, b := 0, 1
			for i := 0; i < 10; i++ {
				ch <- a
				a, b = b, a+b
			}
			close(ch)
		}()
		return ch
	}

	// Use the generator
	fmt.Println("Fibonacci sequence:")
	for num := range fibonacci() {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// Advanced: Concurrent calculation with error handling
func advancedConcurrentCalculation() {
	fmt.Println("\n=== Advanced Concurrent Calculation ===")

	type Calculation struct {
		Operation string
		A, B      float64
		Result    float64
		Error     error
	}

	// Input channel
	calculations := make(chan Calculation, 10)
	results := make(chan Calculation, 10)

	// Worker function
	worker := func(workerID int) {
		for calc := range calculations {
			var result float64
			var err error

			switch calc.Operation {
			case "add":
				result = calc.A + calc.B
			case "multiply":
				result = calc.A * calc.B
			case "divide":
				if calc.B == 0 {
					err = fmt.Errorf("division by zero")
				} else {
					result = calc.A / calc.B
				}
			default:
				err = fmt.Errorf("unknown operation: %s", calc.Operation)
			}

			calc.Result = result
			calc.Error = err
			fmt.Printf("Worker %d: %s(%.2f, %.2f) = %.2f (error: %v)\n", 
				workerID, calc.Operation, calc.A, calc.B, result, err)
			results <- calc
		}
	}

	// Start workers
	numWorkers := 2
	for i := 0; i < numWorkers; i++ {
		go worker(i)
	}

	// Send calculations
	go func() {
		calculations <- Calculation{"add", 10, 5, 0, nil}
		calculations <- Calculation{"multiply", 3, 4, 0, nil}
		calculations <- Calculation{"divide", 15, 3, 0, nil}
		calculations <- Calculation{"divide", 10, 0, 0, nil} // Error case
		calculations <- Calculation{"unknown", 1, 1, 0, nil} // Error case
		close(calculations)
	}()

	// Collect results
	go func() {
		for i := 0; i < 5; i++ {
			result := <-results
			if result.Error != nil {
				fmt.Printf("Error in calculation: %v\n", result.Error)
			} else {
				fmt.Printf("Successful calculation: %.2f\n", result.Result)
			}
		}
		close(results)
	}()

	// Wait a bit for processing
	time.Sleep(1 * time.Second)
}

func main() {
	concurrentCalculationPatterns()
	advancedConcurrentCalculation()
}
