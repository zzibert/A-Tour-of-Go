package main

import "fmt"

func main() {
	fmt.Println(firstReccuringCharacter([]int{2, 5, 1, 2, 3, 5, 1, 2, 4}))
	fmt.Println(firstReccuringCharacter([]int{2, 1, 1, 2, 3, 5, 1, 2, 4}))
	fmt.Println(firstReccuringCharacter([]int{2, 3, 4, 5}))
}

// O(n)
func firstReccuringCharacter(nums []int) int {
	characters := make(map[int]bool)

	for _, number := range nums {
		if _, ok := characters[number]; ok {
			return number
		} else {
			characters[number] = true
		}
	}

	return 0
}
