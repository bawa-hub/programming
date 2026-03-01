// https://leetcode.com/problems/asteroid-collision/

package main

import (
	"fmt"
	"math"
)

func asteroidCollision(ast []int) []int {

	stack := []int{}

	for _, a := range ast {

		if a > 0 || len(stack) == 0 {
			stack = append(stack, a)
		} else {

			// current asteroid moving left
			destroyed := false

			for len(stack) > 0 &&
				stack[len(stack)-1] > 0 &&
				stack[len(stack)-1] < int(math.Abs(float64(a))) {

				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 &&
				stack[len(stack)-1] == int(math.Abs(float64(a))) {

				// both destroy
				stack = stack[:len(stack)-1]
				destroyed = true
			} else if len(stack) == 0 ||
				stack[len(stack)-1] < 0 {

				stack = append(stack, a)
			}

			if destroyed {
				continue
			}
		}
	}

	return stack
}

func main() {

	ast := []int{5, 10, -5}
	fmt.Println(asteroidCollision(ast)) // [5 10]

	ast2 := []int{8, -8}
	fmt.Println(asteroidCollision(ast2)) // []
}