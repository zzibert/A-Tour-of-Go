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
	fmt.Println("What is the name of the text file?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	nameSlice := make([]Name, 0)

	file, err := os.Open(text)
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
