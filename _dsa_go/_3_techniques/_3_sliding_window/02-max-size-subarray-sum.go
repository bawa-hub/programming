// https://www.codingninjas.com/studio/problems/longest-subarray-with-sum-k_6682399

package main

import "fmt"

func longestSubarrayWithSumK(a []int, k int64) int {
	i, j := 0, 0
	var sum int64 = 0
	maxi := 0

	n := len(a)

	for j < n {
		sum += int64(a[j])

		// Shrink window if sum exceeds k
		for sum > k && i <= j {
			sum -= int64(a[i])
			i++
		}

		// Check if equal
		if sum == k {
			length := j - i + 1
			if length > maxi {
				maxi = length
			}
		}

		j++
	}

	return maxi
}

func main() {
	arr := []int{1, 2, 3, 1, 1, 1, 1}
	k := int64(3)
	fmt.Println(longestSubarrayWithSumK(arr, k)) // Output: 3
}