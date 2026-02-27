// https://leetcode.com/problems/next-greater-element-ii/

package main

import "fmt"

func nextGreaterElements(nums []int) []int {

	n := len(nums)
	nge := make([]int, n)

	// initialize with -1
	for i := range nge {
		nge[i] = -1
	}

	stack := []int{}

	// traverse twice (circular array)
	for i := 2*n - 1; i >= 0; i-- {

		curr := nums[i%n]

		for len(stack) > 0 && stack[len(stack)-1] <= curr {
			stack = stack[:len(stack)-1]
		}

		if i < n {
			if len(stack) > 0 {
				nge[i] = stack[len(stack)-1]
			}
		}

		stack = append(stack, curr)
	}

	return nge
}

func main() {

	v := []int{5, 7, 1, 2, 6, 0}

	res := nextGreaterElements(v)

	fmt.Println("Next greater elements:")
	for _, val := range res {
		fmt.Print(val, " ")
	}
}