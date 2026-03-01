// https://leetcode.com/problems/spiral-matrix/

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	R := len(matrix)
	C := len(matrix[0])

	ans := []int{}

	top, left := 0, 0
	bottom, right := R-1, C-1

	for top <= bottom && left <= right {

		// Traverse from left to right
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		// Traverse top to bottom
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		// Traverse right to left
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
		}

		// Traverse bottom to top
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}

	return ans
}
// Time Complexity: O(R x C)
// Reason: We are printing every element of the matrix so the time complexity is O(R x C) where R and C are rows and columns of the matrix.
// Space Complexity: O(1)


func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	result := spiralOrder(matrix)
	fmt.Println(result)
}