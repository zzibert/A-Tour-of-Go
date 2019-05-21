package main

import (
	"fmt"

	"./deck"
)

func main() {
	ranks := []deck.Rank{deck.Ace, deck.Two}
	cards := deck.New(deck.Filter(ranks))

	// cards = deck.Shuffle(cards)

	// cards = deck.DefaultSort(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}
