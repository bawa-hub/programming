package exercises

import "fmt"

// Exercise 6: Pipeline Pattern
// Create a pipeline that processes data through multiple stages.
func Exercise6() {
	fmt.Println("\nExercise 6: Pipeline Pattern")
	fmt.Println("============================")
	
	// Create pipeline stages
	numbers := make(chan int)
	squares := make(chan int)
	results := make(chan int)
	
	// Stage 1: Generate numbers
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
			fmt.Printf("Stage 1: Generated %d\n", i)
		}
	}()
	
	// Stage 2: Square numbers
	go func() {
		defer close(squares)
		for n := range numbers {
			squared := n * n
			squares <- squared
			fmt.Printf("Stage 2: Squared %d = %d\n", n, squared)
		}
	}()
	
	// Stage 3: Add 10 to squares
	go func() {
		defer close(results)
		for s := range squares {
			result := s + 10
			results <- result
			fmt.Printf("Stage 3: Added 10 to %d = %d\n", s, result)
		}
	}()
	
	// Collect results
	fmt.Println("Pipeline results:")
	for result := range results {
		fmt.Printf("Final result: %d\n", result)
	}
}