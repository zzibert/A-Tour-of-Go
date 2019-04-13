package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
