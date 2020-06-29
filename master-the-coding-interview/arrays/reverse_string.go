package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverse("Hi My name is Andrei"))
	fmt.Println(reverse2("Hi My name is Andrei"))
}

// This function has time complexity of O(n)
func reverse(str string) string {
	chars := strings.Split(str, "")
	var newString []string

	for i, _ := range chars {
		j := (len(chars) - i) - 1
		newString = append(newString, chars[j])
	}

	return strings.Join(newString, "")
}

// Alternative version of reverse string , much simpler
func reverse2(str string) string {
	chars := strings.Split(str, "")
	var newString []string
	for i := len(chars) - 1; i >= 0; i-- {
		newString = append(newString, chars[i])
	}

	return strings.Join(newString, "")
}
