// https://practice.geeksforgeeks.org/problems/minimum-spanning-tree/1

package main

import (
	"fmt"
	"sort"
)

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

/*************** KRUSKAL ****************/

type Edge struct {
	wt int
	u  int
	v  int
}

func spanningTree(V int, adj [][][2]int) int {

	var edges []Edge

	// Build edge list
	for node := 0; node < V; node++ {
		for _, it := range adj[node] {
			adjNode := it[0]
			wt := it[1]
			edges = append(edges, Edge{wt, node, adjNode})
		}
	}

	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].wt < edges[j].wt
	})

	ds := NewDisjointSet(V)

	mstWt := 0

	for _, e := range edges {
		if ds.FindUPar(e.u) != ds.FindUPar(e.v) {
			mstWt += e.wt
			ds.UnionBySize(e.u, e.v)
		}
	}

	return mstWt
}

/*************** DRIVER ****************/

func main() {

	V := 5

	edges := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 2, 1},
		{2, 3, 2},
		{3, 4, 1},
		{4, 2, 2},
	}

	adj := make([][][2]int, V)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]

		adj[u] = append(adj[u], [2]int{v, w})
		adj[v] = append(adj[v], [2]int{u, w})
	}

	mstWt := spanningTree(V, adj)

	fmt.Println("Sum of MST weights:", mstWt)
}