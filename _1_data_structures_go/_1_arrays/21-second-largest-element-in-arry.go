// https://practice.geeksforgeeks.org/problems/second-largest3735/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=second-largest
// https://takeuforward.org/data-structure/find-second-smallest-and-second-largest-element-in-an-array/


package main

import (
	"fmt"
	"math"
)

func getElements(arr []int) {
	n := len(arr)

	if n < 2 {
		fmt.Println(-1, -1)
		return
	}

	small, secondSmall := math.MaxInt, math.MaxInt
	large, secondLarge := math.MinInt, math.MinInt

	// First pass: find smallest & largest
	for _, v := range arr {
		if v < small {
			small = v
		}
		if v > large {
			large = v
		}
	}

	// Second pass: find second smallest & second largest
	for _, v := range arr {
		if v < secondSmall && v != small {
			secondSmall = v
		}
		if v > secondLarge && v != large {
			secondLarge = v
		}
	}

	fmt.Println("Second smallest is", secondSmall)
	fmt.Println("Second largest is", secondLarge)
}
// Time Complexity: O(2N), We do two linear traversals in our array
// Space Complexity: O(1)

// best
func secondSmallest(arr []int) int {
	n := len(arr)
	if n < 2 {
		return -1
	}

	small := math.MaxInt
	secondSmall := math.MaxInt

	for _, v := range arr {
		if v < small {
			secondSmall = small
			small = v
		} else if v < secondSmall && v != small {
			secondSmall = v
		}
	}

	if secondSmall == math.MaxInt {
		return -1
	}
	return secondSmall
}

func secondLargest(arr []int) int {
	n := len(arr)
	if n < 2 {
		return -1
	}

	large := math.MinInt
	secondLarge := math.MinInt

	for _, v := range arr {
		if v > large {
			secondLarge = large
			large = v
		} else if v > secondLarge && v != large {
			secondLarge = v
		}
	}

	if secondLarge == math.MinInt {
		return -1
	}
	return secondLarge
}

// Time Complexity: O(N), Single-pass solution
// Space Complexity: O(1)

func main() {
	arr := []int{1, 2, 4, 7, 7, 5}

	fmt.Println("Second smallest is", secondSmallest(arr))
	fmt.Println("Second largest is", secondLargest(arr))
}