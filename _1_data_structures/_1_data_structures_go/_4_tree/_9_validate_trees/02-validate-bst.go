// https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	return validate(root, &prev)
}

func validate(node *TreeNode, prev **TreeNode) bool {
	if node == nil {
		return true
	}

	if !validate(node.Left, prev) {
		return false
	}

	if *prev != nil && (*prev).Val >= node.Val {
		return false
	}

	*prev = node

	return validate(node.Right, prev)
}