// https://leetcode.com/problems/binary-search-tree-iterator/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushAll(root)
	return it
}

// Returns whether next smallest exists
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// Returns next smallest element
func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	this.pushAll(node.Right)

	return node.Val
}

func (this *BSTIterator) pushAll(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}