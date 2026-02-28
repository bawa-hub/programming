
// indegree === no of incoming edges to a node is called indegree of that node

package main

import "fmt"

type Solution struct{}

func (s *Solution) topo(N int, adj [][]int) []int {

	indegree := make([]int, N)

	// calculate indegree
	for i := 0; i < N; i++ {
		for _, it := range adj[i] {
			indegree[it]++
		}
	}

	queue := []int{}

	// push nodes having indegree 0
	for i := 0; i < N; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topo := []int{}

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		topo = append(topo, node)

		// remove node edges
		for _, it := range adj[node] {
			indegree[it]--
			if indegree[it] == 0 {
				queue = append(queue, it)
			}
		}
	}

	return topo
}

func main() {

	N := 6
	adj := make([][]int, N)

	adj[5] = append(adj[5], 2)
	adj[5] = append(adj[5], 0)
	adj[4] = append(adj[4], 0)
	adj[4] = append(adj[4], 1)
	adj[3] = append(adj[3], 1)
	adj[2] = append(adj[2], 3)

	obj := Solution{}
	res := obj.topo(N, adj)

	for _, v := range res {
		fmt.Print(v, " ")
	}
}


// Time Complexity: O(V+E), where V = no. of nodes and E = no. of edges. This is a simple BFS algorithm.
// Space Complexity: O(N) + O(N) ~ O(2N), O(N) for the indegree array, and O(N) for the queue data structure used in BFS(where N = no.of nodes).