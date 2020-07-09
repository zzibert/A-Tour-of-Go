package main

func majorityElement(nums []int) int {
	half := len(nums) / 2

	numbers := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numbers[nums[i]]++
	}

	for k, v := range numbers {
		if v > half {
			return k
		}
	}
	return 0
}
