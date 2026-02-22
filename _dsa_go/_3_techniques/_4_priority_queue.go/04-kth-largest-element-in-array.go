// https://leetcode.com/problems/kth-largest-element-in-an-array/
// https://www.geeksforgeeks.org/k-largestor-smallest-elements-in-an-array/
// https://takeuforward.org/data-structure/kth-largest-smallest-element-in-an-array/

package main

import (
	"container/heap"
	"fmt"
)

// ---------- MAX HEAP ----------
type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // MAX heap
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func kthLargest(arr []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, val := range arr {
		heap.Push(h, val)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return heap.Pop(h).(int)
}


// -------- OPTIMIZED FUNCTION --------
func kthSmallest(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)

		// keep heap size ≤ k
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// root is kth smallest
	return (*h)[0]
}
// (O(n log k))

// quick sort algo
func findKthLargest(arr []int, k int) int {
	left := 0
	right := len(arr) - 1

	for {
		idx := partitionLargest(arr, left, right)

		if idx == k-1 {
			return arr[idx]
		}

		if idx < k-1 {
			left = idx + 1
		} else {
			right = idx - 1
		}
	}
}

func partitionLargest(arr []int, left, right int) int {
	pivot := arr[left]
	l := left + 1
	r := right

	for l <= r {
		if arr[l] < pivot && arr[r] > pivot {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
		if l <= r && arr[l] >= pivot {
			l++
		}
		if l <= r && arr[r] <= pivot {
			r--
		}
	}

	arr[left], arr[r] = arr[r], arr[left]
	return r
}
// Time complexity: O(n) , where n = size of the array
// Space complexity: O(1) 

func main() {
	arr := []int{1, 2, 6, 4, 5, 3}

	fmt.Println("3rd Largest:", kthLargest(arr, 3))
}