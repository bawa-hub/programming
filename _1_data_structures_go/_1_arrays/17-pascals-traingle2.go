// https://leetcode.com/problems/pascals-triangle-ii/description/

package main

import "fmt"

func getRow(rowIndex int) []int {
	rows := generate(rowIndex + 1)
	return rows[rowIndex]
}

func generate(numRows int) [][]int {
	r := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		r[i] = make([]int, i+1)

		r[i][0] = 1
		r[i][i] = 1

		for j := 1; j < i; j++ {
			r[i][j] = r[i-1][j-1] + r[i-1][j]
		}
	}

	return r
}

func main() {
	fmt.Println(getRow(4)) // Output: [1 4 6 4 1]
}