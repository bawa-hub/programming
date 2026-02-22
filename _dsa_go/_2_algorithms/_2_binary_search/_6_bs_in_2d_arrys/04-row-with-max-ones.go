// https://www.codingninjas.com/studio/problems/row-of-a-matrix-with-maximum-ones_982768

package main

import "fmt"

// brute force
func rowWithMax1s(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return -1
	}
	m := len(matrix[0])

	cntMax := 0
	index := -1

	for i := 0; i < n; i++ {
		cntOnes := 0

		for j := 0; j < m; j++ {
			cntOnes += matrix[i][j]
		}

		if cntOnes > cntMax {
			cntMax = cntOnes
			index = i
		}
	}

	return index
}
// Time Complexity: O(n X m), where n = given row number, m = given column number.
// Reason: We are using nested loops running for n and m times respectively.
// Space Complexity: O(1) as we are not using any extra space.

// binary search
func lowerBound(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	ans := len(arr)

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] >= x {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func rowWithMax1sOptimized(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return -1
	}

	m := len(matrix[0])
	cntMax := 0
	index := -1

	for i := 0; i < n; i++ {
		firstOne := lowerBound(matrix[i], 1)
		cntOnes := m - firstOne

		if cntOnes > cntMax {
			cntMax = cntOnes
			index = i
		}
	}

	return index
}
// Time Complexity: O(n X logm), where n = given row number, m = given column number.
// Reason: We are using a loop running for n times to traverse the rows. Then we are applying binary search on each row with m columns.
// Space Complexity: O(1) as we are not using any extra space.

func main() {
	matrix := [][]int{
		{0, 1, 1, 1},
		{0, 0, 1, 1},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	}

	fmt.Println(rowWithMax1s(matrix)) // Expected: 2
}