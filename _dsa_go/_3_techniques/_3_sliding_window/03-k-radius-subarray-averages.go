// https://leetcode.com/problems/k-radius-subarray-averages

package main

import "fmt"

func getAverages(nums []int, k int) []int {
	n := len(nums)

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	if k == 0 {
		return nums
	}

	windowSize := 2*k + 1
	if windowSize > n {
		return res
	}

	i, j := 0, 0
	center := k
	var sum int64 = 0

	for j < n {
		sum += int64(nums[j])

		if j-i+1 == windowSize {
			avg := sum / int64(windowSize)
			res[center] = int(avg)

			sum -= int64(nums[i])
			i++
			center++
		}

		j++
	}

	return res
}

func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	fmt.Println(getAverages(nums, k))
}