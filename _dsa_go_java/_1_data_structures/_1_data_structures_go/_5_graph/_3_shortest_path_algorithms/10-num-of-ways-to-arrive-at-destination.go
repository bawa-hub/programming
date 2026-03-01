// https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to int
	wt int
}

type Node struct {
	dist int
	node int
}

// ---------- Min Heap ----------
type MinHeap []Node

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// ---------- Solution ----------
func countPaths(n int, roads [][]int) int {

	graph := make([][]Edge, n)
	for _, r := range roads {
		u, v, w := r[0], r[1], r[2]
		graph[u] = append(graph[u], Edge{v, w})
		graph[v] = append(graph[v], Edge{u, w})
	}

	const mod = 1_000_000_007

	dist := make([]int, n)
	ways := make([]int, n)

	for i := range dist {
		dist[i] = math.MaxInt64
	}

	dist[0] = 0
	ways[0] = 1

	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, Node{0, 0})

	for pq.Len() > 0 {

		cur := heap.Pop(pq).(Node)
		dis := cur.dist
		node := cur.node

		for _, edge := range graph[node] {

			newDist := dis + edge.wt

			// better path found
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				ways[edge.to] = ways[node]
				heap.Push(pq, Node{newDist, edge.to})
			} else if newDist == dist[edge.to] {
				// another shortest path found
				ways[edge.to] =
					(ways[edge.to] + ways[node]) % mod
			}
		}
	}

	return ways[n-1] % mod
}

// ---------- Main ----------
func main() {

	n := 7

	roads := [][]int{
		{0, 6, 7},
		{0, 1, 2},
		{1, 2, 3},
		{1, 3, 3},
		{6, 3, 3},
		{3, 5, 1},
		{6, 5, 1},
		{2, 5, 1},
		{0, 4, 5},
		{4, 6, 2},
	}

	fmt.Println(countPaths(n, roads))
}

// Time Complexity: O( E* log(V)) { As we are using simple Dijkstra’s algorithm here, the time complexity will be or the order E*log(V)}
// Where E = Number of edges and V = No. of vertices.
// Space Complexity :  O(N) { for dist array + ways array + approximate complexity for priority queue }
// Where, N = Number of nodes.