package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	first := head
	var last *ListNode

	for first != last {
		current := first
		for current.Next != last {
			current = current.Next
		}
		if first.Val != current.Val {
			return false
		}
		last = current
		first = first.Next
		if first.Next == last {
			return true
		}
	}
	return true
}
