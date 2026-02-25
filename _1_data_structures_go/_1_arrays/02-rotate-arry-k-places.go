// https://leetcode.com/problems/rotate-array/

package main

import "fmt"

func rotateToRight(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	temp := make([]int, k)

	// Store last k elements
	for i := n - k; i < n; i++ {
		temp[i-(n-k)] = arr[i]
	}

	// Shift remaining elements right
	for i := n - k - 1; i >= 0; i-- {
		arr[i+k] = arr[i]
	}

	// Copy temp back
	for i := 0; i < k; i++ {
		arr[i] = temp[i]
	}
}

func rotateToLeft(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	temp := make([]int, k)

	// Store first k elements
	for i := 0; i < k; i++ {
		temp[i] = arr[i]
	}

	// Shift elements left
	for i := 0; i < n-k; i++ {
		arr[i] = arr[i+k]
	}

	// Copy temp back
	for i := n - k; i < n; i++ {
		arr[i] = temp[i-(n-k)]
	}
}

// Time Complexity: O(n)
// Space Complexity: O(k) since k array element needs to be stored in temp array


// reversal algo
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func rotateRightOptimal(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}

	k = k % n

	// Reverse first n-k
	reverse(arr, 0, n-k-1)

	// Reverse last k
	reverse(arr, n-k, n-1)

	// Reverse whole array
	reverse(arr, 0, n-1)
}

func rotateLeftOptimal(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}

	k = k % n

	// Reverse first k
	reverse(arr, 0, k-1)

	// Reverse remaining
	reverse(arr, k, n-1)

	// Reverse whole array
	reverse(arr, 0, n-1)
}
// Time Complexity – O(N) where N is the number of elements in an array
// Space Complexity – O(1) since no extra space is required


func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 2

	rotateRightOptimal(arr, k)

	fmt.Println("After rotating right:", arr)
}