package main

import "fmt"

type Solution struct{}

// 4 steps to understand dfs

// ---------- DFS ----------
func (s *Solution) dfs(node int, adj [][]int, vis []bool, res *[]int) {
	// 1. Action on entering node
	vis[node] = true
	*res = append(*res, node)

	for _, it := range adj[node] {
		// 2. Before entering child
		if !vis[it] {
			s.dfs(it, adj, vis, res)
		}
		// 3. After exiting child (optional place)
	}

	// 4. After exiting node (optional place)
}

// ---------- DFS Of Graph ----------
func (s *Solution) dfsOfGraph(V int, adj [][]int) []int {
	vis := make([]bool, V)
	res := []int{}

	start := 0
	s.dfs(start, adj, vis, &res)

	return res
}

// ---------- Add Edge ----------
func addEdge(adj [][]int, u, v int) {
	adj[u] = append(adj[u], v)
	adj[v] = append(adj[v], u)
}

func printAns(ans []int) {
	for _, v := range ans {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	adj := make([][]int, n)

	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)
		addEdge(adj, v1, v2)
	}

	obj := Solution{}
	ans := obj.dfsOfGraph(n, adj)

	printAns(ans)
}

// Time Complexity: For an undirected graph, O(N) + O(2E), For a directed graph, O(N) + O(E), 
// Because for every node we are calling the recursive function once, the time taken is O(N) and 2E is for total degrees as we traverse for all adjacent nodes.
// Space Complexity: O(3N) ~ O(N), Space for dfs stack space, visited array and an adjacency list.


// Time complexity explanation

// Step 1: Each Node Is Visited Once → O(V)
// Step 2: Total Work Done in All Loops → O(2E)
// for (auto it : adj[node]) -> runs once for each neighbor of each node.

// If you sum across the whole graph:
// Node 1 loop runs degree(1) times
// Node 2 loop runs degree(2) times
// Node 3 loop runs degree(3) times
// ...
// Node V loop runs degree(V) times

// If you add all degrees:
// degree(1) + degree(2) + ... + degree(V) = 2E (for undirected graph)


// ✔ Why it's not O(V × E)

// Because DFS does not check all edges for each node.
// It only checks the edges that belong to that particular node.

// Total time:
// Visiting nodes = O(V)
// Checking all adj lists = O(2E)
// Total = O(V + E)