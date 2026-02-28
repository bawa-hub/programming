// https://leetcode.com/problems/find-eventual-safe-states/
// https://takeuforward.org/data-structure/find-eventual-safe-states-dfs-g-20/

package main

import "fmt"

type Solution struct{}

func (s *Solution) dfsCheck(
	node int,
	adj [][]int,
	vis []bool,
	pathVis []bool,
	check []bool,
) bool {

	vis[node] = true
	pathVis[node] = true
	check[node] = false

	for _, it := range adj[node] {

		if !vis[it] {
			if s.dfsCheck(it, adj, vis, pathVis, check) {
				check[node] = false
				return true
			}
		} else if pathVis[it] {
			return true
		}
	}

	check[node] = true
	pathVis[node] = false
	return false
}

func (s *Solution) eventualSafeNodes(V int, adj [][]int) []int {

	vis := make([]bool, V)
	pathVis := make([]bool, V)
	check := make([]bool, V)

	for i := 0; i < V; i++ {
		if !vis[i] {
			s.dfsCheck(i, adj, vis, pathVis, check)
		}
	}

	safeNodes := []int{}
	for i := 0; i < V; i++ {
		if check[i] {
			safeNodes = append(safeNodes, i)
		}
	}

	return safeNodes
}

func main() {

	adj := [][]int{
		{1},
		{2},
		{3},
		{4, 5},
		{6},
		{6},
		{7},
		{},
		{1, 9},
		{10},
		{8},
		{9},
	}

	V := len(adj)

	obj := Solution{}
	safeNodes := obj.eventualSafeNodes(V, adj)

	for _, node := range safeNodes {
		fmt.Print(node, " ")
	}
	fmt.Println()
}