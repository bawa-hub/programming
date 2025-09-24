package main

import (
	"container/heap"
	"fmt"
)

// Define a pair (frequency, value)
type Pair struct {
	freq  int
	value int
}

// MaxHeap implementation
type MaxHeap []Pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq } // max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// TopKFrequent function
func TopKFrequent(nums []int, k int) []int {
	// Step 1: Count frequencies
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: Build max heap
	h := &MaxHeap{}
	heap.Init(h)
	for val, freq := range freqMap {
		heap.Push(h, Pair{freq: freq, value: val})
	}

	// Step 3: Extract top k elements
	res := []int{}
	for i := 0; i < k && h.Len() > 0; i++ {
		top := heap.Pop(h).(Pair)
		res = append(res, top.value)
	}

	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(TopKFrequent(nums, k)) // Example output: [1 2]
}


// using min heap

// // Pair stores frequency and value
// type Pair struct {
// 	freq  int
// 	value int
// }

// // MinHeap (compare by frequency)
// type MinHeap []Pair

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq } // min-heap
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x interface{}) {
// 	*h = append(*h, x.(Pair))
// }

// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

// func TopKFrequent(nums []int, k int) []int {
// 	// Step 1: Count frequencies
// 	freqMap := make(map[int]int)
// 	for _, num := range nums {
// 		freqMap[num]++
// 	}

// 	// Step 2: Min-heap of size k
// 	h := &MinHeap{}
// 	heap.Init(h)

// 	for val, freq := range freqMap {
// 		heap.Push(h, Pair{freq: freq, value: val})
// 		if h.Len() > k {
// 			heap.Pop(h) // remove smallest frequency
// 		}
// 	}

// 	// Step 3: Extract results
// 	res := make([]int, 0, k)
// 	for h.Len() > 0 {
// 		top := heap.Pop(h).(Pair)
// 		res = append(res, top.value)
// 	}

// 	return res
// }

// func main() {
// 	nums := []int{1, 1, 1, 2, 2, 3}
// 	k := 2
// 	fmt.Println(TopKFrequent(nums, k)) // Example output: [2 1] or [1 2]
// }


// bucket sort 
// func TopKFrequent(nums []int, k int) []int {
// 	// Step 1: Count frequencies
// 	freqMap := make(map[int]int)
// 	for _, num := range nums {
// 		freqMap[num]++
// 	}

// 	// Step 2: Create buckets
// 	buckets := make([][]int, len(nums)+1)
// 	for val, freq := range freqMap {
// 		buckets[freq] = append(buckets[freq], val)
// 	}

// 	// Step 3: Collect results from highest frequency
// 	res := []int{}
// 	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
// 		for _, val := range buckets[i] {
// 			res = append(res, val)
// 			if len(res) == k {
// 				break
// 			}
// 		}
// 	}

// 	return res
// }