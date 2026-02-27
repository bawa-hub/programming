// https://leetcode.com/problems/sliding-window-maximum/

package main

import (
	"fmt"
	"math"
)

func getMax(nums []int, l int, r int, arr *[]int) {

	maxi := math.MinInt

	for i := l; i <= r; i++ {
		if nums[i] > maxi {
			maxi = nums[i]
		}
	}

	*arr = append(*arr, maxi)
}

func maxSlidingWindow(nums []int, k int) []int {

	left := 0
	right := 0
	arr := []int{}

	// build first window
	for right < k-1 {
		right++
	}

	for right < len(nums) {
		getMax(nums, left, right, &arr)
		left++
		right++
	}

	return arr
}
// Time Complexity: O(N^2)
// Reason: One loop for traversing and another to findMax
// Space Complexity: O(K)
// Reason: No.of windows

func maxSlidingWindowOptimized(nums []int, k int) []int {

	deque := []int{} // stores indices
	ans := []int{}

	for i := 0; i < len(nums); i++ {

		// remove elements out of window
		if len(deque) > 0 && deque[0] == i-k {
			deque = deque[1:]
		}

		// maintain decreasing order
		for len(deque) > 0 &&
			nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		// window formed
		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}

	return ans
}
// Time Complexity: O(N)
// Space Complexity: O(K)

func main() {

	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	result := maxSlidingWindow(nums, k)

	fmt.Println(result)
}