package main

func checkIfExist(arr []int) bool {
	elements := make([]int, 0)

	if len(arr) == 0 {
		return false
	}

	for _, number := range arr {
		if isElement(elements, number) {
			return true
		}
		double := number * 2
		elements = append(elements, double)
		if number%2 == 0 {
			elements = append(elements, (number / 2))
		}
	}

	return false
}

func isElement(nums []int, val int) bool {
	for _, number := range nums {
		if number == val {
			return true
		}
	}

	return false
}
