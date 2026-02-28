// https://leetcode.com/problems/merge-intervals/


package main

import (
	"fmt"
	"sort"
)


type Interval struct {
	Start int
	End   int
}

func merge(arr []Interval) []Interval {

	n := len(arr)

	// Sort intervals by Start
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Start < arr[j].Start
	})

	ans := []Interval{}

	for i := 0; i < n; i++ {

		start := arr[i].Start
		end := arr[i].End

		// If overlaps with last added interval → skip
		if len(ans) > 0 {
			if start <= ans[len(ans)-1].End {
				continue
			}
		}

		// Check all intervals to the right
		for j := i + 1; j < n; j++ {
			if arr[j].Start <= end {
				if arr[j].End > end {
					end = arr[j].End
				}
			}
		}

		ans = append(ans, Interval{Start: start, End: end})
	}

	return ans
}
// Time Complexity: O(NlogN)+O(N*N). O(NlogN) for sorting the array, and O(N*N) because we are checking to the right for each index which is a nested loop.
// Space Complexity: O(N), as we are using a separate data structure.


func mergeIntervals(intervals [][]int) [][]int {

	// Step 1: Sort intervals based on starting time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}

	for _, interval := range intervals {

		// If merged is empty OR no overlap
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			// Overlapping case → merge
			last := merged[len(merged)-1]
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		}
	}

	return merged
}
// Time Complexity: O(NlogN) + O(N). O(NlogN) for sorting and O(N) for traversing through the array.
// Space Complexity: O(N) to return the answer of the merged intervals.

func main() {

	arr := [][]int{
		{1, 3}, {2, 4}, {2, 6},
		{8, 9}, {8, 10}, {9, 11},
		{15, 18}, {16, 17},
	}

	ans := mergeIntervals(arr)

	fmt.Println("Merged Overlapping Intervals:")
	for _, it := range ans {
		fmt.Println(it[0], it[1])
	}
}