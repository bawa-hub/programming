// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/


package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/**************** LCA ****************/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	curr := root.val

	// both nodes in right subtree
	if curr < p.val && curr < q.val {
		return lowestCommonAncestor(root.right, p, q)
	}

	// both nodes in left subtree
	if curr > p.val && curr > q.val {
		return lowestCommonAncestor(root.left, p, q)
	}

	// split point
	return root
}

func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {

	for root != nil {

		if root.val < p.val && root.val < q.val {
			root = root.right
		} else if root.val > p.val && root.val > q.val {
			root = root.left
		} else {
			return root
		}
	}

	return nil
}

/**************** TEST ****************/

func main() {

	root := &TreeNode{val: 6}
	root.left = &TreeNode{val: 2}
	root.right = &TreeNode{val: 8}
	root.left.left = &TreeNode{val: 0}
	root.left.right = &TreeNode{val: 4}
	root.left.right.left = &TreeNode{val: 3}
	root.left.right.right = &TreeNode{val: 5}
	root.right.left = &TreeNode{val: 7}
	root.right.right = &TreeNode{val: 9}

	p := root.left       // 2
	q := root.left.right // 4

	lca := lowestCommonAncestor(root, p, q)

	fmt.Println("LCA:", lca.val)
}