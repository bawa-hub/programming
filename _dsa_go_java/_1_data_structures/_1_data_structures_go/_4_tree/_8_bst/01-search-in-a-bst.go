// https://leetcode.com/problems/search-in-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
// Time Complexity: O(logN)
// Reason: The time required will be proportional to the height of the tree, if the tree is balanced, then the height of the tree is logN.
// Space Complexity: O(1)
// Reason: We are not using any extra space.