// https://leetcode.com/problems/delete-node-in-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		return helper(root)
	}

	dummy := root

	for root != nil {
		if key < root.Val {
			if root.Left != nil && root.Left.Val == key {
				root.Left = helper(root.Left)
				break
			}
			root = root.Left
		} else {
			if root.Right != nil && root.Right.Val == key {
				root.Right = helper(root.Right)
				break
			}
			root = root.Right
		}
	}

	return dummy
}

func helper(root *TreeNode) *TreeNode {
	// case 1: no left child
	if root.Left == nil {
		return root.Right
	}

	// case 2: no right child
	if root.Right == nil {
		return root.Left
	}

	rightChild := root.Right
	lastRight := findLastRight(root.Left)

	lastRight.Right = rightChild

	return root.Left
}

func findLastRight(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}
// Time Complexity: O(H) where H is the height of the tree. This is due to the binary search applied while finding the node with value as key. All other operations performed are in constant time. O(H) ~ O(log N) in case of a full binary search tree (optimal time).
// Space Complexity: O(1) as no additional data structures or memory allocation is done