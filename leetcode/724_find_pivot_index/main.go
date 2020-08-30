package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	var left, right int

	for _, num := range nums {
		right += num
	}

	for i, number := range nums {
		right -= number
		if left == right {
			return i
		}
		left += number
	}

	return -1
}
