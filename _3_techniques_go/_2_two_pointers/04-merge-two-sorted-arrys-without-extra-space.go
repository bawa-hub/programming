// https://leetcode.com/problems/merge-sorted-array/

package main

import (
	"fmt"
	"sort"
)

func mergeBrute(arr1 []int, arr2 []int) {
	n := len(arr1)
	m := len(arr2)

	arr3 := make([]int, 0, n+m)

	arr3 = append(arr3, arr1...)
	arr3 = append(arr3, arr2...)

	sort.Ints(arr3)

	copy(arr1, arr3[:n])
	copy(arr2, arr3[n:])
}
// Time complexity: O(n*log(n))+O(n)+O(n)
// Space Complexity: O(n)

func mergeNoExtraSpace(arr1 []int, arr2 []int) {

	n := len(arr1)
	m := len(arr2)

	for i := 0; i < n; i++ {

		if arr1[i] > arr2[0] {
			arr1[i], arr2[0] = arr2[0], arr1[i]
		}

		first := arr2[0]
		k := 1

		for k < m && arr2[k] < first {
			arr2[k-1] = arr2[k]
			k++
		}
		arr2[k-1] = first
	}
}
// Time complexity: O(n*m)
// Space Complexity: O(1)

// optimized
func merge(nums1 []int, m int, nums2 []int, n int) {

	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
// Time complexity: O(m+n)
// Space complexity: O(1)