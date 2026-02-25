// https://leetcode.com/problems/maximum-subarray/
// https://takeuforward.org/data-structure/kadanes-algorithm-maximum-subarray-sum-in-an-array/

package main

import (
	"fmt"
	"math"
)

func maxSubArrayBrute(nums []int) (int, []int) {
	n := len(nums)
	maxSum := math.MinInt
	subarray := []int{0, 0}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {

			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}

			if sum > maxSum {
				maxSum = sum
				subarray[0] = i
				subarray[1] = j
			}
		}
	}

	return maxSum, subarray
}
// Time Complexity: O(N^3)
// Space Complexity: O(1)

func maxSubArrayBetter(nums []int) (int, []int) {
	n := len(nums)
	maxSum := math.MinInt
	subarray := []int{0, 0}

	for i := 0; i < n; i++ {
		currSum := 0
		for j := i; j < n; j++ {
			currSum += nums[j]

			if currSum > maxSum {
				maxSum = currSum
				subarray[0] = i
				subarray[1] = j
			}
		}
	}

	return maxSum, subarray
}
// Time Complexity: O(N^2)
// Space Complexity: O(1)

// kadanes alog
// Kadane's algorithm runs one for loop over the array and at the beginning of each iteration, 
// if the current sum is negative, it will reset the current sum to zero. 
// This way, we ensure a one-pass and solve the problem in linear time.
func maxSubArrayKadane(nums []int) (int, []int) {
	maxSum := nums[0]
	currSum := 0

	start, end := 0, 0
	tempStart := 0

	for i := 0; i < len(nums); i++ {

		currSum += nums[i]

		if currSum > maxSum {
			maxSum = currSum
			start = tempStart
			end = i
		}

		if currSum < 0 {
			currSum = 0
			tempStart = i + 1
		}
	}

	return maxSum, []int{start, end}
}
// Time Complexity: O(N)
// Space Complexity:O(1)

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum, indices := maxSubArrayKadane(arr)

	fmt.Println("Maximum subarray sum:", maxSum)
	fmt.Println("Subarray:")
	for i := indices[0]; i <= indices[1]; i++ {
		fmt.Print(arr[i], " ")
	}
}