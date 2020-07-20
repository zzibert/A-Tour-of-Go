package main

func moveZeroes(nums []int) {
	var pointer1, zeroCounter int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCounter++
		} else {
			nums[pointer1] = nums[i]
			pointer1++
		}
	}

	for i := len(nums) - zeroCounter; i < len(nums); i++ {
		nums[i] = 0
	}
}
