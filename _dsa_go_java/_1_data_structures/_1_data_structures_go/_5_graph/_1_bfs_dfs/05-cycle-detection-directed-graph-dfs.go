// https://takeuforward.org/data-structure/detect-cycle-in-a-directed-graph-using-dfs-g-19/
// https://leetcode.com/problems/course-schedule-ii/solutions/293048/detecting-cycle-in-directed-graph-problem/

package main

import "fmt"

type Solution struct{}

// ---------- DFS Cycle Check ----------
func (s *Solution) dfsCheck(
	node int,
	adj [][]int,
	vis []bool,
	pathVis []bool,
) bool {

	vis[node] = true
	pathVis[node] = true

	for _, it := range adj[node] {

		// not visited
		if !vis[it] {
			if s.dfsCheck(it, adj, vis, pathVis) {
				return true
			}
		} else if pathVis[it] {
			// visited in current path → cycle
			return true
		}
	}

	// remove from current path
	pathVis[node] = false
	return false
}

// ---------- Detect Cycle ----------
func (s *Solution) isCyclic(V int, adj [][]int) bool {

	vis := make([]bool, V)
	pathVis := make([]bool, V)

	for i := 0; i < V; i++ {
		if !vis[i] {
			if s.dfsCheck(i, adj, vis, pathVis) {
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
		{3},
		{4, 7},
		{5},
		{6},
		{},
		{5},
		{9},
		{10},
		{8},
	}

	V := len(adj)

	obj := Solution{}
	ans := obj.isCyclic(V, adj)

	if ans {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

// Time Complexity: O(V+E)+O(V) , where V = no. of nodes and E = no. of edges. There can be at most V components. So, another O(V) time complexity.
// Space Complexity: O(2N) + O(N) ~ O(2N): O(2N) for two visited arrays and O(N) for recursive stack space.