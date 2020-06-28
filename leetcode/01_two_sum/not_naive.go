package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	numbers := make(map[int]int, 0)
	for i, val := range nums {
		sub := target - val
		if j, ok := numbers[sub]; ok {
			return []int{j, i}
		} else {
			numbers[val] = i
		}
	}
	return []int{0}
}
