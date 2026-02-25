// https://leetcode.com/problems/pascals-triangle/

package main

import "fmt"

func generate(numRows int) [][]int {
	r := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		// Create row of size i+1
		r[i] = make([]int, i+1)

		// First and last elements are always 1
		r[i][0] = 1
		r[i][i] = 1

		// Fill middle elements
		for j := 1; j < i; j++ {
			r[i][j] = r[i-1][j-1] + r[i-1][j]
		}
	}

	return r
}

func main() {
	numRows := 5
	result := generate(numRows)

	for _, row := range result {
		fmt.Println(row)
	}
}