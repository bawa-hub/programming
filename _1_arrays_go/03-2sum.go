// https://leetcode.com/problems/two-sum/

package main

import (
	"fmt"
	"sort"
)

// Brute force
// TC: O(n^2), SC: O(1)
func twoSumBruteForce(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Two-pointer approach
// TC: O(n log n), SC: O(n)
func twoSumTwoPointer(nums []int, target int) []int {
	n := len(nums)
	store := append([]int{}, nums...) // make a copy
	sort.Ints(store)

	left, right := 0, n-1
	n1, n2 := 0, 0

	for left < right {
		sum := store[left] + store[right]
		if sum == target {
			n1, n2 = store[left], store[right]
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	res := []int{}
	for i := 0; i < n; i++ {
		if nums[i] == n1 || nums[i] == n2 {
			res = append(res, i)
		}
	}
	return res
}

// Hashing approach (most efficient)
// TC: O(n), SC: O(n)
func twoSumHashing(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, num := range nums {
		if idx, found := mp[target-num]; found {
			return []int{i, idx}
		}
		mp[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println("Brute Force:", twoSumBruteForce(nums, target))
	fmt.Println("Two Pointer:", twoSumTwoPointer(nums, target))
	fmt.Println("Hashing:", twoSumHashing(nums, target))
}
