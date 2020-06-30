package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "bB"))
}

func numJewelsInStones(J string, S string) int {
	var counter int
	jewels := make(map[string]bool)
	stones := strings.Split(S, "")

	for _, jewel := range J {
		jewels[string(jewel)] = true
	}

	for _, stone := range stones {
		if jewels[stone] {
			counter++
		}
	}
	return counter
}
