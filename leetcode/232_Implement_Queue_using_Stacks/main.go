package main

type MyQueue struct {
	stack []int
	size  int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
	this.size++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	pop := this.stack[0]
	this.stack = this.stack[1:]
	this.size--

	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
