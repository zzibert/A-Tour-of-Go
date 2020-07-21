package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ptr := head
	fast := head
	slow := head
	var intersection *ListNode

	for fast != nil || fast.Next != nil {
		if slow == fast {
			intersection = slow
			if intersection.Next == head {
				return intersection
			}
			for {
				ptr = ptr.Next
				intersection = intersection.Next
				if ptr == intersection {
					return ptr
				}
			}
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return nil
}
