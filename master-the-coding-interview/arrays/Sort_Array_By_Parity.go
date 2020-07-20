package main

func sortArrayByParity(A []int) []int {
	odd := make([]int, 0)
	even := make([]int, 0)

	for _, number := range A {
		if number%2 == 0 {
			even = append(even, number)
		} else {
			odd = append(odd, number)
		}
	}

	return append(even, odd...)
}
