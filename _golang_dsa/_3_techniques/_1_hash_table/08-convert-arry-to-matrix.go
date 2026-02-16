// https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/description/

package main

func findMatrix(nums []int) [][]int {
	mp := make(map[int]int)

	// Count frequency
	for _, num := range nums {
		mp[num]++
	}

	var res [][]int

	for len(mp) > 0 {
		var temp []int
		var keysToDelete []int

		for key, val := range mp {
			temp = append(temp, key)
			mp[key]--

			if val-1 == 0 {
				keysToDelete = append(keysToDelete, key)
			}
		}

		// Delete keys with 0 frequency
		for _, key := range keysToDelete {
			delete(mp, key)
		}

		res = append(res, temp)
	}

	return res
}
