// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ---------- Step 1: Mark Parents ----------
func markParents(root *TreeNode, parent map[*TreeNode]*TreeNode) {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left != nil {
			parent[current.Left] = current
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			parent[current.Right] = current
			queue = append(queue, current.Right)
		}
	}
}

// ---------- Step 2: BFS from target ----------
func distanceK(root, target *TreeNode, k int) []int {

	parentTrack := make(map[*TreeNode]*TreeNode)
	markParents(root, parentTrack)

	visited := make(map[*TreeNode]bool)
	queue := []*TreeNode{target}

	visited[target] = true
	currLevel := 0

	for len(queue) > 0 {

		if currLevel == k {
			break
		}

		size := len(queue)
		currLevel++

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			if current.Left != nil && !visited[current.Left] {
				visited[current.Left] = true
				queue = append(queue, current.Left)
			}

			if current.Right != nil && !visited[current.Right] {
				visited[current.Right] = true
				queue = append(queue, current.Right)
			}

			if parent, ok := parentTrack[current]; ok && !visited[parent] {
				visited[parent] = true
				queue = append(queue, parent)
			}
		}
	}

	result := []int{}
	for _, node := range queue {
		result = append(result, node.Val)
	}

	return result
}

// Time Complexity: O(2N + log N ) The time complexity arises from traversing the tree to create the parent hashmap, which involves visiting every node once hence O(N), exploring all nodes at a distance of ‘K’ which will be O(N) in the worst case, and the logarithmic lookup time for the hashmap is O( log N) in the worst scenario as well hence O(N + N + log N) which simplified to O(N).
// Space Complexity: O(N) The space complexity stems from the data structures used, O(N) for the parent hashmap, O(N) for the queue of DFS, and O(N) for the visited hashmap hence overall our space complexity is O(3N) ~ O(N).