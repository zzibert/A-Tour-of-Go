package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := l1

	for l1 != nil && l2 != nil {
		l1.Val = l1.Val + l2.Val
		if l1.Next == nil && l2.Next != nil {
			l1.Next = l2.Next
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	current := head

	for current != nil {
		if current.Val > 9 && current.Next == nil {
			current.Val -= 10
			node := &ListNode{Val: 1}
			current.Next = node
			break
		}
		if current.Val > 9 {
			current.Val -= 10
			current.Next.Val += 1
		}
		current = current.Next
	}

	return head
}
