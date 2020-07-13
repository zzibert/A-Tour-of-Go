package main

import "strconv"

func fizzBuzz(n int) []string {
	array := make([]string, 0)

	for i := 1; i <= n; i++ {
		array = append(array, strconv.Itoa(i))
	}

	for i := 3; i <= n; i += 3 {
		array[i-1] = "Fizz"
	}

	for i := 5; i <= n; i += 5 {
		array[i-1] = "Buzz"
	}

	for i := 15; i <= n; i += 15 {
		array[i-1] = "FizzBuzz"
	}

	return array
}
