package deck

import (
	"fmt"
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
