// https://practice.geeksforgeeks.org/problems/sum-of-elements-between-k1th-and-k2th-smallest-elements3133/1

package main

import (
	"container/heap"
	"fmt"
)

// ---------------- Max Heap ----------------

type MaxHeap []int64

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	// max heap
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// ---------------- kth Smallest ----------------

func kthSmallest(A []int64, k int64) int64 {
	h := &MaxHeap{}
	heap.Init(h)

	for _, val := range A {
		heap.Push(h, val)

		if int64(h.Len()) > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

// ---------------- Main Logic ----------------

func sumBetweenTwoKth(A []int64, K1, K2 int64) int64 {

	k1th := kthSmallest(A, K1)
	k2th := kthSmallest(A, K2)

	var sum int64 = 0

	for _, v := range A {
		if v > k1th && v < k2th {
			sum += v
		}
	}

	return sum
}

// ---------------- Example ----------------

func main() {
	A := []int64{20, 8, 22, 4, 12, 10, 14}
	K1 := int64(3)
	K2 := int64(6)

	ans := sumBetweenTwoKth(A, K1, K2)
	fmt.Println(ans)
}