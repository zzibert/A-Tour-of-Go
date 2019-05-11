package main

import (
	"fmt"

	"./deck"
)

func main() {
	cards := deck.New(true)

	for _, card := range cards {
		fmt.Println(card)
	}
}
