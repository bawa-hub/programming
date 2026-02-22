package main

import "fmt"

func main() {
	// adjacency matrix
	adj := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}

	V := len(adj)

	// adjacency list
	adjList := make([][]int, V)

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {

			// ignore self loops
			if adj[i][j] == 1 && i != j {
				adjList[i] = append(adjList[i], j)
				adjList[j] = append(adjList[j], i)
			}
		}
	}

	fmt.Println("Adjacency List:")
	for i := 0; i < V; i++ {
		fmt.Println(i, "->", adjList[i])
	}
}