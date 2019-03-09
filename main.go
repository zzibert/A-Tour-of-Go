package main

import "fmt"

func foo() *int {
	x := 1
	return &x
}

func main() {
	var y *int
	y = foo()
	fmt.Printf("Hi %s\n", "Joe")
	fmt.Printf("%d", *y)
}
