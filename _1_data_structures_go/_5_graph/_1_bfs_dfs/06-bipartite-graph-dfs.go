// https://leetcode.com/problems/is-graph-bipartite/

package main

import "fmt"

type Solution struct{}

// ---------- DFS Coloring ----------
func (s *Solution) dfs(node int, col int, color []int, adj [][]int) bool {

	color[node] = col

	for _, it := range adj[node] {

		// uncolored node
		if color[it] == -1 {
			if !s.dfs(it, 1-col, color, adj) {
				return false
			}
		} else if color[it] == col {
			// same color conflict
			return false
		}
	}

	return true
}

// ---------- Main Bipartite Function ----------
func (s *Solution) isBipartite(V int, adj [][]int) bool {

	color := make([]int, V)
	for i := 0; i < V; i++ {
		color[i] = -1
	}

	// handle disconnected components
	for i := 0; i < V; i++ {
		if color[i] == -1 {
			if !s.dfs(i, 0, color, adj) {
				return false
			}
		}
	}

	return true
}

// ---------- Add Edge ----------
func addEdge(adj [][]int, u, v int) {
	adj[u] = append(adj[u], v)
	adj[v] = append(adj[v], u)
}

func main() {

	V := 4
	adj := make([][]int, V)

	addEdge(adj, 0, 2)
	addEdge(adj, 0, 3)
	addEdge(adj, 2, 3)
	addEdge(adj, 3, 1)

	obj := Solution{}
	ans := obj.isBipartite(V, adj)

	if ans {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

// for leetcode
type Solution struct{}

func dfs(cur int, graph [][]int, color []int) bool {

	if color[cur] == -1 {
		color[cur] = 1
	}

	for _, nbr := range graph[cur] {

		if color[nbr] == -1 {
			color[nbr] = 1 - color[cur]

			if !dfs(nbr, graph, color) {
				return false
			}
		} else if color[nbr] == color[cur] {
			return false
		}
	}

	return true
}

func isBipartite(graph [][]int) bool {

	n := len(graph)
	color := make([]int, n)

	for i := 0; i < n; i++ {
		color[i] = -1
	}

	// handle disconnected graph
	for i := 0; i < n; i++ {
		if color[i] == -1 {
			if !dfs(i, graph, color) {
				return false
			}
		}
	}

	return true
}