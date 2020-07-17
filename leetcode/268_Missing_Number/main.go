package main

func missingNumber(nums []int) int {
	numbers := make(map[int]bool, len(nums))

	for _, number := range nums {
		numbers[number] = true
	}

	for i := 0; i <= len(nums); i++ {
		if !numbers[i] {
			return i
		}
	}

	return 0
}
