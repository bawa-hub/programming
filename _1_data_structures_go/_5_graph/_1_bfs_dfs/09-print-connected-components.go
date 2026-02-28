package main

import "fmt"

var vis []bool
var g [][]int
var cc [][]int
var currentCC []int

// ---------- DFS ----------
func dfs(vertex int) {
	vis[vertex] = true
	currentCC = append(currentCC, vertex)

	for _, child := range g[vertex] {
		if vis[child] {
			continue
		}
		dfs(child)
	}
}

func main() {

	var n, e int
	fmt.Scan(&n, &e)

	// adjacency list (1-based indexing)
	g = make([][]int, n+1)
	vis = make([]bool, n+1)

	for i := 0; i < e; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}

		currentCC = []int{}
		dfs(i)
		cc = append(cc, currentCC)
	}

	fmt.Println(len(cc))

	for _, component := range cc {
		for _, vertex := range component {
			fmt.Print(vertex, " ")
		}
		fmt.Println()
	}
}