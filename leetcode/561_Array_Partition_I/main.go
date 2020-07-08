package main

import "math"

func arrayPairSum(nums []int) int {
	ordered := MergeSort(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += int(math.Min(float64(ordered[i]), float64(ordered[i+1])))
	}

	return sum
}

func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	half := len(numbers) / 2
	return SelectionSort(append(SelectionSort(numbers[:half]), SelectionSort(numbers[half:])...))
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
