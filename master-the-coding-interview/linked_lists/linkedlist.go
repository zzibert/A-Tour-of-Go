// package main

// import "fmt"

// type Node struct {
// 	value interface{}
// 	next  *Node
// }

// type DoubleNode struct {
// 	value    interface{}
// 	next     *DoubleNode
// 	previous *DoubleNode
// }

// type LinkedList struct {
// 	root   *Node
// 	Length int
// }

// func main() {
// 	linkedList := &LinkedList{&Node{value: 1, next: nil}, 1}
// 	linkedList.Append(1)
// 	linkedList.Append(1)
// 	linkedList.Insert(0, 1)
// 	linkedList.Insert(2, 5)
// 	linkedList.Insert(2, 6)
// 	linkedList.Insert(2, 7)
// 	linkedList.Insert(234234234, 5)
// 	linkedList.Remove(2)
// 	linkedList.Remove(0)
// 	linkedList.Remove(1)

// 	fmt.Println(linkedList.Traverse())
// }

// func (list *LinkedList) Append(value interface{}) {
// 	list.Insert(list.Length, value)
// }

// func (list *LinkedList) Prepend(value interface{}) {
// 	list.Insert(0, value)
// }

// func (list *LinkedList) Insert(index int, value interface{}) {
// 	if index == 0 {
// 		newNode := &Node{value: value, next: list.root}
// 		list.root = newNode
// 		list.Length++
// 	} else {
// 		current := list.root
// 		for ; (index - 1) > 0; index-- {
// 			if current.next == nil {
// 				current.next = &Node{value: value, next: nil}
// 				list.Length++
// 				return
// 			} else {
// 				current = current.next
// 			}
// 		}
// 		newNode := &Node{value: value, next: current.next}
// 		current.next = newNode
// 		list.Length++
// 	}
// }

// func (list *LinkedList) Remove(index int) (value interface{}) {
// 	if list.Length == 0 || index > list.Length {
// 		return nil
// 	} else if index == 0 {
// 		value = list.root.value
// 		list.root = list.root.next
// 		list.Length--
// 	} else {
// 		current := list.root
// 		for ; (index - 1) > 0; index-- {
// 			if current.next == nil {
// 				return
// 			} else {
// 				current = current.next
// 			}
// 		}
// 		value = current.next.value
// 		current.next = current.next.next
// 		list.Length--
// 	}
// 	return
// }

// func (list *LinkedList) Traverse() (values []interface{}) {
// 	if list != nil {
// 		current := list.root
// 		for current != nil {
// 			values = append(values, current.value)
// 			current = current.next
// 		}
// 	}

// 	return
// }
