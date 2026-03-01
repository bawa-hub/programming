// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/solutions/106059/java-c-three-simple-methods-choose-one-you-like/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack   []*TreeNode
	reverse bool
}

func NewBSTIterator(root *TreeNode, isReverse bool) *BSTIterator {
	it := &BSTIterator{
		reverse: isReverse,
	}
	it.pushAll(root)
	return it
}

func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

func (it *BSTIterator) Next() int {
	node := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]

	if !it.reverse {
		it.pushAll(node.Right)
	} else {
		it.pushAll(node.Left)
	}

	return node.Val
}

func (it *BSTIterator) pushAll(node *TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)

		if it.reverse {
			node = node.Right
		} else {
			node = node.Left
		}
	}
}

// ---------------- Solution ----------------

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	l := NewBSTIterator(root, false) // inorder
	r := NewBSTIterator(root, true)  // reverse inorder

	i := l.Next()
	j := r.Next()

	for i < j {
		sum := i + j

		if sum == k {
			return true
		} else if sum < k {
			i = l.Next()
		} else {
			j = r.Next()
		}
	}

	return false
}