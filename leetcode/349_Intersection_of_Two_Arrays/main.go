package main

func intersection(nums1 []int, nums2 []int) []int {
	presentElements := make(map[int]bool)

	intersection := make([]int, 0)

	for _, element := range nums1 {
		presentElements[element] = true
	}

	for _, element := range nums2 {
		if presentElements[element] {
			intersection = append(intersection, element)
			presentElements[element] = false
		}
	}

	return intersection
}
