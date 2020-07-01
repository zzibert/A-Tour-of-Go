package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	if node.Next.Next != nil {
		deleteNode(node.Next)
	} else {
		node.Next = nil
	}
}
