package main

import "fmt"

type DisjointSet struct {
	parent []int
	rank   []int
}

// Constructor
func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n+1),
		rank:   make([]int, n+1),
	}

	for i := 0; i <= n; i++ {
		ds.parent[i] = i
	}

	return ds
}

// Find with Path Compression
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

// -------- Driver --------

func main() {

	ds := NewDisjointSet(7)

	ds.UnionByRank(1, 2)
	ds.UnionByRank(2, 3)
	ds.UnionByRank(4, 5)
	ds.UnionByRank(6, 7)
	ds.UnionByRank(5, 6)

	// check if 3 and 7 belong to same set
	if ds.FindUPar(3) == ds.FindUPar(7) {
		fmt.Println("Same")
	} else {
		fmt.Println("Not same")
	}

	ds.UnionByRank(3, 7)

	if ds.FindUPar(3) == ds.FindUPar(7) {
		fmt.Println("Same")
	} else {
		fmt.Println("Not same")
	}
}