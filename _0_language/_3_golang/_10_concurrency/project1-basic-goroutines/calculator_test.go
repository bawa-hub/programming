package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	calc := NewCalculator()
	calc.Start()
	defer calc.Stop()

	// Test basic operations
	tests := []struct {
		operation string
		a, b      float64
		expected  float64
		hasError  bool
	}{
		{"add", 2, 3, 5, false},
		{"subtract", 10, 4, 6, false},
		{"multiply", 3, 4, 12, false},
		{"divide", 15, 3, 5, false},
		{"divide", 10, 0, 0, true}, // Division by zero
		{"unknown", 1, 1, 0, true}, // Unknown operation
	}

	for _, test := range tests {
		t.Run(test.operation, func(t *testing.T) {
			calc.AddOperation(test.operation, test.a, test.b)
			
			result, ok := calc.GetResult()
			if !ok {
				t.Fatal("Failed to get result")
			}

			if test.hasError {
				if result.Error == nil {
					t.Errorf("Expected error for %s(%.2f, %.2f), got none", 
						test.operation, test.a, test.b)
				}
			} else {
				if result.Error != nil {
					t.Errorf("Unexpected error for %s(%.2f, %.2f): %v", 
						test.operation, test.a, test.b, result.Error)
				}
				if result.Result != test.expected {
					t.Errorf("Expected %.2f, got %.2f", test.expected, result.Result)
				}
			}
		})
	}
}

func TestConcurrentOperations(t *testing.T) {
	calc := NewCalculator()
	calc.Start()
	defer calc.Stop()

	// Add multiple operations concurrently
	numOperations := 10
	for i := 0; i < numOperations; i++ {
		calc.AddOperation("add", float64(i), float64(i+1))
	}

	// Collect all results
	results := make([]Operation, 0, numOperations)
	for i := 0; i < numOperations; i++ {
		if result, ok := calc.GetResult(); ok {
			results = append(results, result)
		}
	}

	if len(results) != numOperations {
		t.Errorf("Expected %d results, got %d", numOperations, len(results))
	}

	// Verify all results are correct (order may vary due to concurrency)
	expectedResults := make(map[float64]bool)
	for i := 0; i < numOperations; i++ {
		expected := float64(i) + float64(i+1)
		expectedResults[expected] = true
	}

	for _, result := range results {
		if !expectedResults[result.Result] {
			t.Errorf("Unexpected result: %.2f", result.Result)
		}
	}
}

func TestCalculatorTimeout(t *testing.T) {
	calc := NewCalculator()
	calc.Start()
	defer calc.Stop()

	// Try to get result without adding any operations
	_, ok := calc.GetResult()
	if ok {
		t.Error("Expected timeout, but got result")
	}
}

func BenchmarkCalculator(b *testing.B) {
	calc := NewCalculator()
	calc.Start()
	defer calc.Stop()

	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		calc.AddOperation("add", float64(i), float64(i+1))
		calc.GetResult()
	}
}

func ExampleCalculator() {
	calc := NewCalculator()
	calc.Start()
	defer calc.Stop()

	// Add some operations
	calc.AddOperation("add", 10, 5)
	calc.AddOperation("multiply", 3, 4)
	calc.AddOperation("divide", 20, 4)

	// Get results
	for i := 0; i < 3; i++ {
		if result, ok := calc.GetResult(); ok {
			if result.Error != nil {
				println("Error:", result.Error.Error())
			} else {
				println("Result:", result.Result)
			}
		}
	}
}
