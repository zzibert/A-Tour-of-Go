package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two , Suit: Diamond})
	fmt.Println(Card{Rank: Nine , Suit: Club})
	fmt.Println(Card{Rank: Jack , Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Diamonds
	// Nine of Clubs
	// Jack of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 4 Suits * 13 Ranks
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort) 
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	counter := 0
	for _, card := range cards {
		if card.Suit == Joker {
			counter++
		}
	}
	if counter != 3 {
		t.Error("Expected number of jokers in deck is 3. Received:", counter)
	}
}
