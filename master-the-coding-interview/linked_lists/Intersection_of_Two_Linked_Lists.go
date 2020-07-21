package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	nodes := make(map[*ListNode]bool)

	//Traverse first linked list.
	for headA != nil {
		nodes[headA] = true
		headA = headA.Next
	}

	// Traverse second linked list , check if node is in map.
	for headB != nil {
		if nodes[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
