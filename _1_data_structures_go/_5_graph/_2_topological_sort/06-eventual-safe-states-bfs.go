// https://leetcode.com/problems/find-eventual-safe-states/

package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

func (s *Solution) eventualSafeNodes(V int, adj [][]int) []int {

	// reverse graph
	adjRev := make([][]int, V)
	indegree := make([]int, V)

	// build reverse graph
	for i := 0; i < V; i++ {
		for _, it := range adj[i] {
			adjRev[it] = append(adjRev[it], i)
			indegree[i]++
		}
	}

	queue := []int{}
	safeNodes := []int{}

	// terminal nodes
	for i := 0; i < V; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Kahn's BFS
	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		safeNodes = append(safeNodes, node)

		for _, it := range adjRev[node] {
			indegree[it]--
			if indegree[it] == 0 {
				queue = append(queue, it)
			}
		}
	}

	sort.Ints(safeNodes)
	return safeNodes
}

func main() {

	adj := [][]int{
		{1},
		{2},
		{3, 4},
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

	V := 12

	obj := Solution{}
	safeNodes := obj.eventualSafeNodes(V, adj)

	for _, node := range safeNodes {
		fmt.Print(node, " ")
	}
	fmt.Println()
}