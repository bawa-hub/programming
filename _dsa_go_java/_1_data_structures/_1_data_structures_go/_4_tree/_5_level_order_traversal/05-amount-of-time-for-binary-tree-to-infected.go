// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/

// same as burning tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	parent := make(map[*TreeNode]*TreeNode)

	target := bfsToMapParents(root, parent, start)

	return findMaxDistance(parent, target)
}

// ---------- Build Parent Map ----------
func bfsToMapParents(
	root *TreeNode,
	parent map[*TreeNode]*TreeNode,
	start int,
) *TreeNode {

	queue := []*TreeNode{root}
	var target *TreeNode

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val == start {
			target = node
		}

		if node.Left != nil {
			parent[node.Left] = node
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			parent[node.Right] = node
			queue = append(queue, node.Right)
		}
	}

	return target
}

// ---------- BFS Infection Spread ----------
func findMaxDistance(
	parent map[*TreeNode]*TreeNode,
	target *TreeNode,
) int {

	queue := []*TreeNode{target}
	visited := make(map[*TreeNode]bool)
	visited[target] = true

	time := 0

	for len(queue) > 0 {
		size := len(queue)
		spread := false

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil && !visited[node.Left] {
				visited[node.Left] = true
				queue = append(queue, node.Left)
				spread = true
			}

			if node.Right != nil && !visited[node.Right] {
				visited[node.Right] = true
				queue = append(queue, node.Right)
				spread = true
			}

			if p, ok := parent[node]; ok && !visited[p] {
				visited[p] = true
				queue = append(queue, p)
				spread = true
			}
		}

		if spread {
			time++
		}
	}

	return time
}