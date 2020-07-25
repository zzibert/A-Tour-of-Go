package main

func intersection(nums1 []int, nums2 []int) []int {
	numbers := make(map[int]bool)

	for _, val := range nums1 {
		numbers[val] = true
	}

	intersection := make([]int, 0)

	for _, val := range nums2 {
		if numbers[val] {
			intersection = append(intersection, val)
			numbers[val] = false
		}
	}

	return intersection
}
