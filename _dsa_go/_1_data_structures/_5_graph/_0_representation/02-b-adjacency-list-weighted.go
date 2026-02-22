package main

import "fmt"

type Edge struct {
	to int
	wt int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	adj := make([][]Edge, n+1)

	
	for i := 0; i < m; i++ {
		var u, v, wt int
		fmt.Scan(&u, &v, &wt)

		adj[u] = append(adj[u], Edge{to: v, wt: wt})
		adj[v] = append(adj[v], Edge{to: u, wt: wt}) // remove for directed
	}

	fmt.Println(adj)
}