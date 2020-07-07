package main

import "fmt"

func main() {
	nums := []int{6, 5, 8, 2, 1, 7, 3, 4, 9}

	fmt.Println(bubbleSort(nums))
}

func bubbleSort(numbers []int) []int {
	for {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
		if isSorted(numbers) {
			return numbers
		}
	}
}

func isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
