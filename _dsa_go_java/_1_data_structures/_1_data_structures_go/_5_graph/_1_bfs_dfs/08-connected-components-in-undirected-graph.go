// https://www.hackerearth.com/problem/algorithm/connected-components-in-a-graph/
// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph

package main

import "fmt"

var vis []bool
var g [][]int

// ---------- DFS ----------
func dfs(vertex int) {
	vis[vertex] = true

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

	ct := 0

	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		ct++
		dfs(i)
	}

	fmt.Println(ct)
}