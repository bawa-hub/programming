// https://leetcode.com/problems/diameter-of-binary-tree/

/**
 * Diameter of a binary tree:
 * longest path b/w two nodes
 * path doesn't need to pass through root node
 *
 */

 type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    height(root, &diameter)
    return diameter
}

func height(node *TreeNode, diameter *int) int {
    if node == nil {
        return 0
    }

    lh := height(node.Left, diameter)
    rh := height(node.Right, diameter)

    if lh+rh > *diameter {
        *diameter = lh + rh
    }

    if lh > rh {
        return lh + 1
    }
    return rh + 1
}

// Time Complexity: O(N) 
// Space Complexity: O(1) Extra Space + O(H) Recursion Stack space (Where “H”  is the height of binary tree )  