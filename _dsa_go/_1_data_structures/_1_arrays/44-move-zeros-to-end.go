// https://leetcode.com/problems/move-zeroes/

package main

import "fmt"

func zerosToEndBrute(arr []int) {
	n := len(arr)

	temp := make([]int, n)
	k := 0

	// Copy non-zero elements
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			temp[k] = arr[i]
			k++
		}
	}

	// Remaining elements already 0 by default in Go
	fmt.Println(temp)
}
// Time complexity: o(n)
// Space complexity: o(n)

func zerosToEndOptimized(arr []int) {
	n := len(arr)

	// Find first zero
	k := 0
	for k < n {
		if arr[k] == 0 {
			break
		}
		k++
	}

	i := k
	j := k + 1

	for i < n && j < n {
		if arr[j] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		j++
	}

	fmt.Println(arr)
}
// Time complexity: o(n)
// Space complexity: o(1)

func main() {
	arr := []int{1, 2, 0, 1, 0, 4, 0}
	zerosToEndBrute(arr)
}