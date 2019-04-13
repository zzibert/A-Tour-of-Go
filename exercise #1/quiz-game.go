package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//Initializing counter
	counter := 0

	//Getting the flags
	filePtr := flag.String("csv", "problems.csv", "the csv file")
	limitPtr := flag.Int("limit", 3, "time limit for each question")
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
	answer := ""

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
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			fmt.Println(err)
		}
		//Stopping the timer
		timer.Stop()

		//Sets everything to lowercase
		answer = strings.ToLower(answer)

		//Compares results
		if answer == line[1] {
			counter++
		}
	}
	//Prints quiz result
	fmt.Printf("Your test score is %d out of %d", counter, maxScore)
}
