// https://leetcode.com/problems/maximum-product-subarray/

package main

import (
	"fmt"
	"math"
)

func maxProductSubArrayBrute(nums []int) int {
	n := len(nums)
	result := math.MinInt

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			prod := 1
			for k := i; k <= j; k++ {
				prod *= nums[k]
			}
			if prod > result {
				result = prod
			}
		}
	}
	return result
}
// TC: O(N^3)
// SC: O(1)

func maxProductSubArrayBetter(nums []int) int {
	n := len(nums)
	result := nums[0]

	for i := 0; i < n-1; i++ {
		p := nums[i]

		for j := i + 1; j < n; j++ {
			if p > result {
				result = p
			}
			p *= nums[j]
		}

		if p > result {
			result = p
		}
	}
	return result
}
// TC: O(N^2)
// SC: O(1)

func maxProductSubArrayTwoPass(nums []int) int {
	maxLeft := nums[0]
	maxRight := nums[0]
	prod := 1
	zeroPresent := false

	// Left to right
	for _, v := range nums {
		prod *= v
		if v == 0 {
			zeroPresent = true
			prod = 1
			continue
		}
		if prod > maxLeft {
			maxLeft = prod
		}
	}

	prod = 1

	// Right to left
	for i := len(nums) - 1; i >= 0; i-- {
		prod *= nums[i]
		if nums[i] == 0 {
			zeroPresent = true
			prod = 1
			continue
		}
		if prod > maxRight {
			maxRight = prod
		}
	}

	if zeroPresent {
		return max(max(maxLeft, maxRight), 0)
	}
	return max(maxLeft, maxRight)
}
// TC: O(N)
// SC: O(1)

func maxProductSubArrayKadane(nums []int) int {
	prod1 := nums[0] // max product
	prod2 := nums[0] // min product (important for negatives)
	result := nums[0]

	for i := 1; i < len(nums); i++ {

		temp := max(nums[i],
			max(prod1*nums[i], prod2*nums[i]))

		prod2 = min(nums[i],
			min(prod1*nums[i], prod2*nums[i]))

		prod1 = temp

		if prod1 > result {
			result = prod1
		}
	}

	return result
}
// TC: O(N)
// SC: O(1)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, -3, 0, -4, -5}

	fmt.Println("Brute:", maxProductSubArrayBrute(nums))
	fmt.Println("Better:", maxProductSubArrayBetter(nums))
	fmt.Println("TwoPass:", maxProductSubArrayTwoPass(nums))
	fmt.Println("Kadane:", maxProductSubArrayKadane(nums))
}