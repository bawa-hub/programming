// https://practice.geeksforgeeks.org/problems/minimum-cost-of-ropes-1587115620/1

package main

import (
	"container/heap"
	"fmt"
)

// -------- Min Heap --------

type MinHeap []int64

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j] // min heap
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// -------- Solution --------

func minCost(arr []int64) int64 {

	h := &MinHeap{}
	heap.Init(h)

	// push all ropes
	for _, v := range arr {
		heap.Push(h, v)
	}

	var cost int64 = 0

	for h.Len() >= 2 {
		first := heap.Pop(h).(int64)
		second := heap.Pop(h).(int64)

		sum := first + second
		cost += sum

		heap.Push(h, sum)
	}

	return cost
}

// -------- Example --------

func main() {
	arr := []int64{4, 3, 2, 6}

	fmt.Println(minCost(arr))
}