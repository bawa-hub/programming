// https://leetcode.com/problems/top-k-frequent-elements/

package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	freq int
	val  int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq // min heap by frequency
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	// Count frequency
	for _, num := range nums {
		freqMap[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	// Push into heap
	for val, freq := range freqMap {
		heap.Push(h, Pair{freq: freq, val: val})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var result []int

	// Extract elements
	for h.Len() > 0 {
		top := heap.Pop(h).(Pair)
		result = append(result, top.val)
	}

	return result
}

func main() {
	arr := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(arr, 2))
}