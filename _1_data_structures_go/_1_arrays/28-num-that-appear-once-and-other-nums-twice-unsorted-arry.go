// https://leetcode.com/problems/single-number/
// https://takeuforward.org/arrays/find-the-number-that-appears-once-and-the-other-numbers-twice/
// https://practice.geeksforgeeks.org/problems/element-appearing-once2552/0


package main

import "fmt"

func getSingleElementBrute(arr []int) int {
	n := len(arr)

	for i := 0; i < n; i++ {
		num := arr[i]
		count := 0

		for j := 0; j < n; j++ {
			if arr[j] == num {
				count++
			}
		}

		if count == 1 {
			return num
		}
	}

	return -1
}
// Time Complexity: O(N2), where N = size of the given array.
// Reason: For every element, we are performing a linear search to count its occurrence. The linear search takes O(N) time complexity. And there are N elements in the array. So, the total time complexity will be O(N2).
// Space Complexity: O(1) as we are not using any extra space.

// using hashing
func getSingleElementHash(arr []int) int {
	freq := make(map[int]int)

	for _, v := range arr {
		freq[v]++
	}

	for key, value := range freq {
		if value == 1 {
			return key
		}
	}

	return -1
}
// Time Complexity: O(N)+O(N)+O(N), where N = size of the array
// Reason: One O(N) is for finding the maximum, the second one is to hash the elements and the third one is to search the single element in the array.
// Space Complexity: O(maxElement+1) where maxElement = the maximum element of the array.

// using xor
func getSingleElementXOR(arr []int) int {
	xorr := 0

	for _, v := range arr {
		xorr ^= v
	}

	return xorr
}


func main() {
	arr := []int{4, 1, 2, 1, 2}
	ans := getSingleElementXOR(arr)
	fmt.Println("The single element is:", ans)
}