package main

import (
	"fmt"
	"time"
)

// Generator demonstrates the generator pattern
func Generator() {
	fmt.Println("=== Generator Pattern ===")
	
	// Simple number generator
	numbers := generateNumbers(1, 5)
	
	fmt.Println("Generated numbers:")
	for n := range numbers {
		fmt.Printf("  %d\n", n)
	}
	
	// Fibonacci generator
	fib := generateFibonacci(10)
	
	fmt.Println("Fibonacci sequence:")
	for n := range fib {
		fmt.Printf("  %d\n", n)
	}
}

// generateNumbers creates a generator that produces numbers in a range
func generateNumbers(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i
		}
	}()
	return ch
}

// generateFibonacci creates a generator that produces Fibonacci numbers
func generateFibonacci(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 0, 1
		for i := 0; i < count; i++ {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}

// AdvancedGenerator demonstrates advanced generator patterns
func AdvancedGenerator() {
	fmt.Println("\n=== Advanced Generator Patterns ===")
	
	// Prime number generator
	primes := generatePrimes(20)
	
	fmt.Println("Prime numbers:")
	for n := range primes {
		fmt.Printf("  %d\n", n)
	}
	
	// Random number generator
	random := generateRandom(10, 1, 100)
	
	fmt.Println("Random numbers:")
	for n := range random {
		fmt.Printf("  %d\n", n)
	}
}

// generatePrimes creates a generator that produces prime numbers
func generatePrimes(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		
		primes := make([]int, 0)
		candidate := 2
		
		for len(primes) < count {
			isPrime := true
			for _, p := range primes {
				if candidate%p == 0 {
					isPrime = false
					break
				}
			}
			
			if isPrime {
				primes = append(primes, candidate)
				ch <- candidate
			}
			
			candidate++
		}
	}()
	return ch
}

// generateRandom creates a generator that produces random numbers
func generateRandom(count int, min, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		
		for i := 0; i < count; i++ {
			// Simple pseudo-random number generation
			seed := time.Now().UnixNano()
			random := int(seed%int64(max-min+1)) + min
			ch <- random
			time.Sleep(1 * time.Millisecond) // Ensure different seeds
		}
	}()
	return ch
}

// GeneratorWithState demonstrates a generator that maintains state
func GeneratorWithState() {
	fmt.Println("\n=== Generator with State ===")
	
	// Counter generator with state
	counter := newCounterGenerator(1, 10, 2) // start=1, end=10, step=2
	
	fmt.Println("Counter with step 2:")
	for n := range counter {
		fmt.Printf("  %d\n", n)
	}
	
	// Accumulator generator
	accumulator := newAccumulatorGenerator(0, 5)
	
	fmt.Println("Accumulator:")
	for n := range accumulator {
		fmt.Printf("  %d\n", n)
	}
}

// counterGenerator represents a counter generator with state
type counterGenerator struct {
	current int
	end     int
	step    int
}

// newCounterGenerator creates a new counter generator
func newCounterGenerator(start, end, step int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		
		counter := &counterGenerator{
			current: start,
			end:     end,
			step:    step,
		}
		
		for counter.current <= counter.end {
			ch <- counter.current
			counter.current += counter.step
		}
	}()
	return ch
}

// accumulatorGenerator represents an accumulator generator
type accumulatorGenerator struct {
	current int
	count   int
}

// newAccumulatorGenerator creates a new accumulator generator
func newAccumulatorGenerator(start, count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		
		accumulator := &accumulatorGenerator{
			current: start,
			count:   count,
		}
		
		for i := 0; i < accumulator.count; i++ {
			ch <- accumulator.current
			accumulator.current += i + 1
		}
	}()
	return ch
}

// GeneratorWithFilter demonstrates a generator with filtering
func GeneratorWithFilter() {
	fmt.Println("\n=== Generator with Filter ===")
	
	// Even numbers generator
	evens := generateEvenNumbers(1, 20)
	
	fmt.Println("Even numbers:")
	for n := range evens {
		fmt.Printf("  %d\n", n)
	}
	
	// Multiples of 3 generator
	multiples := generateMultiples(1, 30, 3)
	
	fmt.Println("Multiples of 3:")
	for n := range multiples {
		fmt.Printf("  %d\n", n)
	}
}

// generateEvenNumbers creates a generator that produces even numbers
func generateEvenNumbers(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			if i%2 == 0 {
				ch <- i
			}
		}
	}()
	return ch
}

// generateMultiples creates a generator that produces multiples of a number
func generateMultiples(start, end, multiple int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			if i%multiple == 0 {
				ch <- i
			}
		}
	}()
	return ch
}

// GeneratorWithTransform demonstrates a generator with transformation
func GeneratorWithTransform() {
	fmt.Println("\n=== Generator with Transform ===")
	
	// Square generator
	squares := generateSquares(1, 10)
	
	fmt.Println("Squares:")
	for n := range squares {
		fmt.Printf("  %d\n", n)
	}
	
	// String generator
	strings := generateStrings("item", 5)
	
	fmt.Println("Strings:")
	for s := range strings {
		fmt.Printf("  %s\n", s)
	}
}

// generateSquares creates a generator that produces squares
func generateSquares(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i * i
		}
	}()
	return ch
}

// generateStrings creates a generator that produces strings
func generateStrings(prefix string, count int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 1; i <= count; i++ {
			ch <- fmt.Sprintf("%s_%d", prefix, i)
		}
	}()
	return ch
}

// GeneratorWithTimeout demonstrates a generator with timeout
func GeneratorWithTimeout() {
	fmt.Println("\n=== Generator with Timeout ===")
	
	// Slow generator with timeout
	slow := generateSlowNumbers(10)
	
	timeout := time.After(3 * time.Second)
	
	fmt.Println("Slow numbers (with timeout):")
	for {
		select {
		case n, ok := <-slow:
			if !ok {
				fmt.Println("Generator completed")
				return
			}
			fmt.Printf("  %d\n", n)
		case <-timeout:
			fmt.Println("Timeout reached, stopping")
			return
		}
	}
}

// generateSlowNumbers creates a generator that produces numbers slowly
func generateSlowNumbers(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i
		}
	}()
	return ch
}

// GeneratorWithError demonstrates a generator with error handling
func GeneratorWithError() {
	fmt.Println("\n=== Generator with Error Handling ===")
	
	// Generator that can produce errors
	numbers := generateNumbersWithError(10)
	
	fmt.Println("Numbers with error handling:")
	for result := range numbers {
		if result.Error != nil {
			fmt.Printf("  ERROR: %v\n", result.Error)
		} else {
			fmt.Printf("  %d\n", result.Value)
		}
	}
}

// GeneratorResult represents the result of a generator operation
type GeneratorResult struct {
	Value int
	Error error
}

// generateNumbersWithError creates a generator that can produce errors
func generateNumbersWithError(count int) <-chan GeneratorResult {
	ch := make(chan GeneratorResult)
	go func() {
		defer close(ch)
		for i := 1; i <= count; i++ {
			// Simulate error for certain values
			if i%4 == 0 {
				ch <- GeneratorResult{Error: fmt.Errorf("error generating %d", i)}
			} else {
				ch <- GeneratorResult{Value: i}
			}
		}
	}()
	return ch
}

// GeneratorWithBackpressure demonstrates a generator with backpressure
func GeneratorWithBackpressure() {
	fmt.Println("\n=== Generator with Backpressure ===")
	
	// Fast generator with slow consumer
	fast := generateFastNumbers(20)
	
	fmt.Println("Fast numbers (with backpressure):")
	for n := range fast {
		fmt.Printf("  %d\n", n)
		// Slow consumer
		time.Sleep(200 * time.Millisecond)
	}
}

// generateFastNumbers creates a generator that produces numbers quickly
func generateFastNumbers(count int) <-chan int {
	ch := make(chan int, 5) // Buffered channel for backpressure
	go func() {
		defer close(ch)
		for i := 1; i <= count; i++ {
			ch <- i
			fmt.Printf("Generated: %d\n", i)
		}
	}()
	return ch
}

// GeneratorWithMetrics demonstrates a generator with metrics
func GeneratorWithMetrics() {
	fmt.Println("\n=== Generator with Metrics ===")
	
	// Generator with metrics
	numbers := generateNumbersWithMetrics(10)
	
	fmt.Println("Numbers with metrics:")
	for n := range numbers {
		fmt.Printf("  %d\n", n)
	}
}

// metricsGenerator represents a generator with metrics
type metricsGenerator struct {
	generated int
	startTime time.Time
}

// generateNumbersWithMetrics creates a generator with metrics
func generateNumbersWithMetrics(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		
		metrics := &metricsGenerator{
			startTime: time.Now(),
		}
		
		for i := 1; i <= count; i++ {
			metrics.generated++
			ch <- i
		}
		
		// Print metrics
		duration := time.Since(metrics.startTime)
		fmt.Printf("Metrics: Generated %d numbers in %v\n", 
			metrics.generated, duration)
	}()
	return ch
}
