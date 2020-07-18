package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	numbers := make(map[int]bool)

	for n != 1 {
		Itoa := strconv.Itoa(n)
		digits := strings.Split(Itoa, "")
		number := sumOfSquares(digits)
		if numbers[number] {
			return false
		}
		numbers[number] = true
		n = number
	}

	return true
}

func sumOfSquares(numbers []string) int {
	sum := 0

	for _, str := range numbers {
		number, _ := strconv.Atoi(str)
		sum += number * number
	}

	return sum
}
