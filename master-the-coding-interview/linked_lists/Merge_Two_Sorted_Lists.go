package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var current, head *ListNode

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		head, current = l1, l1
		l1 = l1.Next
	} else {
		head, current = l2, l2
		l2 = l2.Next
	}

	for {
		if l1 == nil {
			current.Next = l2
			break
		}
		if l2 == nil {
			current.Next = l1
			break
		}
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	return head
}
