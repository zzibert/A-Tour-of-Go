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
	current_node := head
	even_head := head.Next

	for current_node.Next != nil {
		next_node := current_node.Next
		current_node.Next = current_node.Next.Next
		current_node = next_node
	}
	current_node = head
	for current_node.Next != nil {
		current_node = current_node.Next
	}

	current_node.Next = even_head

	return head
}
