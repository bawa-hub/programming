package operations

import (
	"fmt"
	"math"
)

// Operation represents a mathematical operation
type Operation interface {
	Execute(a, b float64) (float64, error)
	Symbol() string
}

// AddOperation represents addition
type AddOperation struct{}

func (op AddOperation) Execute(a, b float64) (float64, error) {
	return a + b, nil
}

func (op AddOperation) Symbol() string {
	return "+"
}

// SubtractOperation represents subtraction
type SubtractOperation struct{}

func (op SubtractOperation) Execute(a, b float64) (float64, error) {
	return a - b, nil
}

func (op SubtractOperation) Symbol() string {
	return "-"
}

// MultiplyOperation represents multiplication
type MultiplyOperation struct{}

func (op MultiplyOperation) Execute(a, b float64) (float64, error) {
	return a * b, nil
}

func (op MultiplyOperation) Symbol() string {
	return "*"
}

// DivideOperation represents division
type DivideOperation struct{}

func (op DivideOperation) Execute(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func (op DivideOperation) Symbol() string {
	return "/"
}

// PowerOperation represents exponentiation
type PowerOperation struct{}

func (op PowerOperation) Execute(a, b float64) (float64, error) {
	result := math.Pow(a, b)
	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, fmt.Errorf("invalid power operation: %f^%f", a, b)
	}
	return result, nil
}

func (op PowerOperation) Symbol() string {
	return "^"
}

// ModuloOperation represents modulo operation
type ModuloOperation struct{}

func (op ModuloOperation) Execute(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("modulo by zero")
	}
	return math.Mod(a, b), nil
}

func (op ModuloOperation) Symbol() string {
	return "%"
}

// SqrtOperation represents square root (unary operation)
type SqrtOperation struct{}

func (op SqrtOperation) Execute(a, b float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root of negative number")
	}
	return math.Sqrt(a), nil
}

func (op SqrtOperation) Symbol() string {
	return "√"
}

// GetOperation returns an operation by symbol
func GetOperation(symbol string) (Operation, error) {
	switch symbol {
	case "+":
		return AddOperation{}, nil
	case "-":
		return SubtractOperation{}, nil
	case "*":
		return MultiplyOperation{}, nil
	case "/":
		return DivideOperation{}, nil
	case "^":
		return PowerOperation{}, nil
	case "%":
		return ModuloOperation{}, nil
	case "√":
		return SqrtOperation{}, nil
	default:
		return nil, fmt.Errorf("unknown operation: %s", symbol)
	}
}

// GetAllOperations returns all available operations
func GetAllOperations() map[string]Operation {
	return map[string]Operation{
		"+": AddOperation{},
		"-": SubtractOperation{},
		"*": MultiplyOperation{},
		"/": DivideOperation{},
		"^": PowerOperation{},
		"%": ModuloOperation{},
		"√": SqrtOperation{},
	}
}
