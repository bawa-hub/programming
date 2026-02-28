package main

import "fmt"

type Solution struct{}

// ---------- BFS Cycle Detection ----------
func (s *Solution) detect(src int, adj [][]int, vis []bool) bool {

	vis[src] = true

	// {node, parent}
	queue := [][2]int{{src, -1}}

	for len(queue) > 0 {

		p := queue[0]
		queue = queue[1:]

		node := p[0]
		parent := p[1]

		for _, adjacentNode := range adj[node] {

			// unvisited neighbour
			if !vis[adjacentNode] {
				vis[adjacentNode] = true
				queue = append(queue, [2]int{adjacentNode, node})
			} else if adjacentNode != parent {
				// visited but not parent → cycle
				return true
			}
		}
	}

	return false
}

// ---------- Main Cycle Function ----------
func (s *Solution) isCycle(V int, adj [][]int) bool {

	vis := make([]bool, V)

	for i := 0; i < V; i++ {
		if !vis[i] {
			if s.detect(i, adj, vis) {
				return true
			}
		}
	}

	return false
}

func main() {

	adj := [][]int{
		{},
		{2},
		{1, 3},
		{2},
	}

	obj := Solution{}
	ans := obj.isCycle(4, adj)

	if ans {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}


// Time Complexity: O(N + 2E) + O(N), Where N = Nodes, 2E is for total degrees as we traverse all adjacent nodes. In the case of connected components of a graph, it will take another O(N) time.
// Space Complexity: O(N) + O(N) ~ O(N), Space for queue data structure and visited array.