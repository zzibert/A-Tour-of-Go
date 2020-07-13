package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
  nodes := make([]*ListNode, 0)
  length := 0

  for head != nil {
    length++
    nodes = append(nodes, head)
    head = head.Next
  }

  return nodes[(length/2)]
}
