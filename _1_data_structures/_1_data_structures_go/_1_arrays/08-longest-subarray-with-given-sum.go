// https://www.codingninjas.com/codestudio/problems/longest-subarray-with-sum-k_6682399
// https://www.geeksforgeeks.org/problems/longest-sub-array-with-sum-k0809

package main

import "fmt"

func getLongestSubarrayBrute(a []int, k int64) int {
	n := len(a)
	maxLen := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {

			var sum int64 = 0
			for x := i; x <= j; x++ {
				sum += int64(a[x])
			}

			if sum == k && j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}
// Time Complexity: O(n^3)
// Space Complexity: O(1)

// optimized brute force
func getLongestSubarrayBetter(a []int, k int64) int {
	n := len(a)
	maxLen := 0

	for i := 0; i < n; i++ {
		var sum int64 = 0

		for j := i; j < n; j++ {
			sum += int64(a[j])

			if sum == k && j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}
// Time Complexity: O(n^2) time to generate all possible subarrays.
// Space Complexity: O(1), we are not using any extra space.

// prefix sum + hashing (works for negatives)
func getLongestSubarrayOptimal(a []int, k int64) int {
	preSumMap := make(map[int64]int)
	var sum int64 = 0
	maxLen := 0

	for i := 0; i < len(a); i++ {

		sum += int64(a[i])

		// If prefix sum itself equals k
		if sum == k {
			maxLen = i + 1
		}

		rem := sum - k

		if prevIndex, exists := preSumMap[rem]; exists {
			length := i - prevIndex
			if length > maxLen {
				maxLen = length
			}
		}

		// Store first occurrence only
		if _, exists := preSumMap[sum]; !exists {
			preSumMap[sum] = i
		}
	}

	return maxLen
}
// Time Complexity: O(N) or O(N*logN) depending on which map data structure we are using, where N = size of the array.
// Space Complexity: O(N) as we are using a map data structure.

// sliding window (only for positives)
func longestSubarrayWithSumK(a []int, k int64) int {
	i, j := 0, 0
	var sum int64 = 0
	maxi := 0

	for j < len(a) {

		sum += int64(a[j])

		for sum > k && i <= j {
			sum -= int64(a[i])
			i++
		}

		if sum == k {
			if j-i+1 > maxi {
				maxi = j - i + 1
			}
		}

		j++
	}

	return maxi
}
// Time Complexity: O(2*N)
// Space Complexity: O(1)

func main() {
	a := []int{10, 5, 2, 7, 1, 9}
	k := int64(15)

	fmt.Println("Longest subarray length:", getLongestSubarrayOptimal(a, k))
}