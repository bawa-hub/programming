// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/

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

func maxRemove(stones [][]int, n int) int {

	maxRow := 0
	maxCol := 0

	for _, s := range stones {
		if s[0] > maxRow {
			maxRow = s[0]
		}
		if s[1] > maxCol {
			maxCol = s[1]
		}
	}

	ds := NewDisjointSet(maxRow + maxCol + 2)

	stoneNodes := make(map[int]bool)

	for _, s := range stones {
		rowNode := s[0]
		colNode := s[1] + maxRow + 1

		ds.UnionBySize(rowNode, colNode)

		stoneNodes[rowNode] = true
		stoneNodes[colNode] = true
	}

	components := 0

	for node := range stoneNodes {
		if ds.FindUPar(node) == node {
			components++
		}
	}

	return n - components
}

/*************** DRIVER ****************/

func main() {

	stones := [][]int{
		{0, 0},
		{0, 2},
		{1, 3},
		{3, 1},
		{3, 2},
		{4, 3},
	}

	n := len(stones)

	ans := maxRemove(stones, n)

	fmt.Println("Maximum stones removed:", ans)
}