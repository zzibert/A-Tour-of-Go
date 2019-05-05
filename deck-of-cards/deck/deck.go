package deck

import (
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	number string
	color  string
}

func New() []Card {
	colors := []string{"spades", "diamonds", "clubs", "hearts"}
	Deck := make([]Card, 0, 52)

	for _, color := range colors {
		for i := 0; i < 14; i++ {
			var number string
			switch i {
			case 0:
				number = "A"
			case 11:
				number = "J"
			case 12:
				number = "Q"
			case 13:
				number = "K"
			default:
				number = strconv.Itoa(i)
			}
			card := Card{number, color}
			Deck = append(Deck, card)
		}
	}
	return Deck
}

func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	newCards := make([]Card, 52)
	perm := r.Perm(52)
	for i, randIndex := range perm {
		newCards[i] = cards[randIndex]
	}
	return newCards
}
