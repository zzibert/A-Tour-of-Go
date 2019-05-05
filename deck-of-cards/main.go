package main

import (
	"fmt"

	"./deck"
)

func main() {
	cards := deck.New()
	// cards = deck.Shuffle(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}
