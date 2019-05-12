package main

import (
	"fmt"

	"./deck"
)

func main() {
	cards := deck.New(true, 0, []string{}, 1)
	player := make(deck.Deck, 0, 2)
	dealer := make(deck.Deck, 0, 2)

	for _, card := range cards[0:4] {
		fmt.Printf("%+v\n", card)
	}

	player, dealer, cards = deck.Turn(player, dealer, cards)
	player, dealer, cards = deck.Turn(player, dealer, cards)

	fmt.Println("player cards")
	for _, card := range player {
		fmt.Printf("%+v\n", card)
	}

	fmt.Println("dealer cards")
	for _, card := range dealer {
		fmt.Printf("%+v\n", card)
	}

}
