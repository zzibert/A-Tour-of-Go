package main

func findDisappearedNumbers(nums []int) []int {
	array := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		var number int
		if nums[i] < 0 {
			number = -(nums[i])
		} else {
			number = nums[i]
		}
		if nums[number-1] > 0 {
			nums[number-1] = -(nums[number-1])
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			array = append(array, i+1)
		}
	}

	return array
}
