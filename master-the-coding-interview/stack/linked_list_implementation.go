package main

type Node struct {
	val  interface{}
	next *Node
}

type Stack struct {
	top    *Node
	length int
}

func (stack *Stack) Push(value interface{}) {
	newNode := &Node{val: value, next: nil}
	if stack.IsEmpty() {
		stack.top = newNode
	} else {
		newNode.next = stack.top
		stack.top = newNode
	}
	stack.length++
}

func (stack *Stack) Pop() *Node {
	if stack.IsEmpty() {
		return nil
	} else {
		node := stack.top
		stack.top = node.next
		stack.length--
		return node
	}
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		return stack.top.val
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.length == 0
}
