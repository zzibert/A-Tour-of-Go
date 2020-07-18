package main

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	digits1 := Atoi(a)
	digits2 := Atoi(b)

	number := Add(digits1, digits2)

	return Itoa(number)

}

func Add(a []int, b []int) []int {
	length1 := len(a)
	length2 := len(b)
	var c []int
	if length2 > length1 {
		c = Calculate(b, a)
	} else {
		c = Calculate(a, b)
	}

	for i := len(c) - 1; i > 0; i-- {
		if c[i] >= 2 {
			c[i] -= 2
			c[i-1] += 1
		}
	}
	if c[0] >= 2 {
		c = append([]int{1, c[0] - 2}, c[1:]...)
	}

	return c
}

func Calculate(a, b []int) []int {
	length1 := len(a)
	length2 := len(b)

	for length2 > 0 {
		a[length1-1] = a[length1-1] + b[length2-1]
		length1--
		length2--
	}
	return a
}

func Atoi(a string) []int {
	numbers := make([]int, 0)
	digits := strings.Split(a, "")
	for _, digit := range digits {

		number, _ := strconv.Atoi(digit)
		numbers = append(numbers, number)
	}

	return numbers
}

func Itoa(digits []int) string {
	letters := make([]string, 0)
	for _, digit := range digits {
		letter := strconv.Itoa(digit)
		letters = append(letters, letter)
	}

	return strings.Join(letters, "")
}
