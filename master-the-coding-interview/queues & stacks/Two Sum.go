package main

func twoSum(nums []int, target int) []int {
	diff := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff[nums[i]] = i
	}

	for i, number := range nums {
		difference := target - number
		if index, ok := diff[difference]; ok && i != index {
			return []int{i, index}
		}
	}
	return []int{}
}
