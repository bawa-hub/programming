// https://practice.geeksforgeeks.org/problems/number-of-provinces/1

package main

import "fmt"

/*************** DISJOINT SET ****************/

type DisjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for i := 0; i < n; i++ {
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

/*************** SOLUTION ****************/

func numProvinces(adj [][]int, V int) int {

	ds := NewDisjointSet(V)

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if adj[i][j] == 1 {
				ds.UnionByRank(i, j)
			}
		}
	}

	count := 0
	for i := 0; i < V; i++ {
		if ds.FindUPar(i) == i {
			count++
		}
	}

	return count
}

/*************** DRIVER ****************/

func main() {

	adj := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	V := len(adj)

	fmt.Println("Number of Provinces:", numProvinces(adj, V))
}