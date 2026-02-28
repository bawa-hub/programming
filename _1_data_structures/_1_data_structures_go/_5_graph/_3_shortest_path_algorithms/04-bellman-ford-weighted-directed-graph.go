// helps to detect negative cycles, and dijkstra not works in negative cycle

// https://practice.geeksforgeeks.org/problems/distance-from-the-source-bellman-ford-algorithm/1

package main

import (
	"fmt"
	"math"
)

func bellmanFord(V int, edges [][]int, S int) []int {

	dist := make([]int, V)

	// initialize distances
	for i := 0; i < V; i++ {
		dist[i] = math.MaxInt32
	}
	dist[S] = 0

	// Relax edges V-1 times
	for i := 0; i < V-1; i++ {
		for _, e := range edges {
			u := e[0]
			v := e[1]
			wt := e[2]

			if dist[u] != math.MaxInt32 &&
				dist[u]+wt < dist[v] {
				dist[v] = dist[u] + wt
			}
		}
	}

	// Check negative cycle
	for _, e := range edges {
		u := e[0]
		v := e[1]
		wt := e[2]

		if dist[u] != math.MaxInt32 &&
			dist[u]+wt < dist[v] {
			return []int{-1}
		}
	}

	return dist
}

func main() {

	V := 6
	edges := [][]int{
		{3, 2, 6},
		{5, 3, 1},
		{0, 1, 5},
		{1, 5, -3},
		{1, 2, -2},
		{3, 4, -2},
		{2, 4, 3},
	}

	S := 0

	dist := bellmanFord(V, edges, S)

	for _, d := range dist {
		fmt.Print(d, " ")
	}
	fmt.Println()
}

// Time Complexity: O(V*E), where V = no. of vertices and E = no. of Edges.
// Space Complexity: O(V) for the distance array which stores the minimized distances.