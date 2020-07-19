package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	index := 0

	for index < len(nums) {
		if nums[index] == 1 {
			localMax := 0
			localIndex := index
			for localIndex < len(nums) && nums[localIndex] == 1 {
				localMax++
				localIndex++
			}
			if localMax > max {
				max = localMax
			}
			index = localIndex
		} else {
			index++
		}
	}
	return max
}
