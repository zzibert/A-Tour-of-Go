package main

func minSubArrayLen(s int, nums []int) int {
	length := len(nums)

	min := length

	if numsSum(nums) < s {
		return 0
	}

	for i := 0; i < length; i++ {
		localSum := 0
		localMin := 0
		for j := i; localSum < s && j < length && localMin < min; j++ {
			localMin++
			localSum += nums[j]
		}
		if localSum >= s && localMin < min {
			min = localMin
		}
	}

	return min
}

func numsSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}
