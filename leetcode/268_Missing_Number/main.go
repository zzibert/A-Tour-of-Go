package main

func missingNumber(nums []int) int {
	gauss := 0
	for i := 1; i <= len(nums); i++ {
		gauss += i
	}

	numbers := 0
	for i := 0; i < len(nums); i++ {
		numbers += nums[i]
	}

	return gauss - numbers
}
