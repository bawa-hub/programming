// https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/

package main

import "fmt"

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {

		rowSet := make(map[int]bool)
		colSet := make(map[int]bool)

		for j := 0; j < n; j++ {

			// Check row
			if rowSet[matrix[i][j]] {
				return false
			}
			rowSet[matrix[i][j]] = true

			// Check column
			if colSet[matrix[j][i]] {
				return false
			}
			colSet[matrix[j][i]] = true
		}
	}

	return true
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}

	fmt.Println("Is matrix valid?", checkValid(matrix))
}