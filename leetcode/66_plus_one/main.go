package main

func plusOne(digits []int) []int {
	length := len(digits)

	digits[length-1] += 1

	for i := length - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i] = 0
			digits[i-1] += 1
		}
	}

	if digits[0] > 9 {
		digits[0] = 0
		return append([]int{1}, digits...)
	}

	return digits
}
