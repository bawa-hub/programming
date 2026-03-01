// https://leetcode.com/problems/majority-element/

// moore voting algorithm

package main

import "fmt"

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
    // TC: O(n)
    // SC: O(1)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums)) // Output: 2
}