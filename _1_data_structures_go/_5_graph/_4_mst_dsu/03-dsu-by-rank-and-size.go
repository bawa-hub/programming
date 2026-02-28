// https://practice.geeksforgeeks.org/problems/disjoint-set-union-find/1

package main

import "fmt"

type DisjointSet struct {
	parent []int
	rank   []int
	size   []int
}

// Constructor
func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n+1),
		rank:   make([]int, n+1),
		size:   make([]int, n+1),
	}

	for i := 0; i <= n; i++ {
		ds.parent[i] = i
		ds.size[i] = 1
	}

	return ds
}

// Find Ultimate Parent (Path Compression)
func (ds *DisjointSet) FindUPar(node int) int {
	if node == ds.parent[node] {
		return node
	}
	ds.parent[node] = ds.FindUPar(ds.parent[node])
	return ds.parent[node]
}

// Union by Rank
func (ds *DisjointSet) UnionByRank(u, v int) {
	ulpU := ds.FindUPar(u)
	ulpV := ds.FindUPar(v)

	if ulpU == ulpV {
		return
	}

	if ds.rank[ulpU] < ds.rank[ulpV] {
		ds.parent[ulpU] = ulpV
	} else if ds.rank[ulpV] < ds.rank[ulpU] {
		ds.parent[ulpV] = ulpU
	} else {
		ds.parent[ulpV] = ulpU
		ds.rank[ulpU]++
	}
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

// -------- Driver --------

func main() {

	ds := NewDisjointSet(7)

	ds.UnionBySize(1, 2)
	ds.UnionBySize(2, 3)
	ds.UnionBySize(4, 5)
	ds.UnionBySize(6, 7)
	ds.UnionBySize(5, 6)

	if ds.FindUPar(3) == ds.FindUPar(7) {
		fmt.Println("Same")
	} else {
		fmt.Println("Not same")
	}

	ds.UnionBySize(3, 7)

	if ds.FindUPar(3) == ds.FindUPar(7) {
		fmt.Println("Same")
	} else {
		fmt.Println("Not same")
	}
}