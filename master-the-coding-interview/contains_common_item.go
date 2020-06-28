package main

import "fmt"

func main() {
	fmt.Println(ContainsCommonItem([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9}))
	fmt.Println(ContainsCommonItem([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 1}))
}

func ContainsCommonItem(nums1, nums2 []int) bool {

	for _, num := range nums1 {
		for _, num2 := range nums2 {
			if num == num2 {
				return true
			}
		}
	}
	return false
}
