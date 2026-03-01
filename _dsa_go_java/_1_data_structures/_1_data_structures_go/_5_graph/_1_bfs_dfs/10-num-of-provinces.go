// https://leetcode.com/problems/number-of-provinces/
// https://practice.geeksforgeeks.org/problems/number-of-provinces/1

type Solution struct{}

func dfs(node int, adj [][]int, vis []bool) {
	vis[node] = true

	for _, it := range adj[node] {
		if !vis[it] {
			dfs(it, adj, vis)
		}
	}
}

func findCircleNum(isConnected [][]int) int {

	n := len(isConnected)

	// build adjacency list
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 && i != j {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	vis := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if !vis[i] {
			count++
			dfs(i, adj, vis)
		}
	}

	return count
}


// Time Complexity: O(N) + O(V+2E), Where O(N) is for outer loop and inner loop runs in total a single DFS over entire graph, and we know DFS takes a time of O(V+2E). 
// Space Complexity: O(N) + O(N),Space for recursion stack space and visited array.