// https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values

package main

import "fmt"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	i, j := 0, 0
	res := [][]int{}

	for i < len(nums1) && j < len(nums2) {

		id1 := nums1[i][0]
		id2 := nums2[j][0]

		if id1 == id2 {
			res = append(res, []int{
				id1,
				nums1[i][1] + nums2[j][1],
			})
			i++
			j++
		} else if id1 < id2 {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	// remaining elements
	for i < len(nums1) {
		res = append(res, nums1[i])
		i++
	}

	for j < len(nums2) {
		res = append(res, nums2[j])
		j++
	}

	return res
}

func main() {
	nums1 := [][]int{{1, 2}, {2, 3}, {4, 5}}
	nums2 := [][]int{{1, 4}, {3, 2}, {4, 1}}

	result := mergeArrays(nums1, nums2)

	fmt.Println(result)
}