// https://leetcode.com/problems/binary-tree-right-side-view/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	recursion(root, 0, &res)
	return res
}

func recursion(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}

	// first node visited at this level
	if len(*res) == level {
		*res = append(*res, root.Val)
	}

	// visit right first
	recursion(root.Right, level+1, res)
	recursion(root.Left, level+1, res)
}
// Time Complexity : O(N)
// Space Complexity : O(H)(H->Height of the Tree)