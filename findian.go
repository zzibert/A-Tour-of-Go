package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Println("Enter a string in double quotes.")
	fmt.Scanf("%q", &x)
	if strings.ContainsAny(x, "aA") && (x[0] == 'i' || x[0] == 'I') && (x[len(x)-1] == 'n' || x[len(x)-1] == 'N') {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
