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
	linkedList.append(2)
	linkedList.append(3)
	linkedList.append(4)
	linkedList.append(5)

	printLinkedList(linkedList.root)
}

func (list *LinkedList) append(value interface{}) {

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

func printLinkedList(current *node) {
	fmt.Println(current.value)
	if current.next != nil {
		printLinkedList(current.next)
	}
}
