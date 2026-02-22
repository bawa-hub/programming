package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// create adjacency matrix
	adj := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		adj[i] = make([]int, n+1)
	}

	// read edges
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		adj[u][v] = 1
		adj[v][u] = 1 // remove for directed graph
	}

	fmt.Println("Adjacency Matrix:")
	for i := 1; i <= n; i++ {
		fmt.Println(adj[i][1 : n+1])
	}
}