// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** LCA ****************/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// base case
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.left, p, q)
	right := lowestCommonAncestor(root.right, p, q)

	// result logic
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	// both sides found nodes
	return root
}
// Time complexity: O(N) where n is the number of nodes.
// Space complexity: O(N), auxiliary space.

/**************** TEST ****************/

func main() {

	root := &TreeNode{val: 3}
	root.left = &TreeNode{val: 5}
	root.right = &TreeNode{val: 1}
	root.left.left = &TreeNode{val: 6}
	root.left.right = &TreeNode{val: 2}
	root.right.left = &TreeNode{val: 0}
	root.right.right = &TreeNode{val: 8}

	p := root.left        // 5
	q := root.left.right  // 2

	lca := lowestCommonAncestor(root, p, q)

	fmt.Println("LCA:", lca.val)
}