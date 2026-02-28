// https://leetcode.com/problems/cheapest-flights-within-k-stops/

package main

import (
	"fmt"
	"math"
)

type Pair struct {
	node int
	cost int
}

type State struct {
	stops int
	node  int
	cost  int
}

func CheapestFlight(n int, flights [][]int, src int, dst int, K int) int {

	// adjacency list
	adj := make([][]Pair, n)
	for _, f := range flights {
		u := f[0]
		v := f[1]
		w := f[2]
		adj[u] = append(adj[u], Pair{v, w})
	}

	// queue for BFS
	queue := []State{{0, src, 0}}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0

	for len(queue) > 0 {

		cur := queue[0]
		queue = queue[1:]

		stops := cur.stops
		node := cur.node
		cost := cur.cost

		if stops > K {
			continue
		}

		for _, next := range adj[node] {

			newCost := cost + next.cost

			if newCost < dist[next.node] && stops <= K {
				dist[next.node] = newCost
				queue = append(queue,
					State{stops + 1, next.node, newCost})
			}
		}
	}

	if dist[dst] == math.MaxInt32 {
		return -1
	}

	return dist[dst]
}

func main() {

	n := 4
	src := 0
	dst := 3
	K := 1

	flights := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}

	ans := CheapestFlight(n, flights, src, dst, K)
	fmt.Println(ans)
}

// Time Complexity: O( N ) { Additional log(N) of time eliminated here because we’re using a simple queue rather than a priority queue which is usually used in Dijkstra’s Algorithm }.
// Where N = Number of flights / Number of edges.
// Space Complexity:  O( |E| + |V| ) { for the adjacency list, priority queue, and the dist array }.
// Where E = Number of edges (flights.size()) and V = Number of Airports.