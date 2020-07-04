package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	} else {
		values := make(map[&ListNode]bool)
		for headA != nil {
			values[headA] = true
			headA = headA.Next
		}

		for headB != nil {
			if values[headB.Val] {
				return headB
			} else {
				headB = headB.Next
			}
		}

		return nil
	}
}
