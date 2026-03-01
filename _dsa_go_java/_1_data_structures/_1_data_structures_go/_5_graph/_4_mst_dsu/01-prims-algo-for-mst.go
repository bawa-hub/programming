// https://practice.geeksforgeeks.org/problems/minimum-spanning-tree/1

package main

import (
	"container/heap"
	"fmt"
)

// ---------- Min Heap ----------

type Pair struct {
	weight int
	node   int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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

// ---------- Prim's Algorithm ----------

func spanningTree(V int, adj [][][2]int) int {

	visited := make([]bool, V)

	pq := &MinHeap{}
	heap.Init(pq)

	// {weight, node}
	heap.Push(pq, Pair{0, 0})

	sum := 0

	for pq.Len() > 0 {

		top := heap.Pop(pq).(Pair)
		node := top.node
		wt := top.weight

		if visited[node] {
			continue
		}

		visited[node] = true
		sum += wt

		for _, edge := range adj[node] {
			adjNode := edge[0]
			edgeWeight := edge[1]

			if !visited[adjNode] {
				heap.Push(pq, Pair{edgeWeight, adjNode})
			}
		}
	}

	return sum
}

// ---------- Driver ----------

func main() {

	V := 5

	edges := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 2, 1},
		{2, 3, 2},
		{3, 4, 1},
		{4, 2, 2},
	}

	adj := make([][][2]int, V)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]

		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	sum := spanningTree(V, adj)

	fmt.Println("Sum of MST weights:", sum)
}

// Time Complexity: O(E*logE) + O(E*logE)~ O(E*logE), where E = no. of given edges.
// The maximum size of the priority queue can be E so after at most E iterations the priority queue will be empty and the loop will end. Inside the loop, there is a pop operation that will take logE time. This will result in the first O(E*logE) time complexity. Now, inside that loop, for every node, we need to traverse all its adjacent nodes where the number of nodes can be at most E. If we find any node unvisited, we will perform a push operation and for that, we need a logE time complexity. So this will result in the second O(E*logE). 
// Space Complexity: O(E) + O(V), where E = no. of edges and V = no. of vertices. O(E) occurs due to the size of the priority queue and O(V) due to the visited array. If we wish to get the mst, we need an extra O(V-1) space to store the edges of the most.
