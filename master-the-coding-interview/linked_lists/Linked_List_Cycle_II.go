package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nodes := make(map[*ListNode]bool)

	current := head

	for current != nil {
		if nodes[current] {
			return current
		}
		nodes[current] = true
		current = current.Next
	}

	return nil
}
