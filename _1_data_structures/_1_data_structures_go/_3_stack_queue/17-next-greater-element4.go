// https://leetcode.com/problems/next-greater-element-iv/
// https://leetcode.com/problems/next-greater-element-iv/solutions/2799346/monotonic-stack-min-heap/

package main

import (
	"container/heap"
	"fmt"
)

/**************** MIN HEAP ****************/

type Pair struct {
	val int
	idx int
}

type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

/**************** SOLUTION ****************/

func secondGreaterElement(nums []int) []int {

	n := len(nums)
	gE := make([]int, n)

	for i := range gE {
		gE[i] = -1
	}

	stack := []int{}
	pq := &MinHeap{}
	heap.Init(pq)

	for i := 0; i < n; i++ {

		// resolve second greater
		for pq.Len() > 0 && (*pq)[0].val < nums[i] {
			top := heap.Pop(pq).(Pair)
			gE[top.idx] = nums[i]
		}

		// move first greater candidates
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			heap.Push(pq, Pair{nums[idx], idx})
		}

		stack = append(stack, i)
	}

	return gE
}

/**************** TEST ****************/

func main() {

	nums := []int{2, 4, 0, 9, 6}

	res := secondGreaterElement(nums)

	fmt.Println(res)
}