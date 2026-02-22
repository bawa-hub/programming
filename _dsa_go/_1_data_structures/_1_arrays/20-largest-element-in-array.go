// https://practice.geeksforgeeks.org/problems/largest-element-in-array4009/0?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=largest-element-in-array
// https://takeuforward.org/data-structure/find-the-largest-element-in-an-array/

package main

import (
	"fmt"
	"sort"
)

// brute force
func findLargestElement1(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}
// Time Complexity: O(N*log(N))
// Space Complexity: O(n)

func findLargestElement2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxVal := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}

	return maxVal
}
// Time Complexity: O(N)
// Space Complexity: O(n)

func main() {
	arr1 := []int{2, 5, 1, 3, 0}
	fmt.Println("Largest element:", findLargestElement1(arr1))

	arr2 := []int{8, 10, 5, 7, 9}
	fmt.Println("Largest element:", findLargestElement2(arr2))
}
