// https://leetcode.com/problems/number-of-islands/
// https://practice.geeksforgeeks.org/problems/find-the-number-of-islands

package main

import "fmt"

type Cell struct {
	r int
	c int
}

type Solution struct{}

func (s *Solution) bfs(row, col int, vis [][]int, grid [][]byte) {

	n := len(grid)
	m := len(grid[0])

	queue := []Cell{{row, col}}
	vis[row][col] = 1

	for len(queue) > 0 {

		cell := queue[0]
		queue = queue[1:]

		r := cell.r
		c := cell.c

		// 8 directions
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {

				nrow := r + dr
				ncol := c + dc

				if nrow >= 0 && nrow < n &&
					ncol >= 0 && ncol < m &&
					grid[nrow][ncol] == '1' &&
					vis[nrow][ncol] == 0 {

					vis[nrow][ncol] = 1
					queue = append(queue, Cell{nrow, ncol})
				}
			}
		}
	}
}

func (s *Solution) numIslands(grid [][]byte) int {

	n := len(grid)
	m := len(grid[0])

	vis := make([][]int, n)
	for i := range vis {
		vis[i] = make([]int, m)
	}

	cnt := 0

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {

			if grid[r][c] == '1' && vis[r][c] == 0 {
				cnt++
				s.bfs(r, c, vis, grid)
			}
		}
	}

	return cnt
}


// dfs
func dfs(grid [][]byte, r, c int, dr, dc []int) {

	grid[r][c] = '0'

	for i := 0; i < 4; i++ {

		nrow := r + dr[i]
		ncol := c + dc[i]

		if nrow >= 0 && nrow < len(grid) &&
			ncol >= 0 && ncol < len(grid[0]) &&
			grid[nrow][ncol] == '1' {

			dfs(grid, nrow, ncol, dr, dc)
		}
	}
}

func numIslandsDFS(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])

	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	cnt := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == '1' {
				cnt++
				dfs(grid, i, j, dr, dc)
			}
		}
	}

	return cnt
}