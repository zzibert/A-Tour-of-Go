package main

func findMaxConsecutiveOnes(nums []int) int {
	var max, localMax int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			localMax++

		} else {
			if localMax > max {
				max = localMax
			}
			localMax = 0
		}
	}

	if localMax > max {
		return localMax
	}

	return max
}
