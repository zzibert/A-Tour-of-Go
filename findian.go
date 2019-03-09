package main

import (
	"fmt"
	"strings"
)

func main() {
	x := [5]int {1, 2, 3, 4, 5}
	for i, v range x {
		fmt.Println("index %d value %d", i, v)
	}
}
