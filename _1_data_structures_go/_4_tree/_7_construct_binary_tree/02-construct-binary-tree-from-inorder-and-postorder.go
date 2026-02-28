// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

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
	postorder []int,
	postStart, postEnd int,
	inorder []int,
	inStart, inEnd int,
	indexMap map[int]int,
) *Node {

	if postStart > postEnd || inStart > inEnd {
		return nil
	}

	root := newNode(postorder[postEnd])

	elemIndex := indexMap[root.Data]
	nElem := elemIndex - inStart

	root.Left = constructTree(
		postorder,
		postStart,
		postStart+nElem-1,
		inorder,
		inStart,
		elemIndex-1,
		indexMap,
	)

	root.Right = constructTree(
		postorder,
		postStart+nElem,
		postEnd-1,
		inorder,
		elemIndex+1,
		inEnd,
		indexMap,
	)

	return root
}

func buildTree(postorder []int, inorder []int) *Node {

	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}

	return constructTree(
		postorder,
		0,
		len(postorder)-1,
		inorder,
		0,
		len(inorder)-1,
		indexMap,
	)
}

func main() {
	postorder := []int{40, 50, 20, 60, 30, 10}
	inorder := []int{40, 20, 50, 10, 60, 30}

	root := buildTree(postorder, inorder)

	_ = root
}

// Time Complexity: O(N)
// Assumption: Hashmap returns the answer in constant time.
// Space Complexity: O(N)
// Reason: We are using an external hashmap of size ‘N’.