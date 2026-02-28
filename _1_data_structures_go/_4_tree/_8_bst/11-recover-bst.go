// https://leetcode.com/problems/recover-binary-search-tree/


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	first  *TreeNode
	middle *TreeNode
	last   *TreeNode
	prev   *TreeNode
}

func (s *Solution) recoverTree(root *TreeNode) {
	s.first = nil
	s.middle = nil
	s.last = nil
	s.prev = &TreeNode{Val: -1 << 63} // equivalent to INT_MIN

	s.inorder(root)

	if s.first != nil && s.last != nil {
		s.first.Val, s.last.Val = s.last.Val, s.first.Val
	} else if s.first != nil && s.middle != nil {
		s.first.Val, s.middle.Val = s.middle.Val, s.first.Val
	}
}

func (s *Solution) inorder(root *TreeNode) {
	if root == nil {
		return
	}

	s.inorder(root.Left)

	if s.prev != nil && root.Val < s.prev.Val {

		// first violation
		if s.first == nil {
			s.first = s.prev
			s.middle = root
		} else {
			// second violation
			s.last = root
		}
	}

	s.prev = root

	s.inorder(root.Right)
}