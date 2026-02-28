// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

// recursive
func preorder(root *Node) []int {
	result := []int{}
	travel(root, &result)
	return result
}

func travel(root *Node, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)

	for _, child := range root.Children {
		travel(child, result)
	}
}

// iterative
func preorderIterative(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*Node{root}

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, n.Val)

		// push children in reverse order
		for i := len(n.Children) - 1; i >= 0; i-- {
			stack = append(stack, n.Children[i])
		}
	}

	return result
}