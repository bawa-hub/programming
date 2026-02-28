// https://leetcode.com/problems/add-two-numbers/

// in case of multiple list, dummy node is very useful as it simplifies the code
// note that temp->next in while loop is changing the dummy->next as initially temp and dummy is indeed same.


package main

import "fmt"

/************ LIST NODE ************/

type ListNode struct {
	Val  int
	Next *ListNode
}

/************ ADD TWO NUMBERS ************/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	temp := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {

		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		temp.Next = &ListNode{Val: sum % 10}
		temp = temp.Next
	}

	return dummy.Next
}
// TC: O(max(m,n))
// SC: O(max(m,n))

/************ HELPER ************/

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

/************ MAIN ************/

func main() {

	l1 := &ListNode{2,
		&ListNode{4,
			&ListNode{3, nil},
		},
	}

	l2 := &ListNode{5,
		&ListNode{6,
			&ListNode{4, nil},
		},
	}

	result := addTwoNumbers(l1, l2)
	printList(result)
}