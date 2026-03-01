// https://leetcode.com/problems/trapping-rain-water/

package main

import "fmt"

// brute force
func trap(arr []int) int {

	n := len(arr)
	waterTrapped := 0

	for i := 0; i < n; i++ {

		leftMax := 0
		rightMax := 0

		// find left max
		for j := i; j >= 0; j-- {
			if arr[j] > leftMax {
				leftMax = arr[j]
			}
		}

		// find right max
		for j := i; j < n; j++ {
			if arr[j] > rightMax {
				rightMax = arr[j]
			}
		}

		waterTrapped += min(leftMax, rightMax) - arr[i]
	}

	return waterTrapped
}
// Time Complexity: O(N*N) as for each index we are calculating leftMax and rightMax so it is a nested loop.
// Space Complexity: O(1).

// better approach
func trapBetter(arr []int) int {

	n := len(arr)
	if n == 0 {
		return 0
	}

	prefix := make([]int, n)
	suffix := make([]int, n)

	// prefix max
	prefix[0] = arr[0]
	for i := 1; i < n; i++ {
		prefix[i] = max(prefix[i-1], arr[i])
	}

	// suffix max
	suffix[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], arr[i])
	}

	waterTrapped := 0

	for i := 0; i < n; i++ {
		waterTrapped += min(prefix[i], suffix[i]) - arr[i]
	}

	return waterTrapped
}
// Time Complexity: O(3*N) as we are traversing through the array only once. And O(2*N) for computing prefix and suffix array.
// Space Complexity: O(N)+O(N) for prefix and suffix arrays

// two pointer
func trapOptimized(height []int) int {

	n := len(height)
	left, right := 0, n-1

	maxLeft := 0
	maxRight := 0
	res := 0

	for left <= right {

		if height[left] <= height[right] {

			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++

		} else {

			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}

	return res
}
// Time Complexity: O(N) because we are using 2 pointer approach.
// Space Complexity: O(1) because we are not using anything extra.

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println("Water trapped:", trap(arr))
}