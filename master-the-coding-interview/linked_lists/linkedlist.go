package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}
type LinkedList struct {
	root   *Node
	Length int
}

func main() {
	linkedList := &LinkedList{&Node{value: 1, next: nil}, 1}
	linkedList.Append(1)
	linkedList.Append(1)
	linkedList.Insert(0, 1)
	linkedList.Insert(2, 5)
	linkedList.Insert(2, 6)
	linkedList.Insert(2, 7)
	linkedList.Insert(234234234, 5)

	printLinkedList(linkedList.root)
}

func (list *LinkedList) Append(value interface{}) {

	current := list.root

	for current.next != nil {
		current = current.next
	}
	newNode := &Node{
		value: value,
		next:  nil,
	}
	current.next = newNode
	list.Length++
}

func (list *LinkedList) Prepend(value interface{}) {
	newNode := &Node{value, list.root}
	list.root = newNode
	list.Length++
}

func (list *LinkedList) Insert(index int, value interface{}) {
	if index == 0 {
		newNode := &Node{value: value, next: list.root}
		list.root = newNode
		list.Length++
	} else {
		current := list.root
		for ; (index - 1) > 0; index-- {
			if current.next == nil {
				current.next = &Node{value: value, next: nil}
				list.Length++
				return
			} else {
				current = current.next
			}
		}
		newNode := &Node{value: value, next: current.next}
		current.next = newNode
		list.Length++
	}
}

func printLinkedList(current *Node) {
	fmt.Println(current.value)
	if current.next != nil {
		printLinkedList(current.next)
	}
}
