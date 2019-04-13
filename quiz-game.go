package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	//Initializing counter
	counter := 0

	//Getting the flags
	filePtr := flag.String("csv", "problems.csv", "the csv file")
	limitPtr := flag.Int("limit", 5, "time limit for each question")
	flag.Parse()

	//Opening the file
	f, err := os.Open(*filePtr)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	//Reading the file
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	//Initializing max score
	maxScore := len(lines)

	//Initializing the answer
	answer := 0

	//Loop through every line and compare the answer with result
	for _, line := range lines {
		fmt.Printf("What is the result of %s: ", line[0])

		//Set the timer for each loop
		timer := time.NewTimer(time.Duration(*limitPtr) * time.Second)

		//Run the timer in separate goroutine
		go func() {
			<-timer.C
			fmt.Printf("\nYour test score is %d out of %d", counter, maxScore)
			os.Exit(0)
		}()

		//Inputs the aswer
		_, err := fmt.Scanf("%d\n", &answer)
		if err != nil {
			fmt.Println(err)
		}
		//Stopping the timer
		timer.Stop()

		//Converts from string to int
		i, err := strconv.Atoi(line[1])

		//Compares results
		if answer == i {
			counter++
		}
	}
	fmt.Printf("Your test score is %d out of %d", counter, maxScore)
}
