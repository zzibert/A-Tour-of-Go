package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib1, fib2 := 1, 1
	return func() int {
		temp := fib2
		fib2 = fib2 + fib1
		fib1 = temp
		return fib2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}