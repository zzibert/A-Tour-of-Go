package main

import (
	"strconv"
	"strings"
)

func findNumbers(nums []int) int {
	counter := 0

	for _, number := range nums {
		if splitInt(number)%2 == 0 {
			counter++
		}
	}
	return counter
}

func splitInt(number int) (digits int) {
	str := strconv.Itoa(number)
	digits = len(strings.Split(str, ""))

	return
}
