// https://www.geeksforgeeks.org/problems/detect-cycle-in-an-undirected-graph/1

package main

import "fmt"

type Solution struct{}

// ---------- DFS Cycle Check ----------
func (s *Solution) dfs(node int, parent int, vis []bool, adj [][]int) bool {
	vis[node] = true

	for _, adjacentNode := range adj[node] {

		// unvisited neighbour
		if !vis[adjacentNode] {
			if s.dfs(adjacentNode, node, vis, adj) {
				return true
			}
		} else if adjacentNode != parent {
			// visited but not parent → cycle
			return true
		}
	}

	return false
}

// ---------- Detect Cycle ----------
func (s *Solution) isCycle(V int, adj [][]int) bool {

	vis := make([]bool, V)

	// handle disconnected components
	for i := 0; i < V; i++ {
		if !vis[i] {
			if s.dfs(i, -1, vis, adj) {
				return true
			}
		}
	}
	return false
}

func main() {

	adj := [][]int{
		{},
		{2},
		{1, 3},
		{2},
	}

	obj := Solution{}
	ans := obj.isCycle(4, adj)

	if ans {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

// Time Complexity: O(N + 2E) + O(N), Where N = Nodes, 2E is for total degrees as we traverse all adjacent nodes. In the case of connected components of a graph, it will take another O(N) time.
// Space Complexity: O(N) + O(N) ~ O(N), Space for recursive stack space and visited array