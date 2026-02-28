// https://practice.geeksforgeeks.org/problems/left-view-of-binary-tree/1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leftSideView(root *TreeNode) []int {
	res := []int{}
	recursion(root, 0, &res)
	return res
}

func recursion(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, root.Val)
	}

	recursion(root.Left, level+1, res)
	recursion(root.Right, level+1, res)
}

// Time Complexity : O(N)
// Space Complexity : O(H)(H->Height of the Tree)