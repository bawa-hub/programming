// https://leetcode.com/problems/sort-colors/

package main

import "fmt"

// brute force
func sortArray(arr []int) {
	n := len(arr)

	cnt0, cnt1, cnt2 := 0, 0, 0

	// Count
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			cnt0++
		} else if arr[i] == 1 {
			cnt1++
		} else {
			cnt2++
		}
	}

	// Replace 0s
	for i := 0; i < cnt0; i++ {
		arr[i] = 0
	}

	// Replace 1s
	for i := cnt0; i < cnt0+cnt1; i++ {
		arr[i] = 1
	}

	// Replace 2s
	for i := cnt0 + cnt1; i < n; i++ {
		arr[i] = 2
	}
}
// Time Complexity: O(N) + O(N), where N = size of the array. First O(N) for counting the number of 0’s, 1’s, 2’s, and second O(N) for placing them correctly in the original array.
// Space Complexity: O(1) as we are not using any extra space.

// Optimal Approach (Dutch National flag algorithm):
func sortArrayOptimized(arr []int) {
	low, mid := 0, 0
	high := len(arr) - 1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}
// Time Complexity: O(N), where N = size of the given array.
// Reason: We are using a single loop that can run at most N times.
// Space Complexity: O(1) as we are not using any extra space.

func main() {
	arr := []int{0, 2, 1, 2, 0, 1, 0}
	sortArrayOptimized(arr)
	fmt.Println(arr)
}
