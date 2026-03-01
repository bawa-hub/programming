
package main

import (
	"fmt"
	"math"
)

func maxSubarraySum(arr []int) (int64, int, int) {
	var maxi int64 = math.MinInt64
	var sum int64 = 0

	start := 0
	ansStart, ansEnd := -1, -1

	for i := 0; i < len(arr); i++ {

		if sum == 0 {
			start = i
		}

		sum += int64(arr[i])

		if sum > maxi {
			maxi = sum
			ansStart = start
			ansEnd = i
		}

		if sum < 0 {
			sum = 0
		}
	}

	return maxi, ansStart, ansEnd
}
// Time Complexity: O(N), where N = size of the array.
// Reason: We are using a single loop running N times.

// Space Complexity: O(1) as we are not using any extra space.

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum, start, end := maxSubarraySum(arr)

	fmt.Println("The subarray is:")
	fmt.Print("[ ")
	for i := start; i <= end; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("]")

	fmt.Println("The maximum subarray sum is:", maxSum)
}