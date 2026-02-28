// https://leetcode.com/problems/01-matrix/

package main

import "fmt"

type Cell struct {
	r, c int
	dist int
}

type Solution struct{}

func (s *Solution) nearest(grid [][]int) [][]int {

	n := len(grid)
	m := len(grid[0])

	vis := make([][]int, n)
	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
		dist[i] = make([]int, m)
	}

	queue := []Cell{}

	// push all 0-cells (multi-source BFS)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				queue = append(queue, Cell{i, j, 0})
				vis[i][j] = 1
			}
		}
	}

	delRow := []int{-1, 0, 1, 0}
	delCol := []int{0, 1, 0, -1}

	for len(queue) > 0 {

		cell := queue[0]
		queue = queue[1:]

		row, col, steps := cell.r, cell.c, cell.dist
		dist[row][col] = steps

		for i := 0; i < 4; i++ {
			nrow := row + delRow[i]
			ncol := col + delCol[i]

			if nrow >= 0 && nrow < n &&
				ncol >= 0 && ncol < m &&
				vis[nrow][ncol] == 0 {

				vis[nrow][ncol] = 1
				queue = append(queue, Cell{nrow, ncol, steps + 1})
			}
		}
	}

	return dist
}

func main() {

	grid := [][]int{
		{0, 1, 1, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 1},
	}

	obj := Solution{}
	ans := obj.nearest(grid)

	for _, row := range ans {
		for _, v := range row {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

// Time Complexity: O(NxM + NxMx4) ~ O(N x M)
// For the worst case, the BFS function will be called for (N x M) nodes, and for every node, we are traversing for 4 neighbors, so it will take O(N x M x 4) time.
// Space Complexity: O(N x M) + O(N x M) + O(N x M) ~ O(N x M)
// O(N x M) for the visited array, distance matrix, and queue space takes up N x M locations at max.