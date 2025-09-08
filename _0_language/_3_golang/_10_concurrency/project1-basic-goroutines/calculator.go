package main

import (
	"fmt"
	"time"
)

// Operation represents a calculation operation
type Operation struct {
	Type      string  // "add", "subtract", "multiply", "divide"
	A, B      float64 // Operands
	Result    float64 // Result of the operation
	Error     error   // Any error that occurred
	Timestamp time.Time
}

// Calculator handles concurrent calculations
type Calculator struct {
	operations chan Operation
	results    chan Operation
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{
		operations: make(chan Operation, 10), // Buffered channel
		results:    make(chan Operation, 10),
	}
}

// Start begins the calculator's worker goroutines
func (c *Calculator) Start() {
	// Start worker goroutines
	for i := 0; i < 3; i++ { // 3 worker goroutines
		go c.worker(i)
	}
}

// Stop stops the calculator
func (c *Calculator) Stop() {
	close(c.operations)
	close(c.results)
}

// worker processes operations from the operations channel
func (c *Calculator) worker(workerID int) {
	for op := range c.operations {
		fmt.Printf("Worker %d processing: %s(%.2f, %.2f)\n", 
			workerID, op.Type, op.A, op.B)
		
		// Perform the calculation
		result, err := c.calculate(op.Type, op.A, op.B)
		
		// Update the operation with result
		op.Result = result
		op.Error = err
		op.Timestamp = time.Now()
		
		// Send result back
		c.results <- op
		
		// Simulate some processing time
		time.Sleep(100 * time.Millisecond)
	}
}

// calculate performs the actual calculation
func (c *Calculator) calculate(operation string, a, b float64) (float64, error) {
	switch operation {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "divide":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}

// AddOperation queues an operation for processing
func (c *Calculator) AddOperation(operationType string, a, b float64) {
	op := Operation{
		Type:  operationType,
		A:     a,
		B:     b,
		Timestamp: time.Now(),
	}
	
	select {
	case c.operations <- op:
		fmt.Printf("Queued operation: %s(%.2f, %.2f)\n", operationType, a, b)
	case <-time.After(1 * time.Second):
		fmt.Println("Failed to queue operation: timeout")
	}
}

// GetResult retrieves a result from the results channel
func (c *Calculator) GetResult() (Operation, bool) {
	select {
	case result := <-c.results:
		return result, true
	case <-time.After(2 * time.Second):
		return Operation{}, false
	}
}

// ProcessOperations processes a batch of operations concurrently
func (c *Calculator) ProcessOperations(operations []Operation) {
	fmt.Println("\n=== Processing Operations Concurrently ===")
	
	// Queue all operations
	for _, op := range operations {
		c.AddOperation(op.Type, op.A, op.B)
	}
	
	// Collect results
	results := make([]Operation, 0, len(operations))
	for i := 0; i < len(operations); i++ {
		if result, ok := c.GetResult(); ok {
			results = append(results, result)
		}
	}
	
	// Display results
	fmt.Println("\n=== Results ===")
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("%s(%.2f, %.2f) = ERROR: %v\n", 
				result.Type, result.A, result.B, result.Error)
		} else {
			fmt.Printf("%s(%.2f, %.2f) = %.2f (processed by worker at %v)\n", 
				result.Type, result.A, result.B, result.Result, result.Timestamp.Format("15:04:05"))
		}
	}
}

// Example usage and demonstration
func demonstrateCalculator() {
	calc := NewCalculator()
	calc.Start()
	defer calc.Stop()
	
	// Create some test operations
	operations := []Operation{
		{Type: "add", A: 10, B: 5},
		{Type: "subtract", A: 20, B: 8},
		{Type: "multiply", A: 7, B: 6},
		{Type: "divide", A: 15, B: 3},
		{Type: "divide", A: 10, B: 0}, // This will cause an error
		{Type: "add", A: 100, B: 200},
		{Type: "multiply", A: 3.14, B: 2},
		{Type: "subtract", A: 50, B: 25},
	}
	
	// Process all operations
	calc.ProcessOperations(operations)
}
