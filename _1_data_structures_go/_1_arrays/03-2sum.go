// https://leetcode.com/problems/two-sum/

package main

import (
	"fmt"
	"sort"
)

func twoSumBrute(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
// TC: O(n^2)
// SC: O(1)


type Pair struct {
	val int
	idx int
}

func twoSumTwoPointer(nums []int, target int) []int {
	n := len(nums)

	store := make([]Pair, n)

	for i := 0; i < n; i++ {
		store[i] = Pair{val: nums[i], idx: i}
	}

	// Sort by value
	sort.Slice(store, func(i, j int) bool {
		return store[i].val < store[j].val
	})

	left, right := 0, n-1

	for left < right {
		sum := store[left].val + store[right].val

		if sum == target {
			return []int{store[left].idx, store[right].idx}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
// TC: O(nlogn)
// SC: O(n)

// using map
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int) // value → index

	for i, num := range nums {

		if idx, found := mp[target-num]; found {
			return []int{i, idx}
		}

		mp[num] = i
	}

	return nil
}

// TC: O(n)
// SC: O(n)