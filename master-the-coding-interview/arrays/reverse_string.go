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

	for i, _ := range chars {
		j := (len(chars) - i) - 1
		newString = append(newString, chars[j])
	}

	return strings.Join(newString, "")
}
