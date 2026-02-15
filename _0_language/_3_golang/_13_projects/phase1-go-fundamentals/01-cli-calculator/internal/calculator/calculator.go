package calculator

import (
	"cli-calculator/pkg/errors"
	"cli-calculator/pkg/operations"
	"fmt"
	"strconv"
	"strings"
)

// Calculator represents the main calculator
type Calculator struct {
	history []string
}

// New creates a new calculator instance
func New() *Calculator {
	return &Calculator{
		history: make([]string, 0),
	}
}

// Calculate performs a simple two-operand calculation
func (c *Calculator) Calculate(a, b float64, operation string) (float64, error) {
	op, err := operations.GetOperation(operation)
	if err != nil {
		return 0, errors.NewInvalidOperationError(operation)
	}

	result, err := op.Execute(a, b)
	if err != nil {
		return 0, errors.NewMathError(err.Error(), fmt.Sprintf("%f %s %f", a, operation, b))
	}

	// Add to history
	historyEntry := fmt.Sprintf("%.2f %s %.2f = %.2f", a, operation, b, result)
	c.history = append(c.history, historyEntry)

	return result, nil
}

// CalculateUnary performs a unary operation (like square root)
func (c *Calculator) CalculateUnary(a float64, operation string) (float64, error) {
	op, err := operations.GetOperation(operation)
	if err != nil {
		return 0, errors.NewInvalidOperationError(operation)
	}

	result, err := op.Execute(a, 0) // For unary operations, second operand is ignored
	if err != nil {
		return 0, errors.NewMathError(err.Error(), fmt.Sprintf("%s %.2f", operation, a))
	}

	// Add to history
	historyEntry := fmt.Sprintf("%s %.2f = %.2f", operation, a, result)
	c.history = append(c.history, historyEntry)

	return result, nil
}

// ParseAndCalculate parses an expression and calculates the result
func (c *Calculator) ParseAndCalculate(expression string) (float64, error) {
	expression = strings.TrimSpace(expression)
	if expression == "" {
		return 0, errors.NewInvalidInputError("empty expression", expression)
	}

	// Handle unary operations (like √25)
	if strings.HasPrefix(expression, "√") {
		// Handle Unicode square root symbol
		numberStr := strings.TrimSpace(strings.TrimPrefix(expression, "√"))
		number, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			return 0, errors.NewParseError("invalid number for square root", numberStr)
		}
		return c.CalculateUnary(number, "√")
	}

	// Find the operation in the expression
	operation, err := c.findOperation(expression)
	if err != nil {
		return 0, err
	}

	// Split the expression by the operation
	parts := strings.Split(expression, operation)
	if len(parts) != 2 {
		return 0, errors.NewParseError("invalid expression format", expression)
	}

	// Parse the operands
	leftStr := strings.TrimSpace(parts[0])
	rightStr := strings.TrimSpace(parts[1])

	left, err := strconv.ParseFloat(leftStr, 64)
	if err != nil {
		return 0, errors.NewParseError("invalid left operand", leftStr)
	}

	right, err := strconv.ParseFloat(rightStr, 64)
	if err != nil {
		return 0, errors.NewParseError("invalid right operand", rightStr)
	}

	return c.Calculate(left, right, operation)
}

// findOperation finds the operation in the expression
func (c *Calculator) findOperation(expression string) (string, error) {
	// Order matters - check for multi-character operations first
	operations := []string{"**", "^", "*", "/", "%", "+", "-"}
	
	for _, op := range operations {
		if strings.Contains(expression, op) {
			return op, nil
		}
	}

	return "", errors.NewParseError("no valid operation found", expression)
}

// GetHistory returns the calculation history
func (c *Calculator) GetHistory() []string {
	return c.history
}

// ClearHistory clears the calculation history
func (c *Calculator) ClearHistory() {
	c.history = make([]string, 0)
}

// GetAvailableOperations returns all available operations
func (c *Calculator) GetAvailableOperations() map[string]operations.Operation {
	return operations.GetAllOperations()
}
