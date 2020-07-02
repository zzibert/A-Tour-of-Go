package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// ITERATIVE VERSION
func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		testNode := current
		for testNode != nil && current.Val == testNode.Val {
			testNode = testNode.Next
		}
		current.Next = testNode
		current = current.Next
	}

	return head
}
