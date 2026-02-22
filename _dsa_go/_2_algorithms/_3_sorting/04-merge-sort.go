// this is a typical example of divide and conqure algorithm

package main

import "fmt"

// Merge two sorted halves
func merge(arr []int, low, mid, high int) {
	temp := []int{}

	left := low
	right := mid + 1

	// Merge in sorted order
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	// Remaining left side
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	// Remaining right side
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	// Copy back to original array
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

func mergeSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2

	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)

	merge(arr, low, mid, high)
}

// 📊 Time Complexity
// Divide:
// log n levels
// Merge at each level:
// O(n)
// Total:
// O(n log n)
// Space Complexity - O(n)

func main() {
	arr := []int{9, 4, 7, 6, 3, 1, 5}

	fmt.Println("Before Sorting:", arr)

	mergeSort(arr, 0, len(arr)-1)

	fmt.Println("After Sorting:", arr)
}