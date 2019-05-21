package main

import (
	"fmt"

	"./deck"
)

func main() {
	cards := deck.New()

	// cards = deck.Shuffle(cards)

	// cards = deck.DefaultSort(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}
