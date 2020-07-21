package main

func findDisappearedNumbers(nums []int) []int {
	numbers := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		numbers[i+1] = false
	}

	for _, number := range nums {
		numbers[number] = true
	}
	array := make([]int, 0)
	for number, value := range numbers {
		if !value {
			array = append(array, number)
		}
	}
	return array
}
