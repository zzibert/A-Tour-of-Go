// package main

// type MyCircularQueue struct {
// 	queue []int
// 	size  int
// 	first int
// 	last  int
// }

// /** Initialize your data structure here. Set the size of the queue to be k. */
// func Constructor(k int) MyCircularQueue {
// 	queue := make([]int, k)
// 	return MyCircularQueue{queue: queue, size: 0, first: 0, last: -1}
// }

// /** Insert an element into the circular queue. Return true if the operation is successful. */
// func (this *MyCircularQueue) EnQueue(value int) bool {
// 	if !this.IsFull() {
// 		this.last++
// 		this.last = this.last % len(this.queue)
// 		this.queue[this.last] = value
// 		this.size++
// 		return true
// 	}

// 	return false
// }

// /** Delete an element from the circular queue. Return true if the operation is successful. */
// func (this *MyCircularQueue) DeQueue() bool {
// 	if this.IsEmpty() {
// 		return false
// 	}
// 	array := append(this.queue[:this.first], 0)
// 	this.queue = append(array, this.queue[(this.first+1):]...)
// 	this.first++
// 	this.first = this.first % len(this.queue)
// 	this.size--
// 	return true
// }

// /** Get the front item from the queue. */
// func (this *MyCircularQueue) Front() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
// 	return this.queue[this.first]
// }

// /** Get the last item from the queue. */
// func (this *MyCircularQueue) Rear() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}

// 	return this.queue[this.last]
// }

// /** Checks whether the circular queue is empty or not. */
// func (this *MyCircularQueue) IsEmpty() bool {
// 	return this.size == 0
// }

// /** Checks whether the circular queue is full or not. */
// func (this *MyCircularQueue) IsFull() bool {
// 	return len(this.queue) <= this.size
// }
