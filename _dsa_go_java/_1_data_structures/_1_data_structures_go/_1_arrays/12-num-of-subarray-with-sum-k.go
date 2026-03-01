

package main

import "fmt"

func findAllSubarraysBrute(arr []int, k int) int {
	n := len(arr)
	cnt := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {

			sum := 0
			for x := i; x <= j; x++ {
				sum += arr[x]
			}

			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}
// Time Complexity: O(N3), where N = size of the array.
// Reason: We are using three nested loops here. Though all are not running for exactly N times, the time complexity will be approximately O(N3).
// Space Complexity: O(1) as we are not using any extra space.

func findAllSubarraysBetter(arr []int, k int) int {
	n := len(arr)
	cnt := 0

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}
// Time Complexity: O(N2), where N = size of the array.
// Reason: We are using two nested loops here. As each of them is running for exactly N times, the time complexity will be approximately O(N2).

// Space Complexity: O(1) as we are not using any extra space.

// prefix sum + hashing
func findAllSubarraysOptimal(arr []int, k int) int {
	mpp := make(map[int]int) // prefixSum -> count
	preSum := 0
	cnt := 0

	mpp[0] = 1 // Important initialization

	for _, val := range arr {
		preSum += val

		remove := preSum - k

		if freq, exists := mpp[remove]; exists {
			cnt += freq
		}

		mpp[preSum]++
	}

	return cnt
}
// Time Complexity: O(N) or O(N*logN) depending on which map data structure we are using, where N = size of the array.
// Reason: For example, if we are using an unordered_map data structure in C++ the time complexity will be O(N) but if we are using a map data structure, the time complexity will be O(N*logN). The least complexity will be O(N) as we are using a loop to traverse the array.

// Space Complexity: O(N) as we are using a map data structure.

func main() {
	arr := []int{3, 1, 2, 4}
	k := 6

	cnt := findAllSubarraysOptimal(arr, k)
	fmt.Println("The number of subarrays is:", cnt)
}