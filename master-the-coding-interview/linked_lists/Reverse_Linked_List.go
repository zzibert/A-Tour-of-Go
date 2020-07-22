package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	current := head.Next
	head.Next = nil

	if current.Next == nil {
		current.Next = head
		return current
	}

	temp := current.Next
	for temp != nil {
		current.Next = head
		head = current
		current = temp
		temp = temp.Next
	}
	current.Next = head
	return current
}
