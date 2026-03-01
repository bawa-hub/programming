// https://leetcode.com/problems/course-schedule-ii/

package main

import "fmt"

type Solution struct{}

func (s *Solution) findOrder(V int, prerequisites [][]int) []int {

	// adjacency list
	adj := make([][]int, V)

	for _, it := range prerequisites {
		// it[1] -> it[0]
		adj[it[1]] = append(adj[it[1]], it[0])
	}

	// indegree array
	indegree := make([]int, V)

	for i := 0; i < V; i++ {
		for _, it := range adj[i] {
			indegree[it]++
		}
	}

	queue := []int{}

	// nodes having no prerequisites
	for i := 0; i < V; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topo := []int{}

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		topo = append(topo, node)

		for _, it := range adj[node] {
			indegree[it]--
			if indegree[it] == 0 {
				queue = append(queue, it)
			}
		}
	}

	// if cycle exists
	if len(topo) != V {
		return []int{}
	}

	return topo
}

func main() {

	N := 4

	prerequisites := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}

	obj := Solution{}
	ans := obj.findOrder(N, prerequisites)

	for _, task := range ans {
		fmt.Print(task, " ")
	}
	fmt.Println()
}