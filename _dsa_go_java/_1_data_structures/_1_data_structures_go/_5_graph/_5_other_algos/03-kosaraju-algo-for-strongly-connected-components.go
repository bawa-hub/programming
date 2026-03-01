// https://practice.geeksforgeeks.org/problems/strongly-connected-components-kosarajus-algo/1

package main

import (
	"fmt"
)

type Solution struct{}

func (s *Solution) dfs(
	node int,
	vis []bool,
	adj [][]int,
	st *[]int,
) {
	vis[node] = true

	for _, it := range adj[node] {
		if !vis[it] {
			s.dfs(it, vis, adj, st)
		}
	}

	*st = append(*st, node)
}

func (s *Solution) dfsTranspose(
	node int,
	vis []bool,
	adjT [][]int,
) {
	vis[node] = true

	for _, it := range adjT[node] {
		if !vis[it] {
			s.dfsTranspose(it, vis, adjT)
		}
	}
}

func (s *Solution) kosaraju(V int, adj [][]int) int {

	vis := make([]bool, V)
	stack := []int{}

	// Step 1: order by finish time
	for i := 0; i < V; i++ {
		if !vis[i] {
			s.dfs(i, vis, adj, &stack)
		}
	}

	// Step 2: transpose graph
	adjT := make([][]int, V)
	for i := 0; i < V; i++ {
		for _, it := range adj[i] {
			adjT[it] = append(adjT[it], i)
		}
	}

	// reset visited
	for i := range vis {
		vis[i] = false
	}

	// Step 3: count SCC
	scc := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !vis[node] {
			scc++
			s.dfsTranspose(node, vis, adjT)
		}
	}

	return scc
}

func main() {

	n := 5
	edges := [][]int{
		{1, 0},
		{0, 2},
		{2, 1},
		{0, 3},
		{3, 4},
	}

	adj := make([][]int, n)

	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}

	obj := Solution{}
	ans := obj.kosaraju(n, adj)

	fmt.Println("Number of SCC:", ans)
}

// Time Complexity: O(V+E) + O(V+E) + O(V+E) ~ O(V+E) , where V = no. of vertices, E = no. of edges. The first step is a simple DFS, so the first term is O(V+E). The second step of reversing the graph and the third step, containing DFS again, will take O(V+E) each.
// Space Complexity: O(V)+O(V)+O(V+E), where V = no. of vertices, E = no. of edges. Two O(V) for the visited array and the stack we have used. O(V+E) space for the reversed adjacent list.