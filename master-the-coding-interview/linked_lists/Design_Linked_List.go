package main

type MyLinkedList struct {
	head *Node
}

type Node struct {
	val  int
	next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	current := this.head
	for i := 0; i < index; i++ {
		current = current.next
		if current == nil {
			return -1
		}
	}

	return current.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := Node{val: val, next: this.head}
	this.head = &node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	current := this.head
	if current == nil {
		this.head = &Node{val: val}
		return
	}
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{val: val}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	node := Node{val: val}
	current := this.head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return
		}
		current = current.next
	}
	node.next = current.next
	current.next = &node
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.head = this.head.next
		return
	}

	current := this.head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return
		}
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}
