package main

func removeDuplicates(nums []int) int {
	length := len(nums)

	for i := 0; i < length; i++ {
		number := countSame(nums, nums[i], i)
		moveElements(nums, i, number)
		length = (length - (number - 1))
		nums = nums[:length]
	}

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

func moveElements(nums []int, index int, number int) {
	for i := index + number; i < len(nums); i++ {
		nums[i-(number-1)] = nums[i]
	}
}
