package main

import "fmt"

type Solution struct{}

// ---------- BFS ----------
func (s *Solution) bfsOfGraph(V int, adj [][]int) []int {

	visited := make([]bool, V)
	bfs := []int{}

	queue := []int{0}
	visited[0] = true

	for len(queue) > 0 {

		// pop front
		node := queue[0]
		queue = queue[1:]

		bfs = append(bfs, node)

		// visit neighbours
		for _, it := range adj[node] {
			if !visited[it] {
				visited[it] = true
				queue = append(queue, it)
			}
		}
	}

	return bfs
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

	V := 5
	adj := make([][]int, V)

	addEdge(adj, 0, 1)
	addEdge(adj, 1, 2)
	addEdge(adj, 1, 3)
	addEdge(adj, 0, 4)

	obj := Solution{}
	ans := obj.bfsOfGraph(V, adj)

	printAns(ans)
}

// Time complexity: O(V+2E) -> same concept as DFS
// Spacce complexity: O(V)