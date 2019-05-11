package main

import (
	"fmt"

	"./deck"
)

func main() {
	cards := deck.New(false, 5, []string{"2", "3", "4"})

	for _, card := range cards {
		fmt.Println(card)
	}
}
