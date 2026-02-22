package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int { 
    return len(h) 
}

func (h MaxHeap) Less(i, j int) bool { 
    return h[i] > h[j]   // MAX heap
}

func (h MaxHeap) Swap(i, j int) { 
    h[i], h[j] = h[j], h[i] 
}

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func main() {
	h := &MaxHeap{}
	heap.Init(h)

	heap.Push(h, 10)
	heap.Push(h, 5)
	heap.Push(h, 20)

	top := heap.Pop(h).(int)
	fmt.Println("top: ", top)
}