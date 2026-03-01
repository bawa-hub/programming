// https://practice.geeksforgeeks.org/problems/implementing-dijkstra-set-1-adjacency-matrix/1

package main

import (
	"container/heap"
	"fmt"
	"math"
)

/************* PRIORITY QUEUE *************/

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

/************* DIJKSTRA *************/

func dijkstra(V int, adj [][][]int, S int) []int {

	dist := make([]int, V)
	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt32
	}

	dist[S] = 0

	pq := &MinHeap{}
	heap.Init(pq)

	heap.Push(pq, Pair{0, S})

	for pq.Len() > 0 {

		top := heap.Pop(pq).(Pair)
		dis := top.dist
		node := top.node

		for _, it := range adj[node] {
			v := it[0]
			w := it[1]

			if dis+w < dist[v] {
				dist[v] = dis + w
				heap.Push(pq, Pair{dist[v], v})
			}
		}
	}

	return dist
}

/************* DRIVER *************/

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

// Time Complexity: O( E log(V) ), Where E = Number of edges and V = Number of Nodes.
// Space Complexity: O( |E| + |V| ), Where E = Number of edges and V = Number of Nodes.