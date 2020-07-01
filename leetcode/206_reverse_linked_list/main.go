package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ITERATED VERSION, NOT GOOD USED ARRAYS
func reverseList(head *ListNode) *ListNode {
	var counter int
	nodes := make([]*ListNode, 0)
	if head == nil {
		return nil
	}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
		counter += 1
	}
	for i := counter - 1; i >= 0; i-- {
		if i == 0 {
			nodes[i].Next = nil
		} else {
			nodes[i].Next = nodes[i-1]
		}
	}
	return nodes[counter-1]
}
