// https://leetcode.com/problems/valid-parentheses/

package main

import "fmt"

func isValid(s string) bool {

	stack := []rune{}

	for _, ch := range s {

		// opening brackets
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {

			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (ch == ')' && top == '(') ||
				(ch == ']' && top == '[') ||
				(ch == '}' && top == '{') {
				continue
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func isValidOptimized(s string) bool {

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for _, ch := range s {
		if open, ok := pairs[ch]; ok {

			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

func main() {

	s := "()[{}()]"

	if isValid(s) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}