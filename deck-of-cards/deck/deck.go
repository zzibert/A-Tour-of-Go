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

type Deck []Card

func (deck Deck) Len() int {
	return len(deck)
}

func (deck Deck) Swap(i, j int) {
	deck[i], deck[j] = deck[j], deck[i]
}

func (deck Deck) Less(i, j int) bool {
	return true
}

func New() Deck {
	colors := []string{"spades", "diamonds", "clubs", "hearts"}
	Deck := make(Deck, 0, 52)

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

func Shuffle(cards Deck) Deck {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	newCards := make(Deck, 52)
	perm := r.Perm(52)
	for i, randIndex := range perm {
		newCards[i] = cards[randIndex]
	}
	return newCards
}
