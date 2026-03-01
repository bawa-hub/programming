// https://leetcode.com/problems/rotate-image/

package main

import "fmt"

func rotateBrute(matrix [][]int) [][]int {
	n := len(matrix)

	rotated := make([][]int, n)
	for i := range rotated {
		rotated[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-i-1] = matrix[i][j]
		}
	}

	return rotated
}
// Time Complexity: O(N*N) to linearly iterate and put it into some other matrix.
// Space Complexity: O(N*N) to copy it into some other matrix.

func rotate(matrix [][]int) {
	n := len(matrix)

	// Step 1: Transpose
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: Reverse each row
	for i := 0; i < n; i++ {
		reverse(matrix[i])
	}
}

func reverse(row []int) {
	left, right := 0, len(row)-1
	for left < right {
		row[left], row[right] = row[right], row[left]
		left++
		right--
	}
}
// Time Complexity: O(N*N) + O(N*N).One O(N*N) for transposing the matrix and the other for reversing the matrix.
// Space Complexity: O(1).

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(arr)

	fmt.Println("Rotated Image:")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}