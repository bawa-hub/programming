package main

import (
	"container/heap"
	"fmt"
	"math"
)

/************ MIN HEAP ************/

type Pair struct {
	dist int
	node int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
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

/************ DIJKSTRA ************/

func dijkstra(V int, adj [][][]int, S int) []int {

	dist := make([]int, V)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[S] = 0

	pq := &MinHeap{}
	heap.Init(pq)

	heap.Push(pq, Pair{0, S})

	for pq.Len() > 0 {

		top := heap.Pop(pq).(Pair)
		node := top.node
		dis := top.dist

		// Equivalent to set erase outdated entry
		if dis > dist[node] {
			continue
		}

		for _, it := range adj[node] {
			adjNode := it[0]
			weight := it[1]

			if dis+weight < dist[adjNode] {
				dist[adjNode] = dis + weight
				heap.Push(pq, Pair{dist[adjNode], adjNode})
			}
		}
	}

	return dist
}

/************ DRIVER ************/

func main() {

	V := 3
	S := 2

	adj := make([][][]int, V)

	adj[0] = append(adj[0], []int{1, 1})
	adj[0] = append(adj[0], []int{2, 6})

	adj[1] = append(adj[1], []int{2, 3})
	adj[1] = append(adj[1], []int{0, 1})

	adj[2] = append(adj[2], []int{1, 3})
	adj[2] = append(adj[2], []int{0, 6})

	res := dijkstra(V, adj, S)

	for _, v := range res {
		fmt.Print(v, " ")
	}
}