// https://practice.geeksforgeeks.org/problems/alien-dictionary/1
// https://leetcode.ca/2016-08-25-269-Alien-Dictionary/

package main

import (
	"fmt"
)

type Solution struct{}

// ---------- Topological Sort ----------
func topoSort(K int, adj [][]int) []int {

	indegree := make([]int, K)

	for i := 0; i < K; i++ {
		for _, it := range adj[i] {
			indegree[it]++
		}
	}

	queue := []int{}

	for i := 0; i < K; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topo := []int{}

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		topo = append(topo, node)

		for _, it := range adj[node] {
			indegree[it]--
			if indegree[it] == 0 {
				queue = append(queue, it)
			}
		}
	}

	return topo
}

// ---------- Alien Dictionary ----------
func (s *Solution) findOrder(dict []string, N int, K int) string {

	adj := make([][]int, K)

	// build graph from dictionary
	for i := 0; i < N-1; i++ {

		s1 := dict[i]
		s2 := dict[i+1]

		length := len(s1)
		if len(s2) < length {
			length = len(s2)
		}

		for ptr := 0; ptr < length; ptr++ {
			if s1[ptr] != s2[ptr] {
				u := int(s1[ptr] - 'a')
				v := int(s2[ptr] - 'a')
				adj[u] = append(adj[u], v)
				break
			}
		}
	}

	topo := topoSort(K, adj)

	ans := ""
	for _, it := range topo {
		ans += string(rune(it + 'a'))
	}

	return ans
}

func main() {

	N := 5
	K := 4

	dict := []string{
		"baa",
		"abcd",
		"abca",
		"cab",
		"cad",
	}

	obj := Solution{}
	ans := obj.findOrder(dict, N, K)

	for _, ch := range ans {
		fmt.Print(string(ch), " ")
	}
	fmt.Println()
}