// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/

package main

import "fmt"

func check(nums []int) bool {
	count := 0
	n := len(nums)

	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			count++
		}
	}

	// Check circular condition
	if nums[0] < nums[n-1] {
		count++
	}

	return count <= 1
}

func main() {
	fmt.Println(check([]int{3, 4, 5, 1, 2})) // true
	fmt.Println(check([]int{2, 1, 3, 4}))    // false
	fmt.Println(check([]int{1, 2, 3}))       // true
}