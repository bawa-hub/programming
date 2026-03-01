// https://leetcode.com/problems/remove-outermost-parentheses/

package main

import "fmt"

func removeOuterParentheses(s string) string {
	stack := []rune{}
	ans := []rune{}

	for _, c := range s {

		if len(stack) > 0 {
			ans = append(ans, c)
		}

		if c == '(' {
			stack = append(stack, c)
		} else {
			// pop
			stack = stack[:len(stack)-1]

			// if stack becomes empty, remove last added char
			if len(stack) == 0 {
				ans = ans[:len(ans)-1]
			}
		}
	}

	return string(ans)
}
// Time complexity:
// O(N)
// Space complexity:
// O(N)

func removeOuterParenthesesWithoutStack(s string) string {
	count := 0
	result := make([]rune, 0)

	for _, c := range s {

		if c == '(' {
			if count > 0 {
				result = append(result, c)
			}
			count++
		} else { // ')'
			count--
			if count > 0 {
				result = append(result, c)
			}
		}
	}

	return string(result)
}
// TC: O(n)
// SC: O(1)

func main() {
	s := "(()())(())"
	fmt.Println(removeOuterParentheses(s)) // "()()()"
}