// https://leetcode.com/problems/reverse-pairs/

package main

import "fmt"

func reversePairsBrute(arr []int) int {
	pairs := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > 2*arr[j] {
				pairs++
			}
		}
	}
	return pairs
}
//     Time Complexity: O (N^2) ( Nested Loops )
// Space Complexity:  O(1)

func merge(nums []int, low, mid, high int) int {
	total := 0
	j := mid + 1

	// Count reverse pairs
	for i := low; i <= mid; i++ {
		for j <= high && int64(nums[i]) > 2*int64(nums[j]) {
			j++
		}
		total += (j - (mid + 1))
	}

	// Standard merge process
	temp := []int{}
	left := low
	right := mid + 1

	for left <= mid && right <= high {
		if nums[left] <= nums[right] {
			temp = append(temp, nums[left])
			left++
		} else {
			temp = append(temp, nums[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, nums[left])
		left++
	}
	for right <= high {
		temp = append(temp, nums[right])
		right++
	}

	for i := low; i <= high; i++ {
		nums[i] = temp[i-low]
	}

	return total
}

func mergeSort(nums []int, low, high int) int {
	if low >= high {
		return 0
	}

	mid := (low + high) / 2
	count := 0

	count += mergeSort(nums, low, mid)
	count += mergeSort(nums, mid+1, high)
	count += merge(nums, low, mid, high)

	return count
}

func reversePairsOptimal(arr []int) int {
	return mergeSort(arr, 0, len(arr)-1)
}
// Time Complexity : O( N log N ) + O (N) + O (N)
// Reason: O(N) – Merge operation, O(N) – counting operation ( at each iteration of i, j doesn’t start from 0 . Both of them move linearly )
// Space Complexity : O(N)
// Reason : O(N) – Temporary vector


func main() {
	arr := []int{1, 3, 2, 3, 1}
	fmt.Println("The Total Reverse Pairs are", reversePairsOptimal(arr))
}