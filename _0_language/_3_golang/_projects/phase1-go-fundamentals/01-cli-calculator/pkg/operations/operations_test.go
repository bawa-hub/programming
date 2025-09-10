package operations

import (
	"math"
	"testing"
)

func TestAddOperation(t *testing.T) {
	op := AddOperation{}
	
	tests := []struct {
		a, b, expected float64
	}{
		{2, 3, 5},
		{-2, 3, 1},
		{0, 0, 0},
		{1.5, 2.5, 4.0},
		{-1.5, -2.5, -4.0},
	}

	for _, test := range tests {
		result, err := op.Execute(test.a, test.b)
		if err != nil {
			t.Errorf("AddOperation.Execute(%f, %f) returned error: %v", test.a, test.b, err)
		}
		if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("AddOperation.Execute(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
		}
	}

	if op.Symbol() != "+" {
		t.Errorf("AddOperation.Symbol() = %s, expected +", op.Symbol())
	}
}

func TestSubtractOperation(t *testing.T) {
	op := SubtractOperation{}
	
	tests := []struct {
		a, b, expected float64
	}{
		{5, 3, 2},
		{3, 5, -2},
		{0, 0, 0},
		{2.5, 1.5, 1.0},
		{-2.5, 1.5, -4.0},
	}

	for _, test := range tests {
		result, err := op.Execute(test.a, test.b)
		if err != nil {
			t.Errorf("SubtractOperation.Execute(%f, %f) returned error: %v", test.a, test.b, err)
		}
		if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("SubtractOperation.Execute(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
		}
	}

	if op.Symbol() != "-" {
		t.Errorf("SubtractOperation.Symbol() = %s, expected -", op.Symbol())
	}
}

func TestMultiplyOperation(t *testing.T) {
	op := MultiplyOperation{}
	
	tests := []struct {
		a, b, expected float64
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{0, 5, 0},
		{1.5, 2, 3.0},
		{-1.5, -2, 3.0},
	}

	for _, test := range tests {
		result, err := op.Execute(test.a, test.b)
		if err != nil {
			t.Errorf("MultiplyOperation.Execute(%f, %f) returned error: %v", test.a, test.b, err)
		}
		if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("MultiplyOperation.Execute(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
		}
	}

	if op.Symbol() != "*" {
		t.Errorf("MultiplyOperation.Symbol() = %s, expected *", op.Symbol())
	}
}

func TestDivideOperation(t *testing.T) {
	op := DivideOperation{}
	
	tests := []struct {
		a, b, expected float64
		expectError    bool
	}{
		{6, 2, 3, false},
		{5, 2, 2.5, false},
		{-6, 2, -3, false},
		{0, 5, 0, false},
		{5, 0, 0, true}, // Division by zero
	}

	for _, test := range tests {
		result, err := op.Execute(test.a, test.b)
		if test.expectError {
			if err == nil {
				t.Errorf("DivideOperation.Execute(%f, %f) expected error but got none", test.a, test.b)
			}
		} else {
			if err != nil {
				t.Errorf("DivideOperation.Execute(%f, %f) returned error: %v", test.a, test.b, err)
			}
			if math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("DivideOperation.Execute(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
			}
		}
	}

	if op.Symbol() != "/" {
		t.Errorf("DivideOperation.Symbol() = %s, expected /", op.Symbol())
	}
}

func TestPowerOperation(t *testing.T) {
	op := PowerOperation{}
	
	tests := []struct {
		a, b, expected float64
		expectError    bool
	}{
		{2, 3, 8, false},
		{4, 0.5, 2, false},
		{2, -1, 0.5, false},
		{0, 5, 0, false},
		{-2, 0.5, 0, true}, // Invalid power operation
	}

	for _, test := range tests {
		result, err := op.Execute(test.a, test.b)
		if test.expectError {
			if err == nil {
				t.Errorf("PowerOperation.Execute(%f, %f) expected error but got none", test.a, test.b)
			}
		} else {
			if err != nil {
				t.Errorf("PowerOperation.Execute(%f, %f) returned error: %v", test.a, test.b, err)
			}
			if math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("PowerOperation.Execute(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
			}
		}
	}

	if op.Symbol() != "^" {
		t.Errorf("PowerOperation.Symbol() = %s, expected ^", op.Symbol())
	}
}

func TestModuloOperation(t *testing.T) {
	op := ModuloOperation{}
	
	tests := []struct {
		a, b, expected float64
		expectError    bool
	}{
		{7, 3, 1, false},
		{8, 4, 0, false},
		{5.5, 2.5, 0.5, false},
		{5, 0, 0, true}, // Modulo by zero
	}

	for _, test := range tests {
		result, err := op.Execute(test.a, test.b)
		if test.expectError {
			if err == nil {
				t.Errorf("ModuloOperation.Execute(%f, %f) expected error but got none", test.a, test.b)
			}
		} else {
			if err != nil {
				t.Errorf("ModuloOperation.Execute(%f, %f) returned error: %v", test.a, test.b, err)
			}
			if math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("ModuloOperation.Execute(%f, %f) = %f, expected %f", test.a, test.b, result, test.expected)
			}
		}
	}

	if op.Symbol() != "%" {
		t.Errorf("ModuloOperation.Symbol() = %s, expected %%", op.Symbol())
	}
}

func TestSqrtOperation(t *testing.T) {
	op := SqrtOperation{}
	
	tests := []struct {
		a, expected float64
		expectError bool
	}{
		{4, 2, false},
		{9, 3, false},
		{0, 0, false},
		{2, math.Sqrt(2), false},
		{-4, 0, true}, // Square root of negative number
	}

	for _, test := range tests {
		result, err := op.Execute(test.a, 0) // Second parameter ignored for unary operations
		if test.expectError {
			if err == nil {
				t.Errorf("SqrtOperation.Execute(%f) expected error but got none", test.a)
			}
		} else {
			if err != nil {
				t.Errorf("SqrtOperation.Execute(%f) returned error: %v", test.a, err)
			}
			if math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("SqrtOperation.Execute(%f) = %f, expected %f", test.a, result, test.expected)
			}
		}
	}

	if op.Symbol() != "√" {
		t.Errorf("SqrtOperation.Symbol() = %s, expected √", op.Symbol())
	}
}

func TestGetOperation(t *testing.T) {
	tests := []struct {
		symbol     string
		expectNil  bool
	}{
		{"+", false},
		{"-", false},
		{"*", false},
		{"/", false},
		{"^", false},
		{"%", false},
		{"√", false},
		{"unknown", true},
	}

	for _, test := range tests {
		op, err := GetOperation(test.symbol)
		if test.expectNil {
			if err == nil {
				t.Errorf("GetOperation(%s) expected error but got none", test.symbol)
			}
		} else {
			if err != nil {
				t.Errorf("GetOperation(%s) returned error: %v", test.symbol, err)
			}
			if op == nil {
				t.Errorf("GetOperation(%s) returned nil operation", test.symbol)
			}
		}
	}
}

func TestGetAllOperations(t *testing.T) {
	operations := GetAllOperations()
	expectedSymbols := []string{"+", "-", "*", "/", "^", "%", "√"}

	if len(operations) != len(expectedSymbols) {
		t.Errorf("GetAllOperations() returned %d operations, expected %d", len(operations), len(expectedSymbols))
	}

	for _, symbol := range expectedSymbols {
		if _, exists := operations[symbol]; !exists {
			t.Errorf("GetAllOperations() missing operation for symbol: %s", symbol)
		}
	}
}
