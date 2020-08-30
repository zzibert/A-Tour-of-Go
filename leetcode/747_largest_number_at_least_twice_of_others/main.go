package main

func dominantIndex(nums []int) int {
	var largest, secondLargest, index int

	for i, number := range nums {
		if number > largest {
			secondLargest = largest
			largest = number
			index = i
		} else if number > secondLargest {
			secondLargest = number
		}
	}

	if largest >= secondLargest*2 {
		return index
	}

	return -1
}
