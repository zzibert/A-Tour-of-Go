package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	even_head, even_current := head.Next, head.Next
	odd_head, odd_current := head, head
	// ODD LINKED LIST

	for {
		odd_node := odd_current.Next.Next
		odd_current.Next = odd_node
		if odd_current.Next == nil {
			break
		}
		odd_current = odd_current.Next
	}

	// EVEN LINKED lIST
	for {
		if even_current.Next == nil || even_current == nil {
			break
		}
		even_node := even_current.Next.Next
		even_current.Next = even_node
		even_current = even_current.Next
	}

	// COMBINE THE TWO LINKED LISTS
	odd_current.Next = even_head

	return odd_head
}
