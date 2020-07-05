package main

import "fmt"

func main() {
	stack := &ArrayStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)

	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Pop()

}

type ArrayStack struct {
	values []interface{}
}

func (stack *ArrayStack) IsEmpty() bool {
	return len(stack.values) == 0
}

func (stack *ArrayStack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		length := len(stack.values)
		value, values := stack.values[length-1], stack.values[:(length-1)]
		stack.values = values
		return value
	}
}

func (stack *ArrayStack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		length := len(stack.values)
		return stack.values[(length - 1)]
	}
}
