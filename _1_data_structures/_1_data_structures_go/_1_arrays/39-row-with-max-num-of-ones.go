// https://practice.geeksforgeeks.org/problems/row-with-max-1s0023/1

package main

import "fmt"

func rowWithMax1s(arr [][]int, n, m int) int {
	r := 0
	c := m - 1
	maxRowIndex := -1

	// Start from top-right corner
	for r < n && c >= 0 {
		if arr[r][c] == 1 {
			maxRowIndex = r
			c-- // move left
		} else {
			r++ // move down
		}
	}

	return maxRowIndex
}

func main() {
	arr := [][]int{
		{0, 0, 0, 1},
		{0, 1, 1, 1},
		{0, 0, 1, 1},
		{1, 1, 1, 1},
	}

	n := len(arr)
	m := len(arr[0])

	fmt.Println(rowWithMax1s(arr, n, m)) // Output: 3
}