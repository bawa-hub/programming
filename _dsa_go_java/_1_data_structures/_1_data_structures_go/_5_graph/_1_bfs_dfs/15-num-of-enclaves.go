// https://leetcode.com/problems/number-of-enclaves/description/

package main

type Solution struct {
	rows int
	cols int
}

func (s *Solution) dfs(grid [][]int, i, j int) {

	if i < 0 || j < 0 || i >= s.rows || j >= s.cols {
		return
	}

	if grid[i][j] != 1 {
		return
	}

	// mark reachable land from boundary
	grid[i][j] = -1

	s.dfs(grid, i+1, j)
	s.dfs(grid, i-1, j)
	s.dfs(grid, i, j+1)
	s.dfs(grid, i, j-1)
}

func (s *Solution) numEnclaves(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	s.rows = len(grid)
	s.cols = len(grid[0])

	// DFS from boundary
	for i := 0; i < s.rows; i++ {
		for j := 0; j < s.cols; j++ {
			if i == 0 || j == 0 || i == s.rows-1 || j == s.cols-1 {
				s.dfs(grid, i, j)
			}
		}
	}

	ans := 0
	for i := 0; i < s.rows; i++ {
		for j := 0; j < s.cols; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}

	return ans
}

// bfs
package main

type Cell struct {
	r int
	c int
}

func numberOfEnclaves(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])

	vis := make([][]int, n)
	for i := range vis {
		vis[i] = make([]int, m)
	}

	queue := []Cell{}

	// push boundary land cells
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			if i == 0 || j == 0 || i == n-1 || j == m-1 {
				if grid[i][j] == 1 {
					queue = append(queue, Cell{i, j})
					vis[i][j] = 1
				}
			}
		}
	}

	delRow := []int{-1, 0, 1, 0}
	delCol := []int{0, 1, 0, -1}

	// BFS
	for len(queue) > 0 {

		cell := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nrow := cell.r + delRow[i]
			ncol := cell.c + delCol[i]

			if nrow >= 0 && nrow < n &&
				ncol >= 0 && ncol < m &&
				vis[nrow][ncol] == 0 &&
				grid[nrow][ncol] == 1 {

				vis[nrow][ncol] = 1
				queue = append(queue, Cell{nrow, ncol})
			}
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && vis[i][j] == 0 {
				cnt++
			}
		}
	}

	return cnt
}