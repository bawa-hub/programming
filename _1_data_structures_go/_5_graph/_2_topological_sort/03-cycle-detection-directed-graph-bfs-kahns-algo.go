// https://takeuforward.org/data-structure/detect-a-cycle-in-directed-graph-topological-sort-kahns-algorithm-g-23/

package main

import "fmt"

type Solution struct{}

func (s *Solution) isCyclic(V int, adj [][]int) bool {

	indegree := make([]int, V)

	// calculate indegree
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

	cnt := 0

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		cnt++

		for _, it := range adj[node] {
			indegree[it]--
			if indegree[it] == 0 {
				queue = append(queue, it)
			}
		}
	}

	// if topo sort doesn't include all nodes → cycle exists
	return cnt != V
}

func main() {

	V := 6
	adj := [][]int{
		{},
		{2},
		{3},
		{4, 5},
		{2},
		{},
	}

	obj := Solution{}
	ans := obj.isCyclic(V, adj)

	if ans {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}


// Time Complexity: O(V+E), where V = no. of nodes and E = no. of edges. This is a simple BFS algorithm.
// Space Complexity: O(N) + O(N) ~ O(2N), O(N) for the in-degree array, and O(N) for the queue data structure used in BFS(where N = no.of nodes).