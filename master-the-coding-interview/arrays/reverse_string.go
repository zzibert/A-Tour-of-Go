package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverse("Hi My name is Andrei"))
}

// This function has time complexity of O(n)
func reverse(str string) string {
	chars := strings.Split(str, "")
	var newString []string

	for _, val := range chars {
		newString = append([]string{val}, newString...)
	}

	return strings.Join(newString, "")
}
