package main

func intersection(nums1 []int, nums2 []int) []int {
	presentElements := make(map[int]bool)

	intersection := make([]int, 0)

	for _, element := range nums1 {
		presentElements[element] = true
	}

	for _, element := range nums2 {
		if !contains(intersection, element) && presentElements[element] {
			intersection = append(intersection, element)
		}
	}

	return intersection
}

func contains(array []int, element int) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}

	return false
}
