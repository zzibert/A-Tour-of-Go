package main

func findMaxConsecutiveOnes(nums []int) int {

	max := 1

	zeroes := findZeroes(nums)

	if len(zeroes) == 0 {
		return len(nums)
	}

	for _, zero := range zeroes {
		localMax := countOnes(nums, zero)
		if localMax > max {
			max = localMax
		}
	}

	return max
}

func findZeroes(nums []int) []int {
	array := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			array = append(array, i)
		}
	}
	return array
}

func countOnes(nums []int, index int) int {
	count := 1
	left := index - 1
	right := index + 1

	// count ones before flipped zero
	for left >= 0 && nums[left] == 1 {
		count++
		left--
	}

	// count ones after flipped zero
	for right < len(nums) && nums[right] == 1 {
		count++
		right++
	}

	return count
}
