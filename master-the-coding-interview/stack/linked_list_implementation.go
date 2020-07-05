package main

import "fmt"

func main() {
	stack := &Stack{top: nil, bottom: nil, length: 0}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

type Node struct {
	val  interface{}
	next *Node
}

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

func (stack *Stack) Push(value interface{}) {
	newNode := &Node{val: value, next: nil}
	if stack.IsEmpty() {
		stack.top = newNode
		stack.bottom = newNode
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

func (stack *Stack) Peek() *Node {
	if stack.IsEmpty() {
		return nil
	} else {
		return stack.top
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.length == 0
}
