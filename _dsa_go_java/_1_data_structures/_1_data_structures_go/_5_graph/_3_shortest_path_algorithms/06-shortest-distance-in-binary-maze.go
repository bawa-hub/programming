// https://leetcode.com/problems/shortest-path-in-binary-matrix/
// for leetcode see on leetcode

// https://practice.geeksforgeeks.org/problems/shortest-path-in-a-binary-maze-1655453161/1
// this is gfg solution rather concept is same with minor variation


package main

import (
	"fmt"
	"math"
)

type Pair struct {
	r int
	c int
}

type State struct {
	dist int
	cell Pair
}

func shortestPath(grid [][]int, source Pair, destination Pair) int {

	// Edge case
	if source.r == destination.r && source.c == destination.c {
		return 0
	}

	n := len(grid)
	m := len(grid[0])

	// distance matrix
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	queue := []State{}
	dist[source.r][source.c] = 0
	queue = append(queue, State{0, source})

	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	for len(queue) > 0 {

		cur := queue[0]
		queue = queue[1:]

		dis := cur.dist
		r := cur.cell.r
		c := cur.cell.c

		for i := 0; i < 4; i++ {

			newr := r + dr[i]
			newc := c + dc[i]

			if newr >= 0 && newr < n &&
				newc >= 0 && newc < m &&
				grid[newr][newc] == 1 &&
				dis+1 < dist[newr][newc] {

				dist[newr][newc] = dis + 1

				if newr == destination.r &&
					newc == destination.c {
					return dis + 1
				}

				queue = append(queue,
					State{dis + 1, Pair{newr, newc}})
			}
		}
	}

	return -1
}

func main() {

	source := Pair{0, 1}
	destination := Pair{2, 2}

	grid := [][]int{
		{1, 1, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
		{1, 1, 0, 0},
		{1, 0, 0, 1},
	}

	res := shortestPath(grid, source, destination)
	fmt.Println(res)
}

// Time Complexity: O( 4*N*M ) { N*M are the total cells, for each of which we also check 4 adjacent nodes for the shortest path length}, Where N = No. of rows of the binary maze and M = No. of columns of the binary maze.
// Space Complexity: O( N*M ), Where N = No. of rows of the binary maze and M = No. of columns of the binary maze.