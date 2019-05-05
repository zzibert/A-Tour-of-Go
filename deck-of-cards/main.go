package main

import (
	"fmt"

	"./deck"
)

func main() {
	deck := deck.New()

	for _, card := range deck {
		fmt.Println(card)
	}
}
