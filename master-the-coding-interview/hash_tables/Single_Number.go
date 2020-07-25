package main

func singleNumber(nums []int) int {
	numberCount := make(map[int]int)

	for _, number := range nums {
		numberCount[number]++
	}

	for number, count := range numberCount {
		if count == 1 {
			return number
		}
	}
	return 0
}
