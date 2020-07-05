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
	top    int
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.top == 0
}

func (stack *ArrayStack) Push(value interface{}) {
	stack.values = append(stack.values, value)
	stack.top++
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		value := stack.values[stack.top-1]
		stack.values = stack.values[:(stack.top - 1)]
		stack.top--
		return value
	}
}

func (stack *ArrayStack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		return stack.values[(stack.top - 1)]
	}
}
