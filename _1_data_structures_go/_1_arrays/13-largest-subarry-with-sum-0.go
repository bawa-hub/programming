// https://practice.geeksforgeeks.org/problems/largest-subarray-with-0-sum/1

package main

import "fmt"

func solveBrute(a []int) int {
	n := len(a)
	maxLen := 0

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += a[j]

			if sum == 0 {
				length := j - i + 1
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}
	return maxLen
}
// Time Complexity: O(N^2) as we have two loops for traversal
// Space Complexity: O(1) as we aren’t using any extra space

func maxLen(a []int) int {
	mpp := make(map[int]int) // sum -> first index
	maxi := 0
	sum := 0

	for i := 0; i < len(a); i++ {
		sum += a[i]

		// If sum becomes 0 → subarray from 0 to i
		if sum == 0 {
			maxi = i + 1
		} else {
			if prevIndex, exists := mpp[sum]; exists {
				length := i - prevIndex
				if length > maxi {
					maxi = length
				}
			} else {
				// Store first occurrence only
				mpp[sum] = i
			}
		}
	}

	return maxi
}
// Time Complexity: O(N), as we are traversing the array only once
// Space Complexity: O(N), in the worst case we would insert all array elements prefix sum into our hashmap


func main() {
	a := []int{15, -2, 2, -8, 1, 7, 10, 23}

	fmt.Println("Brute:", solveBrute(a))
	fmt.Println("Optimal:", maxLen(a))
}
