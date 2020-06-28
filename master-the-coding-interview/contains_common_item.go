package main

import "fmt"

func main() {
	fmt.Println(ContainsCommonItem([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9}))
	fmt.Println(ContainsCommonItem([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 1}))
}

func ContainsCommonItem(nums1, nums2 []int) bool {

	map1 := make(map[int]bool)

	for _, val := range nums1 {
		map1[val] = true
	}

	for _, val := range nums2 {
		if _, ok := map1[val]; ok {
			return true
		}
	}

	return false
}
