package main

func sortedSquares(A []int) []int {

	for i, number := range A {
		A[i] = number * number
	}

	sort(A)

	return A
}

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				index = j
			}
		}
		temp := nums[i]
		nums[i] = min
		nums[index] = temp
	}
}
