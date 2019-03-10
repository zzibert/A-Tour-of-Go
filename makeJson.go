package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Person struct {
		Name, Address string
	}
	guy := Person{"", ""}

	fmt.Println("Enter your name!")
	fmt.Scanf("%s\n", &guy.Name)
	fmt.Println("Enter your address!")
	fmt.Scanf("%s\n", &guy.Address)
	jsonString, err := json.Marshal(guy)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonString))
	}
}
