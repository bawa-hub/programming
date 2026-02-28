// infix expression: <operand><operator><operand>. e.g. (A*B)/Q
// postfix expression: <operand><operand><operator>. e.g PQ-C/

package main

import (
	"fmt"
	"unicode"
)

/************ PRECEDENCE ************/

func prec(c rune) int {
	switch c {
	case '^':
		return 3
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	default:
		return -1
	}
}

/************ INFIX → POSTFIX ************/

func infixToPostfix(s string) string {

	stack := []rune{}
	result := ""

	for _, c := range s {

		// operand
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			result += string(c)

		} else if c == '(' {
			stack = append(stack, c)

		} else if c == ')' {

			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				result += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			// remove '('
			stack = stack[:len(stack)-1]

		} else { // operator

			for len(stack) > 0 &&
				prec(c) <= prec(stack[len(stack)-1]) {

				result += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, c)
		}
	}

	// remaining operators
	for len(stack) > 0 {
		result += string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return result
}
// Time Complexity: O(N)
// Space Complexity: O(N) for using the stack

/************ MAIN ************/

func main() {

	expr := "a+b*(c-d)"

	fmt.Println("Postfix expression:",
		infixToPostfix(expr))
}