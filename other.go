package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var intList string

	fmt.Println("Please enter up to 10 integers separated by commas, e.g. 10,5,2")
	fmt.Scan(&intList)

	unsortedList := strings.Split(intList, ",")

	BubbleSort(unsortedList)
	fmt.Print("Sorted list: ", strings.Join(unsortedList, " "))
}

func BubbleSort(unsortedList []string) {
	for i := len(unsortedList) - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if ShouldSwap(unsortedList[j], unsortedList[j+1]) {
				swapped = true
				Swap(unsortedList, j)
			}
		}
		if !swapped {
			break
		}
		swapped = false
	}
}

func ShouldSwap(x, y string) bool {
	firstInt, firstErr := strconv.Atoi(x)
	secondInt, secondErr := strconv.Atoi(y)
	if firstErr != nil || secondErr != nil {
		fmt.Print("Each entry must be an integer. Please try again.")
	}
	return firstInt > secondInt
}

func Swap(intList []string, i int) {
	firstInt := intList[i]
	secondInt := intList[i+1]
	intList[i] = secondInt
	intList[i+1] = firstInt
}
