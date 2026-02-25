// https://leetcode.com/problems/majority-element-ii/

package main

import "fmt"

func majorityElementBrute(arr []int) []int {
	n := len(arr)
	var ans []int

	for i := 0; i < n; i++ {
		count := 1
		for j := i + 1; j < n; j++ {
			if arr[j] == arr[i] {
				count++
			}
		}

		if count > n/3 {
			// Avoid duplicates in answer
			exists := false
			for _, v := range ans {
				if v == arr[i] {
					exists = true
					break
				}
			}
			if !exists {
				ans = append(ans, arr[i])
			}
		}
	}
	return ans
}
// Time Complexity: O(n^2)
// Space Complexity : O(1)

func majorityElementBetter(arr []int) []int {
	n := len(arr)
	mp := make(map[int]int)
	var ans []int

	for _, v := range arr {
		mp[v]++
	}

	for key, count := range mp {
		if count > n/3 {
			ans = append(ans, key)
		}
	}

	return ans
}
// Time Complexity: O(n)
// Space Complexity : O(n)

// Optimal Solution (Extended Boyer Moore’s Voting Algorithm)
func majorityElementOptimal(nums []int) []int {
	n := len(nums)

	var num1, num2 int
	count1, count2 := 0, 0

	// Phase 1: Find potential candidates
	for _, v := range nums {

		if count1 > 0 && v == num1 {
			count1++
		} else if count2 > 0 && v == num2 {
			count2++
		} else if count1 == 0 {
			num1 = v
			count1 = 1
		} else if count2 == 0 {
			num2 = v
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	// Phase 2: Verify candidates
	count1, count2 = 0, 0
	for _, v := range nums {
		if v == num1 {
			count1++
		} else if v == num2 {
			count2++
		}
	}

	var ans []int
	if count1 > n/3 {
		ans = append(ans, num1)
	}
	if count2 > n/3 {
		ans = append(ans, num2)
	}

	return ans
}
// Time Complexity: O(n)
// Space Complexity : O(1)


func main() {
	arr := []int{1, 2, 2, 3, 2}
	majority := majorityElementOptimal(arr)

	fmt.Println("Majority elements (> n/3):", majority)
}