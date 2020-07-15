package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	parentNode := head
	for parentNode != nil && parentNode.Next != nil {

		childNode := parentNode.Next
		for childNode != nil && childNode.Val == val {

			childNode = childNode.Next
		}

		parentNode.Next = childNode
		parentNode = parentNode.Next
	}

	return head
}
