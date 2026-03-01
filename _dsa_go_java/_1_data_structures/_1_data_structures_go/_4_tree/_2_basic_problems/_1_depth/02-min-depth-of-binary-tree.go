// https://leetcode.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    // leaf node
    if root.Left == nil && root.Right == nil {
        return 1
    } else if root.Left != nil && root.Right == nil {
        return minDepth(root.Left) + 1
    } else if root.Left == nil && root.Right != nil {
        return minDepth(root.Right) + 1
    } else {
        return 1 + min(minDepth(root.Left), minDepth(root.Right))
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}