// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package main

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func newNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

func constructTree(
	preorder []int,
	preStart, preEnd int,
	inorder []int,
	inStart, inEnd int,
	indexMap map[int]int,
) *Node {

	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	root := newNode(preorder[preStart])

	elemIndex := indexMap[root.Data]
	nElem := elemIndex - inStart

	root.Left = constructTree(
		preorder,
		preStart+1,
		preStart+nElem,
		inorder,
		inStart,
		elemIndex-1,
		indexMap,
	)

	root.Right = constructTree(
		preorder,
		preStart+nElem+1,
		preEnd,
		inorder,
		elemIndex+1,
		inEnd,
		indexMap,
	)

	return root
}

func buildTree(preorder []int, inorder []int) *Node {
	indexMap := make(map[int]int)

	for i, v := range inorder {
		indexMap[v] = i
	}

	return constructTree(
		preorder,
		0,
		len(preorder)-1,
		inorder,
		0,
		len(inorder)-1,
		indexMap,
	)
}

func main() {
	preorder := []int{10, 20, 40, 50, 30, 60}
	inorder := []int{40, 20, 50, 10, 60, 30}

	root := buildTree(preorder, inorder)

	_ = root
}