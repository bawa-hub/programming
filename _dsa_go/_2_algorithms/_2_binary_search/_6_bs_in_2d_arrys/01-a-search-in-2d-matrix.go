// https://leetcode.com/problems/search-a-2d-matrix/

// naive approach (by linear search)
// Time complexity: O(m*n)
// Space complexity: O(1)

package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	lo := 0
	hi := rows*cols - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		// Convert 1D index to 2D
		row := mid / cols
		col := mid % cols

		value := matrix[row][col]

		if value == target {
			return true
		} else if value < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return false
}
// Time complexity: O(log(m*n))
// Space complexity: O(1)

func main() {
	matrix := [][]int{
		{1, 3, 5},
		{7, 9, 11},
		{13, 15, 17},
	}

	fmt.Println(searchMatrix(matrix, 9))  // true
	fmt.Println(searchMatrix(matrix, 8))  // false
}