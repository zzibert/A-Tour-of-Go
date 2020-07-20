package main

func removeElement(nums []int, val int) int {

	length := len(nums)

	for i := 0; i < length; {
		if nums[i] == val {
			moveElements(nums, i)
			nums = nums[:length-1]
			length--
		} else {
			i++
		}
	}

	return length
}

func moveElements(nums []int, index int) {
	for index < len(nums)-1 {
		nums[index] = nums[index+1]
		index++
	}
}
