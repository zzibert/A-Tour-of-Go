package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, current *ListNode

	// If one of the linked lists is empty, return the other one.
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	// Set the first node
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	//Set current node
	current = head

	for {
		if l1 == nil {
			current.Next = l2
			return head
		} else if l2 == nil {
			current.Next = l1
			return head
		} else {
			if l1.Val < l2.Val {
				current.Next = l1
				l1 = l1.Next
			} else {
				current.Next = l2
				l2 = l2.Next
			}
		}
		current = current.Next
	}

}
