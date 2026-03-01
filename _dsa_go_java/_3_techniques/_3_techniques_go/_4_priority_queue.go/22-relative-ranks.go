// https://leetcode.com/problems/relative-ranks/description/

package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

// -------- Max Heap --------

type Pair struct {
	score int
	index int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }

// max heap
func (h MaxHeap) Less(i, j int) bool {
	return h[i].score > h[j].score
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
	val := old[n-1]
	*h = old[:n-1]
	return val
}

// -------- Solution --------

func findRelativeRanks(score []int) []string {

	h := &MaxHeap{}
	heap.Init(h)

	for i := 0; i < len(score); i++ {
		heap.Push(h, Pair{score[i], i})
	}

	res := make([]string, len(score))
	rank := 1

	for h.Len() > 0 {
		p := heap.Pop(h).(Pair)

		switch rank {
		case 1:
			res[p.index] = "Gold Medal"
		case 2:
			res[p.index] = "Silver Medal"
		case 3:
			res[p.index] = "Bronze Medal"
		default:
			res[p.index] = strconv.Itoa(rank)
		}

		rank++
	}

	return res
}

// -------- Example --------

func main() {
	score := []int{5, 4, 3, 2, 1}
	fmt.Println(findRelativeRanks(score))
}