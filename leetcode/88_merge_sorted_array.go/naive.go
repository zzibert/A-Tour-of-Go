package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var index int

	for len(nums2) != 0 && index < len(nums1) {
		if nums1[index] > nums2[0] {
			pushForward(nums1, index, nums2[0])
			nums2 = nums2[1:]
		} else {
			index++
		}
	}

	nums1 = append(nums1, nums2...)
}

func pushForward(nums []int, index int, val int) {
	temp := nums[index+1]
	for i := index; i < len(nums)-1; i++ {
		nums[i] = val
		val = temp
		temp = nums[i+1]
	}
}
