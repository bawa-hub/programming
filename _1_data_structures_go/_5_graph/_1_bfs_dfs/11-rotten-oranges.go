// https://leetcode.com/problems/rotting-oranges/

package main

import "fmt"

type Cell struct {
	r int
	c int
	t int
}

type Solution struct{}

func (s *Solution) orangesRotting(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])

	queue := []Cell{}
	vis := make([][]int, n)

	for i := range vis {
		vis[i] = make([]int, m)
	}

	cntFresh := 0

	// initialize queue with rotten oranges
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			if grid[i][j] == 2 {
				queue = append(queue, Cell{i, j, 0})
				vis[i][j] = 2
			} else {
				vis[i][j] = 0
			}

			if grid[i][j] == 1 {
				cntFresh++
			}
		}
	}

	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, 1, 0, -1}

	tm := 0
	cnt := 0

	// BFS
	for len(queue) > 0 {

		cell := queue[0]
		queue = queue[1:]

		r, c, t := cell.r, cell.c, cell.t
		if t > tm {
			tm = t
		}

		for i := 0; i < 4; i++ {
			nrow := r + drow[i]
			ncol := c + dcol[i]

			if nrow >= 0 && nrow < n &&
				ncol >= 0 && ncol < m &&
				vis[nrow][ncol] == 0 &&
				grid[nrow][ncol] == 1 {

				queue = append(queue, Cell{nrow, ncol, t + 1})
				vis[nrow][ncol] = 2
				cnt++
			}
		}
	}

	if cnt != cntFresh {
		return -1
	}

	return tm
}

func main() {

	grid := [][]int{
		{0, 1, 2},
		{0, 1, 2},
		{2, 1, 1},
	}

	obj := Solution{}
	fmt.Println(obj.orangesRotting(grid))
}

// Time Complexity: O ( n x n ) x 4    
// Reason: Worst-case – We will be making each fresh orange rotten in the grid and for each rotten orange will check in 4 directions
// Space Complexity: O ( n x n )
// Reason: worst-case –  If all oranges are Rotten, we will end up pushing all rotten oranges into the Queue data structure 