
// A graph that can be colored with two colors such that no two adjacent nodes has the same color
// if a graph has odd length cycle then it is not bipartite else bipartite

// https://leetcode.com/problems/is-graph-bipartite/
// https://takeuforward.org/graph/bipartite-graph-bfs-implementation/

package main

import "fmt"

type Solution struct{}

// ---------- BFS Coloring ----------
func (s *Solution) check(start int, adj [][]int, color []int) bool {

	queue := []int{start}
	color[start] = 0

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		for _, it := range adj[node] {

			// not colored yet
			if color[it] == -1 {
				color[it] = 1 - color[node]
				queue = append(queue, it)

			} else if color[it] == color[node] {
				// same color → not bipartite
				return false
			}
		}
	}

	return true
}

// ---------- Main Function ----------
func (s *Solution) isBipartite(V int, adj [][]int) bool {

	color := make([]int, V)

	for i := 0; i < V; i++ {
		color[i] = -1
	}

	for i := 0; i < V; i++ {
		if color[i] == -1 {
			if !s.check(i, adj, color) {
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