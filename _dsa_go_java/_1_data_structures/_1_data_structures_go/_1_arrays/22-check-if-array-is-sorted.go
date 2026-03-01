// https://practice.geeksforgeeks.org/problems/check-if-an-array-is-sorted0701/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=check-if-an-array-is-sorted

package main

import "fmt"

func isSortedBrute(arr []int) bool {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[i] {
				return false
			}
		}
	}
	return true
}
// Time Complexity: O(N^2)
// Space Complexity: O(1)

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
// Time Complexity: O(N)
// Space Complexity: O(1)


func main() {
	arr := []int{1, 2, 3, 4, 5}

	if isSorted(arr) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}