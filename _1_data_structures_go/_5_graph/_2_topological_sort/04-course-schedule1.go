// https://leetcode.com/problems/course-schedule/

package main

import "fmt"

type Solution struct{}

func (s *Solution) isPossible(V int, prerequisites [][2]int) bool {

	// adjacency list
	adj := make([][]int, V)

	for _, it := range prerequisites {
		u := it[0]
		v := it[1]
		adj[u] = append(adj[u], v)
	}

	// indegree array
	indegree := make([]int, V)

	for i := 0; i < V; i++ {
		for _, it := range adj[i] {
			indegree[it]++
		}
	}

	queue := []int{}

	// push nodes with indegree 0
	for i := 0; i < V; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topoCount := 0

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		topoCount++

		for _, it := range adj[node] {
			indegree[it]--
			if indegree[it] == 0 {
				queue = append(queue, it)
			}
		}
	}

	return topoCount == V
}

func main() {

	prerequisites := [][2]int{
		{1, 0},
		{2, 1},
		{3, 2},
	}

	N := 4

	obj := Solution{}
	ans := obj.isPossible(N, prerequisites)

	if ans {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}