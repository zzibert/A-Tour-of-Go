package main

func heightChecker(heights []int) int {
	students := make([]int, 0)

	for i := 0; i < len(heights); i++ {
		min := heights[i]
		index := i
		for j := i + 1; j < len(heights); j++ {
			if min > heights[j] {
				min = heights[j]
				index = j
			}
		}
		if i != index {
			students = includes(students, i)
			students = includes(students, index)
		}
		temp := heights[i]
		heights[i] = heights[index]
		heights[index] = temp
	}

	return len(students)
}

func isOrdered(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func max(nums []int) int {
	max := nums[0]

	for _, number := range nums {
		if number > max {
			max = number
		}
	}
	return max
}

func includes(nums []int, val int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			return nums
		}
	}
	return append(nums, val)
}
