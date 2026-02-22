package main

import "fmt"

func binarySearchRecursive(arr []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2

	if arr[mid] == target {
		return mid
	} else if target < arr[mid] {
		return binarySearchRecursive(arr, start, mid-1, target)
	} else {
		return binarySearchRecursive(arr, mid+1, end, target)
	}
}
// Time complexity: O(log n)
// Space complexity: O(logn) for auxiliary space

func binarySearchIterative(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
// Time complexity: O(log n)
// Space complexity: O(1)

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	index := binarySearchIterative(arr, 7)
	fmt.Println(index) // 3
}
