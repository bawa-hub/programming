package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// adjacency list
	adj := make([][]int, n+1)

	// for undirected graph
    // time complexity: O(2E)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// for directed graph
    // time complexity: O(E)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Space Comlexity - O(E)

	fmt.Println(adj)
}