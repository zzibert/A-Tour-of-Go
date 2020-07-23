package main

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		node := &Node{Val: x}
		node.Next = node
		return node
	}

	start := aNode
	prev := start.Next
	current := prev.Next

	if aNode == aNode.Next {
		new := &Node{Val: x, Next: current}
		prev.Next = new
		return current
	}

	for prev != start {
		if prev.Val <= x && x <= current.Val {
			new := &Node{Val: x, Next: current}
			prev.Next = new
			return current
		}
		prev = current
		current = current.Next
	}

	start = current
	prev = start.Next
	current = prev.Next

	for start.Next != current {
		if prev.Val <= x {
			new := &Node{Val: x, Next: current}
			prev.Next = new
			return new
		}

		prev = current
		current = current.Next
	}

	new := &Node{Val: x, Next: current}
	prev.Next = new

	return current
}
