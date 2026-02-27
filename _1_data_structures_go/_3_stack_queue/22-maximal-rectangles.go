// https://leetcode.com/problems/maximal-rectangle/

package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	cols := len(matrix[0])
	height := make([]int, cols)
	maxRec := 0

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < cols; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j]++
			}
		}

		area := largestRectangleArea(height)
		if area > maxRec {
			maxRec = area
		}
	}

	return maxRec
}

/************ HISTOGRAM SOLUTION ************/

func largestRectangleArea(height []int) int {

	stack := []int{}
	maxArea := 0

	// sentinel
	h := append(height, 0)

	for i := 0; i < len(h); i++ {

		for len(stack) > 0 &&
			h[i] < h[stack[len(stack)-1]] {

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			heightVal := h[top]

			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			area := heightVal * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}

/************ TEST ************/

func main() {

	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	fmt.Println("Max Rectangle:", maximalRectangle(matrix))
}