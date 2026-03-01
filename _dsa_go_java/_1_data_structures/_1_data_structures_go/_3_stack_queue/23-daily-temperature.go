// https://leetcode.com/problems/daily-temperatures/

package main

import "fmt"

type Pair struct {
	temp  int
	index int
}

func dailyTemperatures(temperatures []int) []int {

	stack := []Pair{}
	n := len(temperatures)
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {

		for len(stack) > 0 &&
			stack[len(stack)-1].temp <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res[i] = stack[len(stack)-1].index - i
		} else {
			res[i] = 0
		}

		stack = append(stack, Pair{temperatures[i], i})
	}

	return res
}

func main() {

	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}

	fmt.Println(dailyTemperatures(temps))
}