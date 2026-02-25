// https://practice.geeksforgeeks.org/problems/inversion-of-array-1587115620/1

package main

import "fmt"

func numberOfInversionsBrute(a []int) int {
	n := len(a)
	cnt := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				cnt++
			}
		}
	}
	return cnt
}
// Time Complexity: O(N2), where N = size of the given array.
// Reason: We are using nested loops here and those two loops roughly run for N times.
// Space Complexity: O(1) as we are not using any extra space to solve this problem.


// using merge sort
func merge(arr []int, low, mid, high int) int {
	temp := []int{}
	left := low
	right := mid + 1
	cnt := 0

	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			cnt += (mid - left + 1) // Count inversions
			right++
		}
	}

	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	// Copy back
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}

	return cnt
}

func mergeSort(arr []int, low, high int) int {
	cnt := 0
	if low >= high {
		return cnt
	}

	mid := (low + high) / 2

	cnt += mergeSort(arr, low, mid)
	cnt += mergeSort(arr, mid+1, high)
	cnt += merge(arr, low, mid, high)

	return cnt
}

func numberOfInversionsOptimal(a []int) int {
	return mergeSort(a, 0, len(a)-1)
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	cnt := numberOfInversionsOptimal(a)

	fmt.Println("The number of inversions are:", cnt)
}