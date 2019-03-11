package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	type Name struct {
		fname string
		lname string
	}

	nameSlice := make([]Name, 0)

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		var name Name
		name.fname = words[0]
		name.lname = words[1]
		nameSlice = append(nameSlice, name)
	}
	for _, name := range nameSlice {
		fmt.Println(name.fname, name.lname)
	}
}
