package main

type MyQueue struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	size := len(this.stack)
	tempStack := make([]int, 0)

	for size > 0 {
		pop := this.Pop()
		tempStack = append(tempStack, pop)
		size--
	}
	this.stack = append(this.stack, x)
	for size = len(tempStack); size > 0; {
		pop := tempStack[size-1]
		tempStack = tempStack[:size-1]
		this.stack = append(this.stack, pop)
		size--
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	length := len(this.stack)
	pop := this.stack[length-1]
	this.stack = this.stack[:length-1]

	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {

	length := len(this.stack)

	return this.stack[length-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
