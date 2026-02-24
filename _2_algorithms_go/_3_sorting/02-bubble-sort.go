// Repeatedly swap 2 adjacent elements if arr[j] > arr[j+1] .
// Here, the maximum element of the unsorted array reaches the end of the unsorted array after each iteration

package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
// Time complexity: O(N2), (where N = size of the array), for the worst, and average cases.
// Reason: If we carefully observe, we can notice that the outer loop, say i, is running from n-1 to 0 i.e. n times, and for each i, the inner loop j runs from 0 to i-1. For, i = n-1, the inner loop runs n-1 times, for i = n-2, the inner loop runs n-2 times, and so on. So, the total steps will be approximately the following: (n-1) + (n-2) + (n-3) + ……..+ 3 + 2 + 1. The summation is approximately the sum of the first n natural numbers i.e. (n*(n+1))/2. The precise time complexity will be O(n2/2 + n/2). Previously, we have learned that we can ignore the lower values as well as the constant coefficients. So, the time complexity is O(n2). Here the value of n is N i.e. the size of the array.
// Space Complexity: O(1)

// optimized
func bubbleSortOptimized(arr []int) {
	n := len(arr)

	for i := n - 1; i >= 0; i-- {
		didSwap := false

		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
				didSwap = true
			}
		}

		// If no swaps happened, array is already sorted
		if !didSwap {
			break
		}
	}
}
// Time Complexity: O(N2) for the worst and average cases and O(N) for the best case. Here, N = size of the array.
// Space Complexity: O(1)

func main() {
	arr := []int{13, 46, 24, 52, 20, 9}

	fmt.Println("Before Bubble Sort:")
	fmt.Println(arr)

	bubbleSort(arr)

	fmt.Println("After Bubble Sort:")
	fmt.Println(arr)
}