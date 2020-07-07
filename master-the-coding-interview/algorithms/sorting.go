package main

import "fmt"

func main() {
	// 6, 5, 8, 2, 1, 7, 3, 4, 9
	nums := []int{6, 5, 3, 1, 8, 7, 2, 4, 9}

	// fmt.Println(bubbleSort(nums))

	fmt.Println(InsertionSort(nums))
}

func InsertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		temp := numbers[i]
		for j := i - 1; j >= 0; j-- {
			if numbers[j] > temp {
				numbers[j+1] = numbers[j]
				if j == 0 {
					numbers[0] = temp
				}
			} else {
				numbers[j+1] = temp
				break
			}

		}
	}
	return numbers
}

func SelectionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		min := numbers[i]
		index := i
		for j := i + 1; j < len(numbers); j++ {
			if min > numbers[j] {
				min = numbers[j]
				index = j
			}
		}
		numbers[index] = numbers[i]
		numbers[i] = min
	}
	return numbers
}

func BubbleSort(numbers []int) []int {
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
