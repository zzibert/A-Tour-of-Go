package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	oddHead, odd := head, head

	if oddHead == nil {
		return nil
	}

	evenHead, even := head.Next, head.Next
	var lastOdd *ListNode

	if evenHead == nil {
		return oddHead
	}

	for odd.Next != nil && even.Next != nil {
		lastOdd = odd
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		even = even.Next
		odd = odd.Next
	}

	if odd == nil {
		lastOdd.Next = evenHead
	} else {
		odd.Next = evenHead
	}

	return oddHead
}
