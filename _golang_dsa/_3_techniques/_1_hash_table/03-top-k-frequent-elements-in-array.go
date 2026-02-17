// https://practice.geeksforgeeks.org/problems/top-k-frequent-elements-in-array/1
// https://leetcode.com/problems/top-k-frequent-elements/

import "container/heap"

type Pair struct {
	freq int
	val  int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }

// Max heap → larger freq has higher priority
func (h MaxHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}


// using max heap
func topKFrequent(nums []int, k int) []int {
	// Frequency map
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	// Create heap
	h := &MaxHeap{}
	heap.Init(h)

	// Push all elements into heap
	for val, freq := range mp {
		heap.Push(h, Pair{freq: freq, val: val})
	}

	// Extract top k
	var res []int
	for k > 0 {
		top := heap.Pop(h).(Pair)
		res = append(res, top.val)
		k--
	}

	return res
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }

// Min heap → smaller freq has higher priority
func (h MinHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
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
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// min heap
func topKFrequent(nums []int, k int) []int {
	// Frequency map
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for val, freq := range mp {
		heap.Push(h, Pair{freq: freq, val: val})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res []int

	for h.Len() > 0 {
		top := heap.Pop(h).(Pair)
		res = append(res, top.val)
	}

	return res
}

// optimal O(n) bucket sort
func topKFrequent(nums []int, k int) []int {
	// Step 1: Build frequency map
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Step 2: Create buckets
	bucket := make([][]int, len(nums)+1)

	for num, count := range freq {
		bucket[count] = append(bucket[count], num)
	}

	// Step 3: Traverse from highest frequency
	var res []int

	for i := len(nums); i >= 0 && len(res) < k; i-- {
		for _, num := range bucket[i] {
			res = append(res, num)
			if len(res) == k {
				break
			}
		}
	}

	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}