package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	parentNode := &ListNode{Next: head}

	if n == 0 {
		return head
	}

	if n == 1 {
		if head.Next == nil {
			return nil
		}
		current := head.Next
		for current != nil {
			parentNode = parentNode.Next
			current = current.Next
		}
		parentNode.Next = nil
		return head
	}
	current := head
	originalHead := head

	for i := 0; i < n; i++ {
		current = current.Next
	}

	for current != nil {
		parentNode = parentNode.Next
		head = head.Next
		current = current.Next
	}
	if head == originalHead {
		return head.Next
	}

	parentNode.Next = head.Next
	return originalHead
}
