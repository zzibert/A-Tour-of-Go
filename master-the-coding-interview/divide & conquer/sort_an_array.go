package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}

	half := len(nums) / 2
	left := sortArray(nums[:half])
	right := sortArray(nums[half:])

	array := make([]int, 0)
	var l, r int
	for {
		if l == len(left) && r == len(right) {
			break
		}
		if l == len(left) {
			array = append(array, right[r])
			r++
		} else if r == len(right) {
			array = append(array, left[l])
			l++
		} else {
			if right[r] < left[l] {
				array = append(array, right[r])
				r++
			} else {
				array = append(array, left[l])
				l++
			}
		}
	}
	return array
}
