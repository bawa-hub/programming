package calculator

import (
	"cli-calculator/pkg/errors"
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	calc := New()

	tests := []struct {
		a, b      float64
		operation string
		expected  float64
		expectError bool
	}{
		{2, 3, "+", 5, false},
		{10, 2, "/", 5, false},
		{5, 0, "/", 0, true}, // Division by zero
		{2, 3, "unknown", 0, true}, // Unknown operation
	}

	for _, test := range tests {
		result, err := calc.Calculate(test.a, test.b, test.operation)
		if test.expectError {
			if err == nil {
				t.Errorf("Calculator.Calculate(%f, %f, %s) expected error but got none", test.a, test.b, test.operation)
			}
		} else {
			if err != nil {
				t.Errorf("Calculator.Calculate(%f, %f, %s) returned error: %v", test.a, test.b, test.operation, err)
			}
			if result != test.expected {
				t.Errorf("Calculator.Calculate(%f, %f, %s) = %f, expected %f", test.a, test.b, test.operation, result, test.expected)
			}
		}
	}
}

func TestCalculator_CalculateUnary(t *testing.T) {
	calc := New()

	tests := []struct {
		a         float64
		operation string
		expected  float64
		expectError bool
	}{
		{25, "√", 5, false},
		{-4, "√", 0, true}, // Square root of negative number
		{2, "unknown", 0, true}, // Unknown operation
	}

	for _, test := range tests {
		result, err := calc.CalculateUnary(test.a, test.operation)
		if test.expectError {
			if err == nil {
				t.Errorf("Calculator.CalculateUnary(%f, %s) expected error but got none", test.a, test.operation)
			}
		} else {
			if err != nil {
				t.Errorf("Calculator.CalculateUnary(%f, %s) returned error: %v", test.a, test.operation, err)
			}
			if result != test.expected {
				t.Errorf("Calculator.CalculateUnary(%f, %s) = %f, expected %f", test.a, test.operation, result, test.expected)
			}
		}
	}
}

func TestCalculator_ParseAndCalculate(t *testing.T) {
	calc := New()

	tests := []struct {
		expression string
		expected   float64
		expectError bool
	}{
		{"2 + 3", 5, false},
		{"10 / 2", 5, false},
		{"√25", 5, false},
		{"2 * 3", 6, false},
		{"8 ^ 2", 64, false},
		{"7 % 3", 1, false},
		{"", 0, true}, // Empty expression
		{"2 +", 0, true}, // Invalid expression
		{"+ 2", 0, true}, // Invalid expression
		{"2 + + 3", 0, true}, // Invalid expression
		{"abc + 3", 0, true}, // Invalid number
	}

	for _, test := range tests {
		result, err := calc.ParseAndCalculate(test.expression)
		if test.expectError {
			if err == nil {
				t.Errorf("Calculator.ParseAndCalculate(%s) expected error but got none", test.expression)
			}
		} else {
			if err != nil {
				t.Errorf("Calculator.ParseAndCalculate(%s) returned error: %v", test.expression, err)
			}
			if result != test.expected {
				t.Errorf("Calculator.ParseAndCalculate(%s) = %f, expected %f", test.expression, result, test.expected)
			}
		}
	}
}

func TestCalculator_History(t *testing.T) {
	calc := New()

	// Initially empty
	history := calc.GetHistory()
	if len(history) != 0 {
		t.Errorf("Expected empty history, got %d entries", len(history))
	}

	// Add some calculations
	calc.Calculate(2, 3, "+")
	calc.Calculate(10, 2, "/")
	calc.CalculateUnary(25, "√")

	// Check history
	history = calc.GetHistory()
	if len(history) != 3 {
		t.Errorf("Expected 3 history entries, got %d", len(history))
	}

	expectedEntries := []string{
		"2.00 + 3.00 = 5.00",
		"10.00 / 2.00 = 5.00",
		"√ 25.00 = 5.00",
	}

	for i, expected := range expectedEntries {
		if i >= len(history) {
			t.Errorf("Missing history entry %d", i)
			continue
		}
		if history[i] != expected {
			t.Errorf("History entry %d = %s, expected %s", i, history[i], expected)
		}
	}

	// Clear history
	calc.ClearHistory()
	history = calc.GetHistory()
	if len(history) != 0 {
		t.Errorf("Expected empty history after clear, got %d entries", len(history))
	}
}

func TestCalculator_GetAvailableOperations(t *testing.T) {
	calc := New()
	operations := calc.GetAvailableOperations()

	expectedSymbols := []string{"+", "-", "*", "/", "^", "%", "√"}
	if len(operations) != len(expectedSymbols) {
		t.Errorf("Expected %d operations, got %d", len(expectedSymbols), len(operations))
	}

	for _, symbol := range expectedSymbols {
		if _, exists := operations[symbol]; !exists {
			t.Errorf("Missing operation for symbol: %s", symbol)
		}
	}
}

func TestCalculator_ErrorTypes(t *testing.T) {
	calc := New()

	// Test division by zero error type
	_, err := calc.Calculate(5, 0, "/")
	if err == nil {
		t.Error("Expected error for division by zero")
	}
	if calcErr, ok := err.(*errors.CalculatorError); ok {
		if calcErr.Type != errors.ErrMathError {
			t.Errorf("Expected error type %s, got %s", errors.ErrMathError, calcErr.Type)
		}
	} else {
		t.Error("Expected CalculatorError type")
	}

	// Test invalid operation error type
	_, err = calc.Calculate(2, 3, "unknown")
	if err == nil {
		t.Error("Expected error for unknown operation")
	}
	if calcErr, ok := err.(*errors.CalculatorError); ok {
		if calcErr.Type != errors.ErrInvalidOperation {
			t.Errorf("Expected error type %s, got %s", errors.ErrInvalidOperation, calcErr.Type)
		}
	} else {
		t.Error("Expected CalculatorError type")
	}
}
