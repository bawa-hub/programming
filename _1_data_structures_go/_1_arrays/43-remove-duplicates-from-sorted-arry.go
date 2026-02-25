// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(arr []int) int {
	// Create a set using map
	set := make(map[int]struct{})

	// Insert elements into set
	for _, v := range arr {
		set[v] = struct{}{}
	}

	// Optional: Since Go map is unordered,
	// we sort keys to mimic C++ set behavior
	keys := make([]int, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Put unique elements back into original array
	for i, v := range keys {
		arr[i] = v
	}

	return len(set)
}
// Time complexity: O(n*log(n))+O(n)
// Space Complexity: O(n)

func removeDuplicatesOptimized(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	i := 0

	for j := 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
	}

	return i + 1
}
// Time complexity: O(n)
// Space Complexity: O(1)

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 4, 5}

	k := removeDuplicates(arr)

	fmt.Println("Unique count:", k)
	fmt.Println("Array after removing duplicates:", arr[:k])
}