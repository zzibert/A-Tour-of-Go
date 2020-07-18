package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)

	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}
	nums1 = nums1[:m]

	for i := 0; i < m; i++ {
		if nums2[0] < nums1[i] {
			temp := nums1[i]
			nums1[i] = nums2[0]
			nums2[0] = temp
			sort(nums2)
		}
	}

	for _, number := range nums2 {
		nums1 = append(nums1, number)
	}
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
		nums[i] = min
		nums[index] = temp
	}
}
