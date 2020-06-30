package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}

//naive solution, O(n), count the zeroes the end at the end.
func moveZeroes(nums []int) []int {
	var counter int
	index := len(nums)

	for i := 0; i < index; {
		if nums[i] == 0 {
			counter++
			if i == index-1 {
				nums = nums[:i]
				i++
			} else {
				nums = append(nums[:i], nums[i+1:]...)
				index--
			}
		} else {
			i++
		}
	}

	for i := 0; i < counter; i++ {
		nums = append(nums, 0)
	}

	return nums
}
