package main

func majorityElement(nums []int) int {
	half := len(nums) / 2

	numbers := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		numbers[nums[i]]++

		if numbers[nums[i]] > half {
			return nums[i]
		}
	}
	return 0
}
