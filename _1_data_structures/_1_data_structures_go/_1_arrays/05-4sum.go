// https://leetcode.com/problems/4sum/

package main

import (
	"fmt"
	"sort"
)

func fourSumBrute(nums []int, target int) [][]int {
	n := len(nums)
	set := make(map[[4]int]bool)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {

					sum := int64(nums[i]) + int64(nums[j]) +
						int64(nums[k]) + int64(nums[l])

					if sum == int64(target) {
						temp := []int{nums[i], nums[j], nums[k], nums[l]}
						sort.Ints(temp)
						key := [4]int{temp[0], temp[1], temp[2], temp[3]}
						set[key] = true
					}
				}
			}
		}
	}

	var ans [][]int
	for k := range set {
		ans = append(ans, []int{k[0], k[1], k[2], k[3]})
	}
	return ans
}
// Time Complexity: O(N4), where N = size of the array.
// Reason: Here, we are mainly using 4 nested loops. But we not considering the time complexity of sorting as we are just sorting 4 elements every time.
// Space Complexity: O(2 * no. of the quadruplets) as we are using a set data structure and a list to store the quads.

// better approach (Using 3 loops and set data structure)
func fourSumBetter(nums []int, target int) [][]int {
	n := len(nums)
	set := make(map[[4]int]bool)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			hashset := make(map[int]bool)

			for k := j + 1; k < n; k++ {

				sum := int64(nums[i]) + int64(nums[j]) + int64(nums[k])
				fourth := int64(target) - sum

				if hashset[int(fourth)] {
					temp := []int{nums[i], nums[j], nums[k], int(fourth)}
					sort.Ints(temp)
					key := [4]int{temp[0], temp[1], temp[2], temp[3]}
					set[key] = true
				}

				hashset[nums[k]] = true
			}
		}
	}

	var ans [][]int
	for k := range set {
		ans = append(ans, []int{k[0], k[1], k[2], k[3]})
	}
	return ans
}
// Time Complexity: O(N3*log(M)), where N = size of the array, M = no. of elements in the set.
// Reason: Here, we are mainly using 3 nested loops, and inside the loops there are some operations on the set data structure which take log(M) time complexity.
// Space Complexity: O(2 * no. of the quadruplets)+O(N)
// Reason: we are using a set data structure and a list to store the quads. This results in the first term. And the second space is taken by the set data structure we are using to store the array elements. At most, the set can contain approximately all the array elements and so the space complexity is O(N).

func fourSumOptimal(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	var ans [][]int

	for i := 0; i < n; i++ {

		// Skip duplicates for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n; j++ {

			// Skip duplicates for j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k := j + 1
			l := n - 1

			for k < l {
				sum := int64(nums[i]) + int64(nums[j]) +
					int64(nums[k]) + int64(nums[l])

				if sum == int64(target) {
					ans = append(ans, []int{
						nums[i], nums[j], nums[k], nums[l],
					})

					k++
					l--

					// Skip duplicates
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}

				} else if sum < int64(target) {
					k++
				} else {
					l--
				}
			}
		}
	}

	return ans
}
// Time Complexity: O(N3), where N = size of the array.
// Reason: Each of the pointers i and j, is running for approximately N times. And both the pointers k and l combined can run for approximately N times including the operation of skipping duplicates. So the total time complexity will be O(N3). 
// Space Complexity: O(no. of quadruplets), This space is only used to store the answer. We are not using any extra space to solve this problem. So, from that perspective, space complexity can be written as O(1).

func main() {
	nums := []int{4, 3, 3, 4, 4, 2, 1, 2, 1, 1}
	target := 9

	ans := fourSumOptimal(nums, target)

	fmt.Println("Quadruplets:")
	for _, quad := range ans {
		fmt.Println(quad)
	}
}