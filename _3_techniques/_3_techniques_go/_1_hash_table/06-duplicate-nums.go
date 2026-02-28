// https://leetcode.com/problems/find-the-duplicate-number/

package main

import "fmt"

func findDuplicate(arr []int) int {
	n := len(arr)

	// create frequency slice of size n+1
	freq := make([]int, n+1)

	for i := 0; i < n; i++ {
		if freq[arr[i]] == 0 {
			freq[arr[i]]++
		} else {
			return arr[i]
		}
	}

	return 0
}

func main() {
	arr := []int{2, 1, 1}
	fmt.Println("The duplicate element is", findDuplicate(arr))
}

// Time: O(n)
// Space: O(n)