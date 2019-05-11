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
	colors := []string{"spades", "diamonds", "clubs", "hearts"}
	numbers := []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	var iColorIndex, jColorIndex int
	for index, color := range colors {
		if deck[i].color == color {
			iColorIndex = index
		}
		if deck[j].color == color {
			jColorIndex = index
		}
	}
	if iColorIndex < jColorIndex {
		return true
	}
	if iColorIndex > jColorIndex {
		return false
	}
	var iNumberIndex, jNumberIndex int
	for index, number := range numbers {
		if deck[i].number == number {
			iNumberIndex = index
		}
		if deck[j].number == number {
			jNumberIndex = index
		}
	}
	if iNumberIndex < jNumberIndex {
		return true
	}
	if iNumberIndex > jNumberIndex {
		return false
	}
	return true
}

func New(shuffle bool, jokers int) Deck {
	colors := []string{"spades", "diamonds", "clubs", "hearts"}
	Deck := make(Deck, 0)
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
	for i := 0; i < jokers; i++ {
		card := Card{"joker", "joker"}
		Deck = append(Deck, card)
	}
	if shuffle {
		return Shuffle(Deck)
	}
	return Deck
}

func Shuffle(cards Deck) Deck {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	newCards := make(Deck, 0)
	perm := r.Perm(len(cards))
	for _, randIndex := range perm {
		newCards = append(newCards, cards[randIndex])
	}
	return newCards
}
