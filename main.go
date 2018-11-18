package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum1 := 1
	sum2 := 1
	return func() int {
		sum3 := sum2
		sum2 = sum2 + sum1
		sum1 = sum3
		return sum2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
