package main

import (
	"./deck"
	"fmt"
)

func main() {
	cards := deck.New()

	for _, card := range cards {
		fmt.Println(card)
	}
}
