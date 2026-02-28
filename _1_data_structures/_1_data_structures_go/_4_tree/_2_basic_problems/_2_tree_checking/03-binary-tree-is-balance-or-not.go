// https://leetcode.com/problems/balanced-binary-tree/

// a binary tree in which the height of the left and right subtree of any node differ by not more than 1. 

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return dfsHeight(root) != -1
}

func dfsHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := dfsHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := dfsHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Time Complexity: O(N)
// Space Complexity: O(1) Extra Space + O(H) Recursion Stack space (Where “H”  is the height of binary tree)