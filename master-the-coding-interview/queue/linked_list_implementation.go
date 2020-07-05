package main

import "fmt"

func main() {
	q := Queue{first: nil, last: nil, length: 0}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println(q.Peek())
	q.Dequeue()
}

type Node struct {
	val  interface{}
	next *Node
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Peek() *Node {
	if q.IsEmpty() {
		return nil
	} else {
		return q.last
	}
}

func (q *Queue) Enqueue(value interface{}) {
	node := &Node{val: value, next: nil}
	if q.IsEmpty() {
		q.first = node
		q.last = node

	} else {
		node.next = q.first
		q.first = node
	}
	q.length++
}

func (q *Queue) Dequeue() *Node {
	if q.IsEmpty() {
		return nil
	} else {
		current := q.first
		last := q.last
		for current.next != q.last && current.next != nil {
			current = current.next
		}
		current.next = nil
		q.last = current
		q.length--
		return last
	}
}
