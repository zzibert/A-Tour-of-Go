package main

import (
	"./deck"
	"fmt"
)

func main() {
	cards := deck.New(deck.Shuffle)

	// cards = deck.Shuffle(cards)

	// cards = deck.DefaultSort(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}
