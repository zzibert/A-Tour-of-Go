package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ITERATIVE VERSION SOLVED
func reverseList(head *ListNode) *ListNode {
	var last *ListNode

	if head == nil {
		return nil
	}

	for head.Next != nil {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}

	head.Next = last

	return head
}
