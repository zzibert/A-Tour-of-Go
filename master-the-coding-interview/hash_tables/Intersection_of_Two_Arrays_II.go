package main

func intersect(nums1 []int, nums2 []int) []int {
	occurences1 := make(map[int]int)
	occurences2 := make(map[int]int)

	intersection := make([]int, 0)

	for _, number := range nums1 {
		occurences1[number]++
	}

	for _, number := range nums2 {
		occurences2[number]++
	}

	for number, count := range occurences1 {
		min := count
		if min > occurences2[number] {
			min = occurences2[number]
		}
		add(&intersection, min, number)
	}

	return intersection

}

func add(array *[]int, count, element int) {
	for i := 0; i < count; i++ {
		*array = append(*array, element)
	}
}
