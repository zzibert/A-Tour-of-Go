package main

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func heightChecker(heights []int) int {
	array := make([]int, 0)
	counter := 0

	for _, number := range heights {
		array = append(array, number)
	}

	sort(array)

	for i := 0; i < len(heights); i++ {
		if heights[i] != array[i] {
			counter++
		}
	}
	return counter
}

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		index := i
		for j := i + 1; j < len(nums); j++ {
			if min > nums[j] {
				min = nums[j]
				index = j
			}
		}
		temp := nums[i]
		nums[i] = nums[index]
		nums[index] = temp
	}
}
