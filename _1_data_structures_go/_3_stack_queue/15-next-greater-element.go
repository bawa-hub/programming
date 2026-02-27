// https://leetcode.com/problems/next-greater-element-i/

package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	mp := make(map[int]int)
	stack := []int{}

	// process nums2 from right to left
	for i := len(nums2) - 1; i >= 0; i-- {

		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			mp[nums2[i]] = stack[len(stack)-1]
		} else {
			mp[nums2[i]] = -1
		}

		stack = append(stack, nums2[i])
	}

	// build answer
	ans := []int{}
	for _, v := range nums1 {
		ans = append(ans, mp[v])
	}

	return ans
}

func main() {

	nums1 := []int{5, 1, 6}
	nums2 := []int{5, 7, 1, 2, 6, 0}

	res := nextGreaterElement(nums1, nums2)

	fmt.Println("Next greater elements:")
	for _, v := range res {
		fmt.Print(v, " ")
	}
}