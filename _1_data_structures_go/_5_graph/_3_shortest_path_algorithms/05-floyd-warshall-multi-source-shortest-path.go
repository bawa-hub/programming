package main

import (
	"fmt"
	"math"
)

func shortestDistance(matrix [][]int) {

	n := len(matrix)
	INF := math.MaxInt32 / 2 // prevent overflow

	// preprocessing
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if matrix[i][j] == -1 {
				matrix[i][j] = INF
			}
			if i == j {
				matrix[i][j] = 0
			}
		}
	}

	// Floyd Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {

				if matrix[i][k]+matrix[k][j] < matrix[i][j] {
					matrix[i][j] = matrix[i][k] + matrix[k][j]
				}
			}
		}
	}

	// restore unreachable nodes
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == INF {
				matrix[i][j] = -1
			}
		}
	}
}

func main() {

	V := 4
	matrix := make([][]int, V)

	for i := range matrix {
		matrix[i] = make([]int, V)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	matrix[0][1] = 2
	matrix[1][0] = 1
	matrix[1][2] = 3
	matrix[3][0] = 3
	matrix[3][1] = 5
	matrix[3][2] = 4

	shortestDistance(matrix)

	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

// Time Complexity: O(V3), as we have three nested loops each running for V times, where V = no. of vertices.
// Space Complexity: O(V2), where V = no. of vertices. This space complexity is due to storing the adjacency matrix of the given graph.