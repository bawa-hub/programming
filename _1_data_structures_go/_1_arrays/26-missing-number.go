// https://leetcode.com/problems/missing-number/

package main

import "fmt"

func missingNumber(nums []int) int {
	size := len(nums)

	total := size * (size + 1) / 2

	for i := 0; i < size; i++ {
		total -= nums[i]
	}

	return total
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums)) // Output: 2
}