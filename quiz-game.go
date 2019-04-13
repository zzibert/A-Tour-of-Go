package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Initializing counter
	counter := 0

	//Opening the file
	f, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	//Reading the file
	lines, err := csv.NewReader(f).ReadAll()

	//Initializing max score
	maxScore := len(lines)

	if err != nil {
		fmt.Println(err)
	}
	answer := 0

	//Loop through every line and compare the answer with result
	for _, line := range lines {
		fmt.Printf("What is the result of %s: ", line[0])

		//Inputs the aswer
		_, err := fmt.Scanf("%d\n", &answer)
		if err != nil {
			fmt.Println(err)
		}
		//Converts from string to int
		i, err := strconv.Atoi(line[1])

		//Compares results
		if answer == i {
			counter++
		}
	}
	fmt.Printf("Your test score is %d out of %d", counter, maxScore)
}
