package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next.Next

	for fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next

		fast = fast.Next
	}
	return false
}
