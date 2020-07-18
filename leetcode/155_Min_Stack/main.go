package main

type Number struct {
	val int
	min int
}

type MinStack struct {
	stack  []Number
	length int
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	number := Number{val: x, min: x}
	if this.length > 0 && this.stack[this.length-1].min < x {
		number.min = this.stack[this.length-1].min
	}

	this.stack = append(this.stack, number)
	this.length++
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}

	this.stack = this.stack[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	if this.length == 0 {
		return 0
	}
	return this.stack[this.length-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[this.length-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
