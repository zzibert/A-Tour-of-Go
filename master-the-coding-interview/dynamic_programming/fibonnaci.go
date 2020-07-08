package main

import "fmt"

func main() {
	fmt.Println(fibonnaci(4))
	fmt.Println(fibonnaci(5))
	fmt.Println(fibonnaci(6))
	fmt.Println(fibonnaci(7))
	fmt.Println(fibonnaci(8))
	fmt.Println(fibonnaci(9))
	fmt.Println(fibonnaci(10))
	fmt.Println(fibonnaci(11))
	fmt.Println(fibonnaci(50))
}

func fibonnaci(n int) int {
	fibs := make(map[int]int)

	fibs[0] = 0
	fibs[1] = 1
	fibs[2] = 1

	for i := 3; i < n; i++ {
		value := fibs[i-1] + fibs[i-2]
		fibs[i] = value
	}

	return fibs[n-1] + fibs[n-2]
}
