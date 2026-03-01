// https://practice.geeksforgeeks.org/problems/burning-tree/1

package main

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// ---------- Find Maximum Distance (Burn Time) ----------
func findMaxDistance(
	parent map[*BinaryTreeNode]*BinaryTreeNode,
	target *BinaryTreeNode,
) int {

	queue := []*BinaryTreeNode{target}
	visited := make(map[*BinaryTreeNode]bool)
	visited[target] = true

	maxTime := 0

	for len(queue) > 0 {
		size := len(queue)
		burned := false

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil && !visited[node.Left] {
				burned = true
				visited[node.Left] = true
				queue = append(queue, node.Left)
			}

			if node.Right != nil && !visited[node.Right] {
				burned = true
				visited[node.Right] = true
				queue = append(queue, node.Right)
			}

			if p, ok := parent[node]; ok && !visited[p] {
				burned = true
				visited[p] = true
				queue = append(queue, p)
			}
		}

		if burned {
			maxTime++
		}
	}

	return maxTime
}

// ---------- Build Parent Map + Find Target ----------
func bfsToMapParents(
	root *BinaryTreeNode,
	parent map[*BinaryTreeNode]*BinaryTreeNode,
	start int,
) *BinaryTreeNode {

	queue := []*BinaryTreeNode{root}
	var target *BinaryTreeNode

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Data == start {
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

// ---------- Main Function ----------
func timeToBurnTree(root *BinaryTreeNode, start int) int {
	parent := make(map[*BinaryTreeNode]*BinaryTreeNode)

	target := bfsToMapParents(root, parent, start)

	return findMaxDistance(parent, target)
}