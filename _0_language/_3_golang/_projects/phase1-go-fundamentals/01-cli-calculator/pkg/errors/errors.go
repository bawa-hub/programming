package errors

import "fmt"

// CalculatorError represents a calculator-specific error
type CalculatorError struct {
	Type    string
	Message string
	Value   interface{}
}

func (e *CalculatorError) Error() string {
	if e.Value != nil {
		return fmt.Sprintf("%s: %s (value: %v)", e.Type, e.Message, e.Value)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// NewCalculatorError creates a new calculator error
func NewCalculatorError(errorType, message string, value interface{}) *CalculatorError {
	return &CalculatorError{
		Type:    errorType,
		Message: message,
		Value:   value,
	}
}

// Predefined error types
var (
	ErrInvalidInput     = "InvalidInput"
	ErrDivisionByZero   = "DivisionByZero"
	ErrInvalidOperation = "InvalidOperation"
	ErrParseError       = "ParseError"
	ErrMathError        = "MathError"
)

// NewInvalidInputError creates an invalid input error
func NewInvalidInputError(message string, value interface{}) *CalculatorError {
	return NewCalculatorError(ErrInvalidInput, message, value)
}

// NewDivisionByZeroError creates a division by zero error
func NewDivisionByZeroError() *CalculatorError {
	return NewCalculatorError(ErrDivisionByZero, "cannot divide by zero", nil)
}

// NewInvalidOperationError creates an invalid operation error
func NewInvalidOperationError(operation string) *CalculatorError {
	return NewCalculatorError(ErrInvalidOperation, "unknown operation", operation)
}

// NewParseError creates a parse error
func NewParseError(message string, value interface{}) *CalculatorError {
	return NewCalculatorError(ErrParseError, message, value)
}

// NewMathError creates a math error
func NewMathError(message string, value interface{}) *CalculatorError {
	return NewCalculatorError(ErrMathError, message, value)
}
