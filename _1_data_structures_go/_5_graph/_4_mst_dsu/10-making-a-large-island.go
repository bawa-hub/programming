// https://leetcode.com/problems/making-a-large-island/

package main

import "fmt"

/*************** DISJOINT SET ****************/

type DisjointSet struct {
	parent []int
	size   []int
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		ds.parent[i] = i
		ds.size[i] = 1
	}
	return ds
}

func (ds *DisjointSet) FindUPar(node int) int {
	if node == ds.parent[node] {
		return node
	}
	ds.parent[node] = ds.FindUPar(ds.parent[node])
	return ds.parent[node]
}

func (ds *DisjointSet) UnionBySize(u, v int) {
	ulpU := ds.FindUPar(u)
	ulpV := ds.FindUPar(v)

	if ulpU == ulpV {
		return
	}

	if ds.size[ulpU] < ds.size[ulpV] {
		ds.parent[ulpU] = ulpV
		ds.size[ulpV] += ds.size[ulpU]
	} else {
		ds.parent[ulpV] = ulpU
		ds.size[ulpU] += ds.size[ulpV]
	}
}

/*************** SOLUTION ****************/

func isValid(r, c, n int) bool {
	return r >= 0 && r < n && c >= 0 && c < n
}

func largestIsland(grid [][]int) int {

	n := len(grid)
	ds := NewDisjointSet(n * n)

	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}

	// Step 1: connect existing islands
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {

			if grid[r][c] == 0 {
				continue
			}

			for i := 0; i < 4; i++ {
				nr := r + dr[i]
				nc := c + dc[i]

				if isValid(nr, nc, n) && grid[nr][nc] == 1 {
					node := r*n + c
					adj := nr*n + nc
					ds.UnionBySize(node, adj)
				}
			}
		}
	}

	maxIsland := 0

	// Step 2: try converting 0 → 1
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {

			if grid[r][c] == 1 {
				continue
			}

			componentSet := make(map[int]bool)

			for i := 0; i < 4; i++ {
				nr := r + dr[i]
				nc := c + dc[i]

				if isValid(nr, nc, n) && grid[nr][nc] == 1 {
					parent := ds.FindUPar(nr*n + nc)
					componentSet[parent] = true
				}
			}

			sizeTotal := 0
			for parent := range componentSet {
				sizeTotal += ds.size[parent]
			}

			if sizeTotal+1 > maxIsland {
				maxIsland = sizeTotal + 1
			}
		}
	}

	// Edge case: grid already full of 1s
	for i := 0; i < n*n; i++ {
		root := ds.FindUPar(i)
		if ds.size[root] > maxIsland {
			maxIsland = ds.size[root]
		}
	}

	return maxIsland
}

/*************** DRIVER ****************/

func main() {

	grid := [][]int{
		{1, 0},
		{0, 1},
	}

	fmt.Println("Largest Island:", largestIsland(grid))
}