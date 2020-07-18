package main

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	digits1 := Atoi(num1)
	digits2 := Atoi(num2)

	length1 := len(digits1)
	length2 := len(digits2)

	digits1 = append(make([]int, 5100-length1), digits1...)
	digits2 = append(make([]int, 5100-length2), digits2...)

	for i := 0; i < 5100; i++ {
		digits2[i] = digits2[i] + digits1[i]
	}

	// summing up the digits of the two numbers.
	for i := len(digits2) - 1; i >= 0; i-- {
		if digits2[i] > 9 {
			digits2[i] -= 10
			digits2[i-1] += 1
		}
	}
	return Itoa(digits2)
}

// Converting the array of digits back into number string.
func Itoa(digits []int) string {
	numbers := make([]string, 0)

	index := removeZeroes(digits)

	if index == len(digits) {
		return "0"
	}

	for _, digit := range digits[index:] {
		numbers = append(numbers, strconv.Itoa(digit))
	}
	return strings.Join(numbers, "")
}

// converting the number string into a array of digits.
func Atoi(number string) []int {
	digits := strings.Split(number, "")
	numbers := make([]int, 0)

	for _, digit := range digits {
		number, _ := strconv.Atoi(digit)
		numbers = append(numbers, number)
	}
	return numbers
}

// Remove redundant zeroes
func removeZeroes(digits []int) int {
	for i, digit := range digits {
		if digit != 0 {
			return i
		}
	}
	return len(digits)
}
