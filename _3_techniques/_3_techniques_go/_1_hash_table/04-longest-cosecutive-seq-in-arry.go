// https://leetcode.com/problems/longest-consecutive-sequence/description/

package main

import (
	"fmt"
	"sort"
)

// brute force
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Sort the slice
	sort.Ints(nums)

	ans := 1
	prev := nums[0]
	cur := 1

	for i := 1; i < len(nums); i++ {

		if nums[i] == prev+1 {
			cur++
		} else if nums[i] != prev {
			cur = 1
		}

		prev = nums[i]

		if cur > ans {
			ans = cur
		}
	}

	return ans
}
//     Time Complexity: We are first sorting the array which will take O(N * log(N)) time and then we are running a for loop which will take O(N) time. Hence, the overall time complexity will be O(N * log(N)).
// Space Complexity: The space complexity for the above approach is O(1) because we are not using any auxiliary space

// optimized

func longestConsecutiveOptimized(nums []int) int {
	set := make(map[int]bool)

	// Insert into set
	for _, num := range nums {
		set[num] = true
	}

	longestStreak := 0

	for _, num := range nums {

		// Check if it's start of sequence
		if !set[num-1] {

			currentNum := num
			currentStreak := 1

			for set[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
// Time Complexity: The time complexity of the above approach is O(N) because we traverse each consecutive subsequence only once. (assuming HashSet takes O(1) to search)
// Space Complexity: The space complexity of the above approach is O(N) because we are maintaining a HashSet.


func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // Output: 4
}
