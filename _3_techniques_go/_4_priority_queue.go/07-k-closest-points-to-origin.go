// https://leetcode.com/problems/k-closest-points-to-origin/

package main

import (
	"container/heap"
	"fmt"
)

// ---------------- Heap Structure ----------------

type Point struct {
	dist int
	x    int
	y    int
}

type MaxHeap []Point

func (h MaxHeap) Len() int { return len(h) }

// Max Heap (reverse comparison)
func (h MaxHeap) Less(i, j int) bool {
	return h[i].dist > h[j].dist
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// ---------------- Solution ----------------

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, p := range points {

		// squared distance (no sqrt needed)
		dist := p[0]*p[0] + p[1]*p[1]

		heap.Push(h, Point{dist, p[0], p[1]})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := [][]int{}
	for h.Len() > 0 {
		p := heap.Pop(h).(Point)
		res = append(res, []int{p.x, p.y})
	}

	return res
}

// ---------------- Main ----------------

func main() {
	points := [][]int{{1, 3}, {-2, 2}, {5, 8}}
	k := 2

	ans := kClosest(points, k)
	fmt.Println(ans)
}