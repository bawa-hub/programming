// https://practice.geeksforgeeks.org/problems/implementing-dijkstra-set-1-adjacency-matrix/1

package main

import (
	"container/heap"
	"fmt"
	"math"
)

/************* MIN HEAP *************/

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

/************* DIJKSTRA PATH *************/

func shortestPath(n int, edges [][]int) []int {

	// adjacency list
	adj := make([][][2]int, n+1)

	for _, e := range edges {
		u := e[0]
		v := e[1]
		w := e[2]

		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	dist := make([]int, n+1)
	parent := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
		parent[i] = i
	}

	dist[1] = 0

	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, Pair{0, 1})

	for pq.Len() > 0 {

		top := heap.Pop(pq).(Pair)
		node := top.node
		dis := top.dist

		if dis > dist[node] {
			continue
		}

		for _, it := range adj[node] {
			adjNode := it[0]
			weight := it[1]

			if dis+weight < dist[adjNode] {
				dist[adjNode] = dis + weight
				parent[adjNode] = node
				heap.Push(pq, Pair{dist[adjNode], adjNode})
			}
		}
	}

	// unreachable
	if dist[n] == math.MaxInt32 {
		return []int{-1}
	}

	/******** PATH RECONSTRUCTION ********/

	path := []int{}
	node := n

	for parent[node] != node {
		path = append(path, node)
		node = parent[node]
	}
	path = append(path, 1)

	// reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

/************* DRIVER *************/

func main() {

	n := 5
	edges := [][]int{
		{1, 2, 2},
		{2, 5, 5},
		{2, 3, 4},
		{1, 4, 1},
		{4, 3, 3},
		{3, 5, 1},
	}

	path := shortestPath(n, edges)

	for _, v := range path {
		fmt.Print(v, " ")
	}
}