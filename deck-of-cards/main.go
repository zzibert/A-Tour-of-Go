package main

import (
	"./deck"
)

func main() {
	cards := deck.New(true, 0, []string{}, 1)

	card, cards := deck.Pop(cards)

}
