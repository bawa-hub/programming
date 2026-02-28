// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/

type Node struct {
    Val      int
    Children []*Node
}

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    maxi := 0

    for _, child := range root.Children {
        depth := maxDepth(child)
        if depth > maxi {
            maxi = depth
        }
    }

    return 1 + maxi
}