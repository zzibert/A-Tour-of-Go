package main

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type LinkedList struct {
	root   *node
	Length int
}

func main() {
	linkedList := &LinkedList{&node{value: 1, next: nil}, 1}
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	linkedList.Prepend(0)
	linkedList.Prepend(-1)
	linkedList.Prepend(-2)
	linkedList.Prepend(-3)

	printLinkedList(linkedList.root)
}

func (list *LinkedList) Append(value interface{}) {

	current := list.root

	for current.next != nil {
		current = current.next
	}
	newNode := &node{
		value: value,
		next:  nil,
	}
	current.next = newNode
}

func (list *LinkedList) Prepend(value interface{}) {
	newNode := &node{value, list.root}
	list.root = newNode
}

func printLinkedList(current *node) {
	fmt.Println(current.value)
	if current.next != nil {
		printLinkedList(current.next)
	}
}
