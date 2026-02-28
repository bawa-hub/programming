// https://practice.geeksforgeeks.org/problems/number-of-distinct-islands/1
// https://leetcode.com/problems/number-of-distinct-islands/description/
// https://leetcode.ca/2017-10-24-694-Number-of-Distinct-Islands/

package main

import "fmt"

type Pair struct {
	r int
	c int
}

type Solution struct{}

func (s *Solution) dfs(
	row, col int,
	vis [][]bool,
	grid [][]int,
	vec *[]Pair,
	row0, col0 int,
) {

	vis[row][col] = true

	// store relative position
	*vec = append(*vec, Pair{row - row0, col - col0})

	n := len(grid)
	m := len(grid[0])

	delRow := []int{-1, 0, 1, 0}
	delCol := []int{0, -1, 0, 1}

	for i := 0; i < 4; i++ {
		nrow := row + delRow[i]
		ncol := col + delCol[i]

		if nrow >= 0 && nrow < n &&
			ncol >= 0 && ncol < m &&
			!vis[nrow][ncol] &&
			grid[nrow][ncol] == 1 {

			s.dfs(nrow, ncol, vis, grid, vec, row0, col0)
		}
	}
}

func (s *Solution) countDistinctIslands(grid [][]int) int {

	n := len(grid)
	m := len(grid[0])

	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}

	shapeSet := make(map[string]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			if !vis[i][j] && grid[i][j] == 1 {

				vec := []Pair{}
				s.dfs(i, j, vis, grid, &vec, i, j)

				// serialize shape
				key := ""
				for _, p := range vec {
					key += fmt.Sprintf("%d,%d|", p.r, p.c)
				}

				shapeSet[key] = true
			}
		}
	}

	return len(shapeSet)
}

func main() {

	grid := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}

	obj := Solution{}
	fmt.Println(obj.countDistinctIslands(grid))
}