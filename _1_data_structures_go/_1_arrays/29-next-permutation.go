// https://leetcode.com/problems/next-permutation/

package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	n := len(nums)
	k := -1

	// Step 1: Find first decreasing element from right
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			k = i
			break
		}
	}

	// If no such element → reverse entire array
	if k == -1 {
		reverse(nums, 0, n-1)
		return
	}

	// Step 2: Find element just greater than nums[k]
	for l := n - 1; l > k; l-- {
		if nums[l] > nums[k] {
			nums[k], nums[l] = nums[l], nums[k]
			break
		}
	}

	// Step 3: Reverse the suffix
	reverse(nums, k+1, n-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // [1 3 2]
}