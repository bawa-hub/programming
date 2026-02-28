// https://leetcode.com/problems/number-of-operations-to-make-network-connected/
// https://practice.geeksforgeeks.org/problems/connecting-the-graph/1

package main

import "fmt"

/*************** DISJOINT SET ****************/

type DisjointSet struct {
	parent []int
	rank   []int
	size   []int
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		ds.parent[i] = i
		ds.size[i] = 1
	}
	return ds
}

// Path Compression
func (ds *DisjointSet) FindUPar(node int) int {
	if node == ds.parent[node] {
		return node
	}
	ds.parent[node] = ds.FindUPar(ds.parent[node])
	return ds.parent[node]
}

// Union by Size
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

func Solve(n int, edges [][]int) int {

	ds := NewDisjointSet(n)
	extraEdges := 0

	for _, e := range edges {
		u := e[0]
		v := e[1]

		// already connected → extra edge
		if ds.FindUPar(u) == ds.FindUPar(v) {
			extraEdges++
		} else {
			ds.UnionBySize(u, v)
		}
	}

	// count components
	components := 0
	for i := 0; i < n; i++ {
		if ds.FindUPar(i) == i {
			components++
		}
	}

	required := components - 1

	if extraEdges >= required {
		return required
	}

	return -1
}

/*************** DRIVER ****************/

func main() {

	V := 9
	edges := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{2, 3},
		{4, 5},
		{5, 6},
		{7, 8},
	}

	ans := Solve(V, edges)

	fmt.Println("Number of operations needed:", ans)
}