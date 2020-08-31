package main

func arrayPairSum(nums []int) int {

	nums = sort(nums)

	sum := 0

	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum

}

func sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				index = j
			}
		}
		temp := nums[i]
		nums[i] = min
		nums[index] = temp
	}

	return nums
}
