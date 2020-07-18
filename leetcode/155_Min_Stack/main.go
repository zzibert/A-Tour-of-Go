package main

type MinStack struct {
	stack  []int
	min    int
	length int
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: MinInt}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.min = x
	} else if x < this.min {
		this.min = x
	}

	this.stack = append(this.stack, x)
	this.length++
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}

	this.stack = this.stack[:this.length-1]
	this.min = findMin(this.stack)
	this.length--
}

func (this *MinStack) Top() int {
	if this.length == 0 {
		return 0
	}
	return this.stack[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func findMin(stack []int) int {
	min := MaxInt
	for _, number := range stack {
		if number < min {
			min = number
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
