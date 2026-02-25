// https://leetcode.com/problems/sum-in-a-matrix/description/

package main

import (
	"fmt"
	"sort"
)

func matrixSum(nums [][]int) int {
	for i := range nums {
		sort.Ints(nums[i])
	}

	sum := 0
	cols := len(nums[0])

	for j := cols - 1; j >= 0; j-- {
		maxVal := 0

		for i := 0; i < len(nums); i++ {
			if nums[i][j] > maxVal {
				maxVal = nums[i][j]
			}
		}

		sum += maxVal
	}

	return sum
}

func main() {
	nums := [][]int{
		{7, 2, 1},
		{6, 4, 2},
		{6, 5, 3},
	}

	fmt.Println(matrixSum(nums))
}

// Overall:

// 👉 O(n * m log m)

// Space:
// 👉 O(1) extra (optimized version)