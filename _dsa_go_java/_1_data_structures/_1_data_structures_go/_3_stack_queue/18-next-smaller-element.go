// https://www.interviewbit.com/problems/nearest-smaller-element/

package main

import "fmt"

func prevSmaller(A []int) []int {

	n := len(A)
	res := make([]int, n)

	// initialize with -1
	for i := range res {
		res[i] = -1
	}

	stack := []int{}

	for i := 0; i < n; i++ {

		for len(stack) > 0 && stack[len(stack)-1] >= A[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res[i] = stack[len(stack)-1]
		}

		stack = append(stack, A[i])
	}

	return res
}

func main() {

	v := []int{4, 5, 2, 10, 8}

	ans := prevSmaller(v)

	for _, val := range ans {
		fmt.Print(val, " ")
	}
	fmt.Println()
}