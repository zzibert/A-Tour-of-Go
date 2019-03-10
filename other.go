package main

import (
	"fmt"
	"sort"
	"strconv"
)

func printSortedSlice(s []int) {
	// Sort a copy of the slice
	sortedSlice := make([]int, len(s), cap(s))
	copy(sortedSlice, s)
	sort.Ints(sortedSlice)

	fmt.Println(sortedSlice)
}

func main() {
	// Variables
	var input string
	var inputs = make([]int, 3)
	var inputsIdx int

	for {
		// Prompt user to enter an int and wait for the input
		fmt.Print("Please enter an int ('X' to quit): ")
		fmt.Scan(&input)

		// Handle "quit" criteria
		if input == "X" {
			break
		}

		if intInput, err := strconv.Atoi(input); err == nil {
			// Evaluate if we exceed initial slice capacity
			if inputsIdx <= 2 {
				inputs[inputsIdx] = intInput
				inputsIdx++
			} else {
				inputs = append(inputs, intInput)
			}

			// Print sorted slice
			printSortedSlice(inputs)
		} else {
			fmt.Println("Cannot evaluate your input as an int, please try again.")
		}
	}
}
