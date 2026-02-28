// https://leetcode.com/problems/insert-into-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	cur := root

	for {
		if val >= cur.Val {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur.Right = &TreeNode{Val: val}
				break
			}
		} else {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur.Left = &TreeNode{Val: val}
				break
			}
		}
	}

	return root
}
// Time Complexity: O(log N) because of the logarithmic height of the Binary Search Tree that is traversed during the insertion process.
// Space Complexity: O(1) as no additional data structures or memory allocation is done