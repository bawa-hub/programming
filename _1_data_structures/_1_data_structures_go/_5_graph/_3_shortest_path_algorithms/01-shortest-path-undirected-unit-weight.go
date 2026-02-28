// https://practice.geeksforgeeks.org/problems/shortest-path-in-undirected-graph-having-unit-distance/1

package main

import (
	"fmt"
)

/************* SHORTEST PATH USING BFS *************/

func shortestPath(edges [][]int, N int, M int, src int) []int {

	// adjacency list
	adj := make([][]int, N)

	for _, e := range edges {
		u := e[0]
		v := e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// distance array
	const INF = int(1e9)
	dist := make([]int, N)

	for i := 0; i < N; i++ {
		dist[i] = INF
	}

	// BFS
	queue := []int{}
	dist[src] = 0
	queue = append(queue, src)

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		for _, next := range adj[node] {
			if dist[node]+1 < dist[next] {
				dist[next] = dist[node] + 1
				queue = append(queue, next)
			}
		}
	}

	// result
	ans := make([]int, N)
	for i := 0; i < N; i++ {
		if dist[i] == INF {
			ans[i] = -1
		} else {
			ans[i] = dist[i]
		}
	}

	return ans
}

/************* DRIVER *************/

func main() {

	N := 9
	M := 10
	edges := [][]int{
		{0, 1}, {0, 3}, {3, 4}, {4, 5},
		{5, 6}, {1, 2}, {2, 6},
		{6, 7}, {7, 8}, {6, 8},
	}

	ans := shortestPath(edges, N, M, 0)

	for _, v := range ans {
		fmt.Print(v, " ")
	}
}

// Time Complexity: O(M) { for creating the adjacency list from given list ‘edges’} + O(N + 2M) { for the BFS Algorithm} + O(N) { for adding the final values of the shortest path in the resultant array} ~ O(N+2M).
// Where N= number of vertices and M= number of edges.

// Space Complexity:  O( N) {for the stack storing the BFS} + O(N) {for the resultant array} + O(N) {for the dist array storing updated shortest paths} + O( N+2M) {for the adjacency list} ~ O(N+M) .
// Where N= number of vertices and M= number of edges.