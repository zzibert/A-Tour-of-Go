package main

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func thirdMax(nums []int) int {
	maximums := make([]int, 3)

	for i := 0; i < 3; i++ {
		max := findMax(nums)
		maximums[i] = max
		removeDuplicates(nums, max)
	}

	if maximums[2] == MinInt {
		return maximums[0]
	}

	return maximums[2]
}

func findMax(nums []int) int {
	max := MinInt

	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func removeDuplicates(nums []int, val int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = MinInt
		}
	}
}
