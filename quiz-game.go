package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//Opening the file
	f, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	//Reading the file
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range lines {
		fmt.Println(line[0])
	}

}
