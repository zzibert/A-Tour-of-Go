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
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return false
}
