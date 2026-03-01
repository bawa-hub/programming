// https://leetcode.com/problems/critical-connections-in-a-network/

package main

import "fmt"

type Solution struct {
	timer int
}

func (s *Solution) dfs(
	node int,
	parent int,
	vis []int,
	adj [][]int,
	tin []int,
	low []int,
	bridges *[][]int,
) {

	vis[node] = 1
	tin[node] = s.timer
	low[node] = s.timer
	s.timer++

	for _, it := range adj[node] {

		if it == parent {
			continue
		}

		if vis[it] == 0 {

			s.dfs(it, node, vis, adj, tin, low, bridges)

			low[node] = min(low[node], low[it])

			// bridge condition
			if low[it] > tin[node] {
				*bridges = append(*bridges, []int{node, it})
			}

		} else {
			low[node] = min(low[node], tin[it])
		}
	}
}

func (s *Solution) criticalConnections(
	n int,
	connections [][]int,
) [][]int {

	adj := make([][]int, n)

	for _, e := range connections {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	vis := make([]int, n)
	tin := make([]int, n)
	low := make([]int, n)

	bridges := [][]int{}
	s.timer = 1

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			s.dfs(i, -1, vis, adj, tin, low, &bridges)
		}
	}

	return bridges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	n := 4
	connections := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
		{1, 3},
	}

	obj := Solution{}
	bridges := obj.criticalConnections(n, connections)

	for _, b := range bridges {
		fmt.Printf("[%d, %d] ", b[0], b[1])
	}
	fmt.Println()
}