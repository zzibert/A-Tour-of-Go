package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	counter := 0

	for i := 0; i < length; {
		number := countSame(nums, nums[i], i)
		counter++
		nums[i+1] = nums[i+number]
		i += number
	}

	nums = nums[:counter]

	return length

}

func countSame(nums []int, val int, index int) int {
	counter := 0
	for index < len(nums) && nums[index] == val {
		counter++
		index++
	}
	return counter
}
