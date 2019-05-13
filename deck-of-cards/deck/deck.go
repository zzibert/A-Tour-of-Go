package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Card struct {
	number string
	color  string
}

type Deck []Card

// Sorting functions

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

// BlackJack functions

func multipleDecks(filterOut []string, numberOfDecks int) Deck {
	FinalDeck := make(Deck, 0)
	colors := []string{"spades", "diamonds", "clubs", "hearts"}
	for j := 0; j < numberOfDecks; j++ {
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
				if !Contains(filterOut, card.number) {
					Deck = append(Deck, card)
				}
			}
		}
		FinalDeck = append(FinalDeck, Deck...)
	}
	return FinalDeck
}

func New(shuffle bool, jokers int, filterOut []string, numberOfDecks int) Deck {
	cards := multipleDecks(filterOut, numberOfDecks)
	if !Contains(filterOut, "joker") {
		for i := 0; i < jokers; i++ {
			card := Card{"joker", "joker"}
			cards = append(cards, card)
		}
	}
	if shuffle {
		return Shuffle(cards)
	}
	sort.Sort(cards)
	return cards
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

func Turn(player Deck, dealer Deck, cards Deck) (Deck, Deck, Deck) {
	card, cards := Pop(cards)
	player = append(player, card)

	card, cards = Pop(cards)
	dealer = append(dealer, card)

	return player, dealer, cards
}

func Hit(player Deck, cards Deck) (Deck, Deck) {
	card, cards := Pop(cards)
	player = Push(card, player)
	return player, cards
}

func Stand() {}

func CalculatePoints(player Deck) int {
	counter, aceCounter := 0, 0
	for _, card := range player {
		switch card.number {
		case "1":
			counter++
		case "2":
			counter += 2
		case "3":
			counter += 3
		case "4":
			counter += 4
		case "5":
			counter += 5
		case "6":
			counter += 6
		case "7":
			counter += 7
		case "8":
			counter += 8
		case "9":
			counter += 9
		case "10":
			counter += 10
		case "J":
			counter += 10
		case "Q":
			counter += 10
		case "K":
			counter += 10
		case "A":
			aceCounter++
		}
	}
	for i := 0; i < aceCounter; i++ {
		if counter+11 > 21 {
			counter++
		} else {
			counter += 11
		}
	}
	return counter
}

func CheckWinner(player Deck, dealer Deck) {
	if CalculatePoints(player) > CalculatePoints(dealer) && CalculatePoints(player) <= 21 {
		fmt.Println("Player Won!")
	} else {
		fmt.Println("Dealer Won!")
	}
}

func CheckLoser(player Deck, dealer Deck) bool {
	if CalculatePoints(player) > 21 {
		fmt.Println("Dealer Won!")
		return true
	}
	return false
}

// Helper functions

func Pop(cards Deck) (Card, Deck) {
	return cards[0], cards[1:]
}

func Push(card Card, cards Deck) Deck {
	return append(cards, card)
}

func Contains(slice []string, element string) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}

func DisplayPlayerCards(player Deck) {
	for i := 0; i < len(player); i++ {
		fmt.Printf("%+v\n", player[i])
	}
}

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
