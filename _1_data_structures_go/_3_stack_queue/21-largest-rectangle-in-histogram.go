// https://leetcode.com/problems/largest-rectangle-in-histogram/

package main

import (
	"fmt"
	"math"
)

func largestArea(arr []int) int {

	n := len(arr)
	maxArea := 0

	for i := 0; i < n; i++ {

		minHeight := math.MaxInt

		for j := i; j < n; j++ {

			if arr[j] < minHeight {
				minHeight = arr[j]
			}

			area := minHeight * (j - i + 1)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
// Time Complexity: O(N*N )
// Space Complexity: O(1)

// optimized
func largestRectangleArea(heights []int) int {

	n := len(heights)
	leftSmall := make([]int, n)
	rightSmall := make([]int, n)

	stack := []int{}

	// LEFT SMALLER
	for i := 0; i < n; i++ {

		for len(stack) > 0 &&
			heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			leftSmall[i] = 0
		} else {
			leftSmall[i] = stack[len(stack)-1] + 1
		}

		stack = append(stack, i)
	}

	// clear stack
	stack = []int{}

	// RIGHT SMALLER
	for i := n - 1; i >= 0; i-- {

		for len(stack) > 0 &&
			heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			rightSmall[i] = n - 1
		} else {
			rightSmall[i] = stack[len(stack)-1] - 1
		}

		stack = append(stack, i)
	}

	// calculate max area
	maxArea := 0

	for i := 0; i < n; i++ {
		width := rightSmall[i] - leftSmall[i] + 1
		area := heights[i] * width

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
// Time Complexity: O( N )
// Space Complexity: O(3N) where 3 is for the stack, left small array and a right small array


func largestRectangleAreaOptimized(histo []int) int {

	stack := []int{}
	maxArea := 0
	n := len(histo)

	for i := 0; i <= n; i++ {

		for len(stack) > 0 &&
			(i == n || histo[stack[len(stack)-1]] >= histo[i]) {

			height := histo[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}
// Time Complexity: O( N ) + O (N)
// Space Complexity: O(N)


func main() {

	arr := []int{2, 1, 5, 6, 2, 3}

	fmt.Println("Largest Area:", largestArea(arr))
}