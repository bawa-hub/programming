// Linear ordering of vertices such that if there is an edge between u & v,
// u appears before v in that ordering.

// it is only applicable to directd acyclic graph (DAG)

package main

import "fmt"

type Solution struct{}

func dfs(node int, vis []int, stack *[]int, adj [][]int) {

	vis[node] = 1

	for _, it := range adj[node] {
		if vis[it] == 0 {
			dfs(it, vis, stack, adj)
		}
	}

	// push after visiting children
	*stack = append(*stack, node)
}

func (s *Solution) topoSort(N int, adj [][]int) []int {

	vis := make([]int, N)
	stack := []int{}

	for i := 0; i < N; i++ {
		if vis[i] == 0 {
			dfs(i, vis, &stack, adj)
		}
	}

	// reverse stack
	topo := []int{}
	for i := len(stack) - 1; i >= 0; i-- {
		topo = append(topo, stack[i])
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
	adj[2] = append(adj[2], 3)
	adj[3] = append(adj[3], 1)

	obj := Solution{}
	res := obj.topoSort(N, adj)

	fmt.Println("Toposort of the given graph is:")
	for _, v := range res {
		fmt.Print(v, " ")
	}
}

// Time Complexity: O(V+E)+O(V), where V = no. of nodes and E = no. of edges. There can be at most V components. So, another O(V) time complexity.
// Space Complexity: O(2N) + O(N) ~ O(2N): O(2N) for the visited array and the stack carried during DFS calls and O(N) for recursive stack space, where N = no. of nodes.