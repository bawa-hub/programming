// https://leetcode.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    lh := maxDepth(root.Left)
    rh := maxDepth(root.Right)

    return 1 + max(lh, rh)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Time Complexity: O(N)
// Space Complexity: O(1) Extra Space + O(H) Recursion Stack space, where “H”  is the height of the binary tree