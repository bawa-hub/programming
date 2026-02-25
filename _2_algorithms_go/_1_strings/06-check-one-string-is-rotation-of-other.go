package main

import "fmt"

func rotateCheck(A, B string, rotation int) bool {
	n := len(A)

	for i := 0; i < n; i++ {
		if A[i] != B[(i+rotation)%n] {
			return false
		}
	}
	return true
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if rotateCheck(s, goal, i) {
			return true
		}
	}

	return false
}
// Time Complexity: O(N^2), where N is the length of string(s). For each rotation string(goal), we check up to N
// elements in string(s) and string(goal).

// Space Complexity: O(1), Constant space. We only use pointers to elements of s and goal.

func main() {
	fmt.Println(rotateString("abcde", "cdeab")) // true
	fmt.Println(rotateString("abcde", "abced")) // false
}