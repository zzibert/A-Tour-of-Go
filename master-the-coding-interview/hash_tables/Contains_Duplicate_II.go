package main

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	elements := make(map[int]int)

	if len(nums) == 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := elements[nums[i]]; ok {
			if math.Abs(float64(elements[nums[i]])-float64(i)) <= float64(k) {
				return true
			} else {
				delete(elements, nums[i])
			}
		}
		elements[nums[i]] = i
	}

	return false
}
