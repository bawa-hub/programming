// https://leetcode.com/problems/path-with-minimum-effort/

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	effort int
	row    int
	col    int
}

// ---------- Min Heap ----------
type MinHeap []Node

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].effort < h[j].effort
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
func minimumEffortPath(heights [][]int) int {

	n := len(heights)
	m := len(heights[0])

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	pq := &MinHeap{}
	heap.Init(pq)

	dist[0][0] = 0
	heap.Push(pq, Node{0, 0, 0})

	for pq.Len() > 0 {

		cur := heap.Pop(pq).(Node)

		diff := cur.effort
		r := cur.row
		c := cur.col

		// reached destination
		if r == n-1 && c == m-1 {
			return diff
		}

		for i := 0; i < 4; i++ {

			nr := r + dr[i]
			nc := c + dc[i]

			if nr >= 0 && nr < n && nc >= 0 && nc < m {

				newEffort := max(
					diff,
					abs(heights[r][c]-heights[nr][nc]),
				)

				if newEffort < dist[nr][nc] {
					dist[nr][nc] = newEffort
					heap.Push(pq, Node{newEffort, nr, nc})
				}
			}
		}
	}

	return 0
}

// ---------- Helpers ----------
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ---------- Main ----------
func main() {

	heights := [][]int{
		{1, 2, 2},
		{3, 8, 2},
		{5, 3, 5},
	}

	ans := minimumEffortPath(heights)
	fmt.Println(ans)
}

// Time Complexity: O( 4*N*M * log( N*M) ) { N*M are the total cells, for each of which we also check 4 adjacent nodes for the minimum effort and additional log(N*M) for insertion-deletion operations in a priority queue }
// Where, N = No. of rows of the binary maze and M = No. of columns of the binary maze.
// Space Complexity: O( N*M ) { Distance matrix containing N*M cells + priority queue in the worst case containing all the nodes ( N*M) }.
// Where, N = No. of rows of the binary maze and M = No. of columns of the binary maze.