package main

import (
	"fmt"
	"sort"

	"./deck"
)

func main() {
	cards := deck.New()
	cards = deck.Shuffle(cards)
	sort.Sort(cards)
	cards = deck.Shuffle(cards)

	for _, card := range cards {
		fmt.Println(card)
	}
}
