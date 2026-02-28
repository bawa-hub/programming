// https://leetcode.com/problems/network-delay-time/
// https://practice.geeksforgeeks.org/problems/alex-travelling/1

// https://www.youtube.com/watch?v=F3PNsWE6_hM&list=PLauivoElc3ghxyYSr_sVnDUc_ynPk6iXE&index=16


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

// ---------- Dijkstra ----------
func dijkstra(source int, n int, graph [][]Edge) int {

	dist := make([]int, n+1)
	visited := make([]bool, n+1)

	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[source] = 0

	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, Node{0, source})

	for pq.Len() > 0 {

		cur := heap.Pop(pq).(Node)
		v := cur.node

		if visited[v] {
			continue
		}
		visited[v] = true

		for _, child := range graph[v] {

			if dist[v]+child.wt < dist[child.to] {
				dist[child.to] = dist[v] + child.wt
				heap.Push(pq, Node{dist[child.to], child.to})
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	return ans
}

// ---------- Network Delay ----------
func networkDelayTime(times [][]int, n int, k int) int {

	graph := make([][]Edge, n+1)

	for _, t := range times {
		u := t[0]
		v := t[1]
		w := t[2]
		graph[u] = append(graph[u], Edge{v, w})
	}

	return dijkstra(k, n, graph)
}

// ---------- Main ----------
func main() {

	times := [][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1},
	}

	n := 4
	k := 2

	fmt.Println(networkDelayTime(times, n, k))
}