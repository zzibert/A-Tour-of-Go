package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var pathToFile string
	sli := make([]name, 0, 3)
	fmt.Println("Please enter path to file:")
	fmt.Scan(&pathToFile)
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Please enter  valid file")
	} else {

		defer file.Close()
		scanner := bufio.NewScanner(file)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		} else {
			for scanner.Scan() {
				inputLine := strings.Fields(scanner.Text())
				runesfName := []rune(inputLine[0])
				fNameTrunc := string(runesfName[0:20])
				runeslName := []rune(inputLine[1])
				lNameTrunc := string(runeslName[0:20])
				sName := name{fname: fNameTrunc, lname: lNameTrunc}
				sli = append(sli, sName)

			}
			if len(sli) > 0 {
				fmt.Println("The content of the file are as follows. Iterating through the slice")

				for _, nameElement := range sli {
					fmt.Printf("First Name: %s LastName: %s\n", nameElement.fname, nameElement.lname)

				}
			} else {
				fmt.Println("The file is empty")

			}
		}
	}
}
