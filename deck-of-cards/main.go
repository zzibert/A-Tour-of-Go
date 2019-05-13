package main

import (
	"bufio"
	"fmt"
	"os"

	"./deck"
)

func main() {
	cards := deck.New(true, 0, []string{}, 1)
	player := make(deck.Deck, 0, 2)
	dealer := make(deck.Deck, 0, 2)

	var action string
	stdin := bufio.NewReader(os.Stdin)

	player, dealer, cards = deck.Turn(player, dealer, cards)
	player, dealer, cards = deck.Turn(player, dealer, cards)

	for {
		deck.DisplayPlayerCards(player)

		fmt.Println("Hit Or Stand?")

		_, err := fmt.Fscan(stdin, &action)
		deck.Error(err)

		switch action {
		case "hit":
			fmt.Println("you selected hit!")
			player, cards = deck.Hit(player, cards)
		case "stand":
			fmt.Println("you selected stand!")
		default:
			fmt.Println("You didnt select nothing.")
		}
	}

}
