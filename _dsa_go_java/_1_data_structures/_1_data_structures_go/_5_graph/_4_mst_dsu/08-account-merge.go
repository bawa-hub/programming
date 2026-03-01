// https://leetcode.com/problems/accounts-merge/

package main

import (
	"fmt"
	"sort"
)

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

func accountsMerge(details [][]string) [][]string {

	n := len(details)
	ds := NewDisjointSet(n)

	sort.Slice(details, func(i, j int) bool {
		return details[i][0] < details[j][0]
	})

	mailToNode := make(map[string]int)

	// Union accounts sharing same email
	for i := 0; i < n; i++ {
		for j := 1; j < len(details[i]); j++ {
			mail := details[i][j]

			if node, ok := mailToNode[mail]; !ok {
				mailToNode[mail] = i
			} else {
				ds.UnionBySize(i, node)
			}
		}
	}

	mergedMail := make([][]string, n)

	// group mails by ultimate parent
	for mail, node := range mailToNode {
		parent := ds.FindUPar(node)
		mergedMail[parent] = append(mergedMail[parent], mail)
	}

	var ans [][]string

	for i := 0; i < n; i++ {
		if len(mergedMail[i]) == 0 {
			continue
		}

		sort.Strings(mergedMail[i])

		temp := []string{details[i][0]}
		temp = append(temp, mergedMail[i]...)
		ans = append(ans, temp)
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})

	return ans
}

/*************** DRIVER ****************/

func main() {

	accounts := [][]string{
		{"John", "j1@com", "j2@com", "j3@com"},
		{"John", "j4@com"},
		{"Raj", "r1@com", "r2@com"},
		{"John", "j1@com", "j5@com"},
		{"Raj", "r2@com", "r3@com"},
		{"Mary", "m1@com"},
	}

	ans := accountsMerge(accounts)

	for _, acc := range ans {
		fmt.Print(acc[0], ": ")
		for i := 1; i < len(acc); i++ {
			fmt.Print(acc[i], " ")
		}
		fmt.Println()
	}
}