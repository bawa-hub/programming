// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/

package main

import (
	"fmt"
	"math"
)

func findCity(n int, edges [][]int, distanceThreshold int) int {

	// distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = math.MaxInt32
		}
		dist[i][i] = 0
	}

	// fill edges
	for _, e := range edges {
		u := e[0]
		v := e[1]
		w := e[2]

		dist[u][v] = w
		dist[v][u] = w
	}

	// Floyd–Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {

				if dist[i][k] == math.MaxInt32 ||
					dist[k][j] == math.MaxInt32 {
					continue
				}

				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// find city
	cntCity := n
	cityNo := -1

	for city := 0; city < n; city++ {
		cnt := 0
		for adjCity := 0; adjCity < n; adjCity++ {
			if dist[city][adjCity] <= distanceThreshold {
				cnt++
			}
		}

		// choose largest index if tie
		if cnt <= cntCity {
			cntCity = cnt
			cityNo = city
		}
	}

	return cityNo
}

func main() {

	n := 4
	edges := [][]int{
		{0, 1, 3},
		{1, 2, 1},
		{1, 3, 4},
		{2, 3, 1},
	}

	distanceThreshold := 4

	ans := findCity(n, edges, distanceThreshold)
	fmt.Println("The answer is node:", ans)
}

// Time Complexity: O(V3), as we have three nested loops each running for V times, where V = no. of vertices.
// Space Complexity: O(V2), where V = no. of vertices. This space complexity is due to storing the adjacency matrix of the given graph.